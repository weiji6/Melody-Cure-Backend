package service

import (
	"melody_cure/DAO"
	"melody_cure/middleware"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(image string, name string, password string, email string, identity string) error
	Login(email string, password string) (string, error)
	Logout() error
	ChangePassword(newPassword string) error
}

type User struct {
	dao DAO.User
	jwt middleware.JwtClient
}

func NewUser(dao DAO.User, jwt middleware.JwtClient) *User {
	return &User{
		dao: dao,
		jwt: jwt,
	}
}

func (u *User) Register(image string, name string, password string, email string, identity string) error {
	if name == "" || password == "" || email == "" {
		return errors.New("用户名、密码、邮箱均不能为空")
	}

	// 邮箱唯一性校验
	var count int64
	if err := DAO.DB.Model(&DAO.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return fmt.Errorf("校验邮箱失败: %w", err)
	}
	if count > 0 {
		return errors.New("该邮箱已被注册")
	}

	// 密码加密
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("密码加密失败: %w", err)
	}

	// 组装并写入
	newUser := DAO.User{
		Image:         image,
		Name:          name,
		Password:      string(hashed),
		Email:         email,
		Identity:      identity,
		Certification: false,
	}

	if err := DAO.DB.Create(&newUser).Error; err != nil {
		return fmt.Errorf("创建用户失败: %w", err)
	}

	return nil
}

func (u *User) Login(email string, password string) (string, error) {
	if email == "" || password == "" {
		return "", errors.New("邮箱与密码不能为空")
	}

	var user DAO.User
	if err := DAO.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", fmt.Errorf("用户不存在或密码错误: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("用户不存在或密码错误")
	}

	// 生成 JWT
	token, err := u.jwt.GenerateToken(user.ID)
	if err != nil {
		return "", fmt.Errorf("生成 token 失败: %w", err)
	}

	return token, nil
}

func (u *User) Logout() error {
	return nil
}

func (u *User) ChangePassword(newPassword string) error {
	if u.dao.Email == "" {
		return errors.New("修改密码需要提供用户邮箱（请在调用前赋值 u.dao.Email）")
	}
	if newPassword == "" {
		return errors.New("新密码不能为空")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("新密码加密失败: %w", err)
	}

	res := DAO.DB.Model(&DAO.User{}).Where("email = ?", u.dao.Email).Update("password", string(hashed))
	if res.Error != nil {
		return fmt.Errorf("更新密码失败: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return errors.New("未找到对应邮箱的用户，更新失败")
	}

	return nil
}

func (u *User) Certificate() error {
	if u.dao.ID == "" {
		return errors.New("需要提供用户 ID（请在调用前赋值 u.dao.ID）")
	}
	if u.dao.Certificate == "" {
		return errors.New("需要提供证书路径（请在调用前赋值 u.dao.Certificate）")
	}

	// 提交证书，认证状态置为未认证，待人工审核
	res := DAO.DB.Model(&DAO.User{}).Where("id = ?", u.dao.ID).Updates(map[string]interface{}{
		"certificate":   u.dao.Certificate,
		"certification": false,
	})
	if res.Error != nil {
		return fmt.Errorf("上传/更新证书失败: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return errors.New("未找到用户或证书未更新")
	}

	return nil
}

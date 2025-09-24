package service

import (
	"melody_cure/DAO"
	"melody_cure/middleware"
	"melody_cure/api/request"
	"errors"
	"fmt"
	"crypto/rand"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

// 生成UUID的简单实现
func generateUUID() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

type UserService interface {
	Register(image string, name string, password string, email string, identity string, phone string) error
	Login(email string, password string) (string, error)
	Logout(token string) error
	ChangePassword(userID string, oldPassword string, newPassword string) error
	Certificate(userID string, req *request.CertificationRequest) error
	
	// 个人信息管理
	GetProfile(userID string) (*DAO.User, error)
	UpdateProfile(userID string, req *request.UpdateProfileRequest) error
	
	// 认证相关
	ApplyCertification(userID string, req request.CertificationRequest) error
	GetCertificationStatus(userID string) (*DAO.Certification, error)
	
	// AI陪伴功能
	CreateAICompanion(userID string, req *request.AICompanionRequest) (*DAO.AICompanion, error)
	GetAICompanions(userID string) ([]DAO.AICompanion, error)
	
	// 虚拟疗愈导师
	CreateVirtualTherapist(userID string, req *request.VirtualTherapistRequest) (*DAO.VirtualTherapist, error)
	GetVirtualTherapists(userID string) ([]DAO.VirtualTherapist, error)
	
	// 儿童档案管理
	CreateChildArchive(userID string, req *request.ChildArchiveRequest) (*DAO.ChildArchive, error)
	GetChildArchives(userID string) ([]DAO.ChildArchive, error)
	UpdateChildArchive(userID string, archiveID string, req *request.ChildArchiveRequest) error
	DeleteChildArchive(archiveID string) error
	
	// 收藏功能
	AddFavorite(userID string, resourceType string, resourceID string) error
	GetFavorites(userID string) ([]DAO.UserFavorite, error)
	RemoveFavorite(userID string, resourceType string, resourceID string) error
	
	// 课程和游戏
	GetCourses() ([]DAO.Course, error)
	GetCourse(courseID string) (*DAO.Course, error)
	GetGames() ([]DAO.Game, error)
	GetGame(gameID string) (*DAO.Game, error)
}

type User struct {
	dao *DAO.UserDAO
	jwt *middleware.JwtClient
}

func NewUser(dao *DAO.UserDAO, jwt *middleware.JwtClient) *User {
	return &User{dao: dao, jwt: jwt}
}

func (u *User) Register(image string, name string, password string, email string, identity string, phone string) error {
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
		ID:            generateUUID(),
		Image:         image,
		Name:          name,
		Password:      string(hashed),
		Email:         email,
		Phone:         phone,
		Identity:      identity,
		Certification: false,
	}

	if err := DAO.DB.Create(&newUser).Error; err != nil {
		return fmt.Errorf("创建用户失败: %w", err)
	}

	return nil
}

func (u *User) Login(email string, password string) (string, error) {
	// 查找用户
	user, err := u.dao.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("用户不存在")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("密码错误")
	}

	// 生成token
	token, err := u.jwt.GenerateToken(user.ID)
	if err != nil {
		return "", errors.New("生成token失败")
	}

	return token, nil
}

func (u *User) Logout(token string) error {
	// 这里可以实现token黑名单机制
	return nil
}

func (u *User) ChangePassword(userID string, oldPassword string, newPassword string) error {
	// 获取用户信息
	user, err := u.dao.GetUserByID(userID)
	if err != nil {
		return errors.New("用户不存在")
	}

	// 验证旧密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if err != nil {
		return errors.New("旧密码错误")
	}

	// 加密新密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("密码加密失败")
	}

	// 更新密码
	user.Password = string(hashedPassword)
	return u.dao.UpdateUser(user)
}

func (u *User) Certificate(userID string, req *request.CertificationRequest) error {
	// 构建认证数据
	cert := &DAO.Certification{
		UserID:           userID,
		CertificateType:  req.CertificateType,
		CertificateName:  req.CertificateName,
		CertificateNo:    req.CertificateNo,
		IssuingAuthority: req.IssuingAuthority,
		IssueDate:        req.IssueDate,
		ExpiryDate:       req.ExpiryDate,
		Status:           "pending", // 默认状态为待审核
	}
	
	// 调用DAO层创建认证记录
	return u.dao.CreateCertification(cert)
}

// 获取个人信息
func (u *User) GetProfile(userID string) (*DAO.User, error) {
	var user DAO.User
	if err := DAO.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, fmt.Errorf("获取用户信息失败: %w", err)
	}
	return &user, nil
}

// 更新个人信息
func (u *User) UpdateProfile(userID string, req *request.UpdateProfileRequest) error {
	updates := make(map[string]interface{})
	
	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Phone != "" {
		updates["phone"] = req.Phone
	}
	if req.Address != "" {
		updates["address"] = req.Address
	}
	if req.Image != "" {
		updates["image"] = req.Image
	}

	if len(updates) == 0 {
		return errors.New("没有需要更新的字段")
	}

	if err := DAO.DB.Model(&DAO.User{}).Where("id = ?", userID).Updates(updates).Error; err != nil {
		return fmt.Errorf("更新用户信息失败: %w", err)
	}

	return nil
}

// 申请认证
func (u *User) ApplyCertification(userID string, req request.CertificationRequest) error {
	certification := DAO.Certification{
		ID:               generateUUID(),
		UserID:           userID,
		CertificateType:  req.CertificateType,
		CertificateName:  req.CertificateName,
		CertificateNo:    req.CertificateNo,
		IssuingAuthority: req.IssuingAuthority,
		IssueDate:        req.IssueDate,
		ExpiryDate:       req.ExpiryDate,
		Status:           "pending",
	}

	if err := DAO.DB.Create(&certification).Error; err != nil {
		return fmt.Errorf("提交认证申请失败: %w", err)
	}

	return nil
}

// 获取认证状态
func (u *User) GetCertificationStatus(userID string) (*DAO.Certification, error) {
	return u.dao.GetCertificationByUserID(userID)
}

// 创建AI陪伴
func (u *User) CreateAICompanion(userID string, req *request.AICompanionRequest) (*DAO.AICompanion, error) {
	// 构建AI陪伴数据
	companion := &DAO.AICompanion{
		UserID:        userID,
		CompanionType: req.CompanionType,
		Name:          req.Name,
		Avatar:        req.Avatar,
		Personality:   req.Personality,
		VoiceType:     req.VoiceType,
	}
	
	// 调用DAO层创建AI陪伴
	err := u.dao.CreateAICompanion(companion)
	if err != nil {
		return nil, err
	}
	
	return companion, nil
}

// 获取AI陪伴列表
func (u *User) GetAICompanions(userID string) ([]DAO.AICompanion, error) {
	return u.dao.GetAICompanionsByUserID(userID)
}

// 创建虚拟疗愈导师
func (u *User) CreateVirtualTherapist(userID string, req *request.VirtualTherapistRequest) (*DAO.VirtualTherapist, error) {
	// 构建虚拟疗愈导师数据
	therapist := &DAO.VirtualTherapist{
		UserID:         userID,
		TherapistType:  req.TherapistType,
		Name:           req.Name,
		Avatar:         req.Avatar,
		Specialization: req.Specialization,
		Experience:     req.Experience,
	}
	
	// 调用DAO层创建虚拟疗愈导师
	err := u.dao.CreateVirtualTherapist(therapist)
	if err != nil {
		return nil, err
	}
	
	return therapist, nil
}

// 获取虚拟疗愈导师列表
func (u *User) GetVirtualTherapists(userID string) ([]DAO.VirtualTherapist, error) {
	return u.dao.GetVirtualTherapistsByUserID(userID)
}

// 创建儿童档案
func (u *User) CreateChildArchive(userID string, req *request.ChildArchiveRequest) (*DAO.ChildArchive, error) {
	// 构建儿童档案数据
	archive := &DAO.ChildArchive{
		UserID:    userID,
		ChildName: req.ChildName,
		Gender:    req.Gender,
		BirthDate: req.BirthDate,
		Avatar:    req.Avatar,
		Condition: req.Condition,
		Diagnosis: req.Diagnosis,
		Treatment: req.Treatment,
		Progress:  req.Progress,
		Notes:     req.Notes,
	}
	
	// 调用DAO层创建儿童档案
	err := u.dao.CreateChildArchive(archive)
	if err != nil {
		return nil, err
	}
	
	return archive, nil
}

// 获取儿童档案列表
func (u *User) GetChildArchives(userID string) ([]DAO.ChildArchive, error) {
	return u.dao.GetChildArchivesByUserID(userID)
}

// 更新儿童档案
func (u *User) UpdateChildArchive(userID string, archiveID string, req *request.ChildArchiveRequest) error {
	// 构建更新数据
	archive := &DAO.ChildArchive{
		ID:        archiveID,
		UserID:    userID,
		ChildName: req.ChildName,
		Gender:    req.Gender,
		BirthDate: req.BirthDate,
		Avatar:    req.Avatar,
		Condition: req.Condition,
		Diagnosis: req.Diagnosis,
		Treatment: req.Treatment,
		Progress:  req.Progress,
		Notes:     req.Notes,
	}
	
	return u.dao.UpdateChildArchive(archive)
}

// 添加收藏
func (u *User) AddFavorite(userID string, resourceType string, resourceID string) error {
	favorite := &DAO.UserFavorite{
		UserID:       userID,
		ResourceType: resourceType,
		ResourceID:   resourceID,
	}
	
	return u.dao.AddFavorite(favorite)
}

// 获取收藏列表
func (u *User) GetFavorites(userID string) ([]DAO.UserFavorite, error) {
	return u.dao.GetFavoritesByUserID(userID)
}

func (u *User) RemoveFavorite(userID string, resourceType string, resourceID string) error {
	return u.dao.RemoveFavorite(userID, resourceType, resourceID)
}

func (u *User) DeleteChildArchive(archiveID string) error {
	return u.dao.DeleteChildArchive(archiveID)
}

func (u *User) GetCourses() ([]DAO.Course, error) {
	return u.dao.GetCourses()
}

func (u *User) GetCourse(courseID string) (*DAO.Course, error) {
	return u.dao.GetCourseByID(courseID)
}

func (u *User) GetGames() ([]DAO.Game, error) {
	return u.dao.GetGames()
}

func (u *User) GetGame(gameID string) (*DAO.Game, error) {
	return u.dao.GetGameByID(gameID)
}

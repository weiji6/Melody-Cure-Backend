package controller

import (
	"melody_cure/api/response"
	"melody_cure/api/request"
	"melody_cure/service"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *User {
	return &User{
		UserService: userService,
	}
}

func (u *User) Register(c *gin.Context) {
	var req request.Register
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, response.Response{Code: "400", Message: "参数错误: " + err.Error()})
		return
	}
	if req.Name == "" || req.Password == "" || req.Email == "" {
		c.JSON(400, response.Response{Code: "400", Message: "用户名、密码、邮箱均不能为空"})
		return
	}

	if err := u.UserService.Register("", req.Name, req.Password, req.Email, req.Identity); err != nil {
		c.JSON(400, response.Response{Code: "400", Message: err.Error()})
		return
	}

	c.JSON(200, response.Response{Code: "200", Message: "注册成功"})
}

func (u *User) Login(c *gin.Context) {
	type loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, response.Response{Code: "400", Message: "参数错误: " + err.Error()})
		return
	}
	if req.Email == "" || req.Password == "" {
		c.JSON(400, response.Response{Code: "400", Message: "邮箱与密码不能为空"})
		return
	}

	token, err := u.UserService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(401, response.Response{Code: "401", Message: err.Error()})
		return
	}

	c.JSON(200, response.Response{Code: "200", Message: "登录成功", Data: gin.H{"token": token}})
}

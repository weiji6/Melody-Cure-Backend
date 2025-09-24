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

// 用户登出
func (u *User) Logout(c *gin.Context) {
	// 从请求头获取token
	token := c.GetHeader("Authorization")
	
	err := u.UserService.Logout(token)
	if err != nil {
		c.JSON(500, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, response.SuccessResponse{
		Code:    200,
		Message: "登出成功",
	})
}

// 修改密码
func (u *User) ChangePassword(c *gin.Context) {
	var req struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, response.ErrorResponse{
			Code:    400,
			Message: "参数错误",
		})
		return
	}

	// 从JWT中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, response.ErrorResponse{
			Code:    401,
			Message: "未认证",
		})
		return
	}

	err := u.UserService.ChangePassword(userID.(string), req.OldPassword, req.NewPassword)
	if err != nil {
		c.JSON(500, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, response.SuccessResponse{
		Code:    200,
		Message: "密码修改成功",
	})
}

// 删除儿童档案
func (u *User) DeleteChildArchive(c *gin.Context) {
	archiveID := c.Param("id")
	if archiveID == "" {
		c.JSON(400, response.ErrorResponse{
			Code:    400,
			Message: "档案ID不能为空",
		})
		return
	}

	err := u.UserService.DeleteChildArchive(archiveID)
	if err != nil {
		c.JSON(500, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, response.SuccessResponse{
		Code:    200,
		Message: "删除成功",
	})
}

// 获取课程列表
func (u *User) GetCourses(c *gin.Context) {
	courses, err := u.UserService.GetCourses()
	if err != nil {
		c.JSON(500, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, response.SuccessResponse{
		Code:    200,
		Message: "获取成功",
		Data:    courses,
	})
}

// 获取单个课程
func (u *User) GetCourse(c *gin.Context) {
	courseID := c.Param("id")
	if courseID == "" {
		c.JSON(400, response.ErrorResponse{
			Code:    400,
			Message: "课程ID不能为空",
		})
		return
	}

	course, err := u.UserService.GetCourse(courseID)
	if err != nil {
		c.JSON(500, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, response.SuccessResponse{
		Code:    200,
		Message: "获取成功",
		Data:    course,
	})
}

// 获取游戏列表
func (u *User) GetGames(c *gin.Context) {
	games, err := u.UserService.GetGames()
	if err != nil {
		c.JSON(500, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, response.SuccessResponse{
		Code:    200,
		Message: "获取成功",
		Data:    games,
	})
}

// 获取单个游戏
func (u *User) GetGame(c *gin.Context) {
	gameID := c.Param("id")
	if gameID == "" {
		c.JSON(400, response.ErrorResponse{
			Code:    400,
			Message: "游戏ID不能为空",
		})
		return
	}

	game, err := u.UserService.GetGame(gameID)
	if err != nil {
		c.JSON(500, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, response.SuccessResponse{
		Code:    200,
		Message: "获取成功",
		Data:    game,
	})
}

func NewUserController(userService service.UserService) *User {
	return &User{
		UserService: userService,
	}
}

func (u *User) Register(c *gin.Context) {
	var req request.Register
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, response.ErrorResponse{
			Code:    400,
			Message: "参数错误: " + err.Error(),
		})
		return
	}
	if req.Name == "" || req.Password == "" || req.Email == "" {
		c.JSON(400, response.ErrorResponse{
			Code:    400,
			Message: "用户名、密码、邮箱均不能为空",
		})
		return
	}

	if err := u.UserService.Register("", req.Name, req.Password, req.Email, req.Identity, req.Phone); err != nil {
		c.JSON(400, response.ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, response.SuccessResponse{
		Code:    200,
		Message: "注册成功",
	})
}

func (u *User) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, response.ErrorResponse{
			Code:    400,
			Message: "参数错误",
		})
		return
	}
	if req.Name == "" || req.Password == "" {
		c.JSON(400, response.ErrorResponse{
			Code:    400,
			Message: "用户名和密码不能为空",
		})
		return
	}

	token, err := u.UserService.Login(req.Name, req.Password)
	if err != nil {
		c.JSON(401, response.ErrorResponse{
			Code:    401,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, response.SuccessResponse{
		Code:    200,
		Message: "登录成功",
		Data:    gin.H{"token": token},
	})
}

// 获取个人信息
func (u *User) GetProfile(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(401, response.ErrorResponse{
			Code:    401,
			Message: "未授权",
		})
		return
	}

	profile, err := u.UserService.GetProfile(userID)
	if err != nil {
		c.JSON(500, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, response.SuccessResponse{
		Code:    200,
		Message: "获取成功",
		Data:    profile,
	})
}

// 更新个人信息
func (u *User) UpdateProfile(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(401, response.ErrorResponse{
			Code:    401,
			Message: "未授权",
		})
		return
	}

	var req request.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, response.ErrorResponse{
			Code:    400,
			Message: "参数错误",
		})
		return
	}

	if err := u.UserService.UpdateProfile(userID, 
		
		&req); err != nil {
		c.JSON(500, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, response.SuccessResponse{
		Code:    200,
		Message: "更新成功",
	})
}

// 申请认证
func (u *User) ApplyCertification(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(401, response.ErrorResponse{
			Code:    401,
			Message: "未授权",
		})
		return
	}

	var req request.CertificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, response.ErrorResponse{
			Code:    400,
			Message: "参数错误",
		})
		return
	}

	if err := u.UserService.Certificate(userID, &req); err != nil {
		c.JSON(500, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, response.SuccessResponse{
		Code:    200,
		Message: "申请提交成功",
	})
}

// 获取认证状态
func (u *User) GetCertificationStatus(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(401, response.ErrorResponse{
			Code:    401,
			Message: "未授权",
		})
		return
	}

	status, err := u.UserService.GetCertificationStatus(userID)
	if err != nil {
		c.JSON(500, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, response.SuccessResponse{
		Code:    200,
		Message: "获取成功",
		Data:    status,
	})
}

// 创建AI陪伴
func (u *User) CreateAICompanion(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(401, response.ErrorResponse{
			Code:    401,
			Message: "未授权",
		})
		return
	}

	var req request.AICompanionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, response.ErrorResponse{
			Code:    400,
			Message: "参数错误",
		})
		return
	}

	companion, err := u.UserService.CreateAICompanion(userID, &req)
	if err != nil {
		c.JSON(500, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, response.SuccessResponse{
		Code:    200,
		Message: "创建成功",
		Data:    companion,
	})
}

// 获取AI陪伴列表
func (u *User) GetAICompanions(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(401, response.ErrorResponse{
			Code:    401,
			Message: "未授权",
		})
		return
	}

	companions, err := u.UserService.GetAICompanions(userID)
	if err != nil {
		c.JSON(500, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, response.SuccessResponse{
		Code:    200,
		Message: "获取成功",
		Data:    companions,
	})
}

// 创建虚拟疗愈导师
func (u *User) CreateVirtualTherapist(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(401, response.ErrorResponse{
			Code:    401,
			Message: "未授权",
		})
		return
	}

	var req request.VirtualTherapistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, response.ErrorResponse{
			Code:    400,
			Message: "参数错误",
		})
		return
	}

	therapist, err := u.UserService.CreateVirtualTherapist(userID, &req)
	if err != nil {
		c.JSON(500, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, response.SuccessResponse{
		Code:    200,
		Message: "创建成功",
		Data:    therapist,
	})
}

// 获取虚拟疗愈导师列表
func (u *User) GetVirtualTherapists(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(401, response.ErrorResponse{
			Code:    401,
			Message: "未授权",
		})
		return
	}

	therapists, err := u.UserService.GetVirtualTherapists(userID)
	if err != nil {
		c.JSON(500, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, response.SuccessResponse{
		Code:    200,
		Message: "获取成功",
		Data:    therapists,
	})
}

// 创建儿童档案
func (u *User) CreateChildArchive(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(401, response.ErrorResponse{
			Code:    401,
			Message: "未授权",
		})
		return
	}

	var req request.ChildArchiveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, response.ErrorResponse{
			Code:    400,
			Message: "参数错误",
		})
		return
	}

	archive, err := u.UserService.CreateChildArchive(userID, &req)
	if err != nil {
		c.JSON(500, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, response.SuccessResponse{
		Code:    200,
		Message: "创建成功",
		Data:    archive,
	})
}

// 获取儿童档案列表
func (u *User) GetChildArchives(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(401, response.ErrorResponse{
			Code:    401,
			Message: "未授权",
		})
		return
	}

	archives, err := u.UserService.GetChildArchives(userID)
	if err != nil {
		c.JSON(500, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, response.SuccessResponse{
		Code:    200,
		Message: "获取成功",
		Data:    archives,
	})
}

// 更新儿童档案
func (u *User) UpdateChildArchive(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(401, response.ErrorResponse{
			Code:    401,
			Message: "未授权",
		})
		return
	}

	archiveID := c.Param("id")
	if archiveID == "" {
		c.JSON(400, response.ErrorResponse{
			Code:    400,
			Message: "档案ID不能为空",
		})
		return
	}

	var req request.ChildArchiveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, response.ErrorResponse{
			Code:    400,
			Message: "参数错误",
		})
		return
	}

	if err := u.UserService.UpdateChildArchive(userID, archiveID, &req); err != nil {
		c.JSON(500, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, response.SuccessResponse{
		Code:    200,
		Message: "更新成功",
	})
}

// 添加收藏
func (u *User) AddFavorite(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(401, response.ErrorResponse{
			Code:    401,
			Message: "未授权",
		})
		return
	}

	var req request.FavoriteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, response.ErrorResponse{
			Code:    400,
			Message: "参数错误",
		})
		return
	}

	if err := u.UserService.AddFavorite(userID, req.ResourceType, req.ResourceID); err != nil {
		c.JSON(500, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, response.SuccessResponse{
		Code:    200,
		Message: "收藏成功",
	})
}

// 获取收藏列表
func (u *User) GetFavorites(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(401, response.ErrorResponse{
			Code:    401,
			Message: "未授权",
		})
		return
	}

	favorites, err := u.UserService.GetFavorites(userID)
	if err != nil {
		c.JSON(500, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, response.SuccessResponse{
		Code:    200,
		Message: "获取成功",
		Data:    favorites,
	})
}

// 取消收藏
func (u *User) RemoveFavorite(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(401, response.ErrorResponse{
			Code:    401,
			Message: "未授权",
		})
		return
	}

	var req request.FavoriteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, response.ErrorResponse{
			Code:    400,
			Message: "参数错误",
		})
		return
	}

	if err := u.UserService.RemoveFavorite(userID, req.ResourceType, req.ResourceID); err != nil {
		c.JSON(500, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, response.SuccessResponse{
		Code:    200,
		Message: "取消收藏成功",
	})
}

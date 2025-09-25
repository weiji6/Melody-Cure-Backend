package controller

import (
	"melody_cure/api/response"
	"melody_cure/api/request"
	"melody_cure/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserService service.UserService
}

// Logout 用户登出
// @Summary 用户登出
// @Description 用户登出，使token失效
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.SuccessResponse "登出成功"
// @Failure 500 {object} response.ErrorResponse "服务器错误"
// @Router /api/user/logout [post]
func (u *User) Logout(c *gin.Context) {
	// 从请求头获取token
	token := c.GetHeader("Authorization")
	
	err := u.UserService.Logout(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Code:    200,
		Message: "登出成功",
	})
}

// ChangePassword 修改密码
// @Summary 修改密码
// @Description 用户修改登录密码
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body object{old_password=string,new_password=string} true "修改密码请求"
// @Success 200 {object} response.SuccessResponse "密码修改成功"
// @Failure 400 {object} response.ErrorResponse "参数错误"
// @Failure 401 {object} response.ErrorResponse "未认证"
// @Failure 500 {object} response.ErrorResponse "服务器错误"
// @Router /api/user/password [put]
func (u *User) ChangePassword(c *gin.Context) {
	var req struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: "参数错误",
		})
		return
	}

	// 从JWT中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Code:    401,
			Message: "未认证",
		})
		return
	}

	err := u.UserService.ChangePassword(userID.(string), req.OldPassword, req.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Code:    200,
		Message: "密码修改成功",
	})
}



// GetCourses 获取课程列表
// @Summary 获取课程列表
// @Description 获取所有课程列表
// @Tags 内容管理
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=[]DAO.Course} "获取成功"
// @Failure 500 {object} response.ErrorResponse "服务器错误"
// @Router /api/courses [get]
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

// GetCourse 获取课程详情
// @Summary 获取课程详情
// @Description 根据ID获取课程详情
// @Tags 内容管理
// @Accept json
// @Produce json
// @Param id path string true "课程ID"
// @Success 200 {object} response.SuccessResponse{data=DAO.Course} "获取成功"
// @Failure 400 {object} response.ErrorResponse "参数错误"
// @Failure 500 {object} response.ErrorResponse "服务器错误"
// @Router /api/course/{id} [get]
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

// GetGames 获取游戏列表
// @Summary 获取游戏列表
// @Description 获取所有游戏列表
// @Tags 内容管理
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=[]DAO.Game} "获取成功"
// @Failure 500 {object} response.ErrorResponse "服务器错误"
// @Router /api/games [get]
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

// GetGame 获取游戏详情
// @Summary 获取游戏详情
// @Description 根据ID获取游戏详情
// @Tags 内容管理
// @Accept json
// @Produce json
// @Param id path string true "游戏ID"
// @Success 200 {object} response.SuccessResponse{data=DAO.Game} "获取成功"
// @Failure 400 {object} response.ErrorResponse "参数错误"
// @Failure 500 {object} response.ErrorResponse "服务器错误"
// @Router /api/game/{id} [get]
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

// Register 用户注册
// @Summary 用户注册
// @Description 用户注册接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body request.Register true "注册请求"
// @Success 200 {object} response.SuccessResponse "注册成功"
// @Failure 400 {object} response.ErrorResponse "参数错误"
// @Router /api/user/register [post]
func (u *User) Register(c *gin.Context) {
	var req request.Register
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: "参数错误: " + err.Error(),
		})
		return
	}
	if req.Name == "" || req.Password == "" || req.Email == "" {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: "用户名、密码、邮箱均不能为空",
		})
		return
	}

	if err := u.UserService.Register("", req.Name, req.Password, req.Email, req.Identity, req.Phone); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Code:    200,
		Message: "注册成功",
	})
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body request.LoginRequest true "登录请求"
// @Success 200 {object} response.SuccessResponse{data=string} "登录成功，返回token"
// @Failure 400 {object} response.ErrorResponse "参数错误"
// @Failure 401 {object} response.ErrorResponse "用户名或密码错误"
// @Router /api/user/login [post]
func (u *User) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: "参数错误",
		})
		return
	}
	if req.Name == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: "用户名和密码不能为空",
		})
		return
	}

	token, err := u.UserService.Login(req.Name, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Code:    401,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Code:    200,
		Message: "登录成功",
		Data:    gin.H{"token": token},
	})
}

// GetProfile 获取个人信息
// @Summary 获取个人信息
// @Description 获取用户个人信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.SuccessResponse{data=DAO.User} "获取成功"
// @Failure 401 {object} response.ErrorResponse "未授权"
// @Failure 500 {object} response.ErrorResponse "服务器错误"
// @Router /api/user/profile [get]
func (u *User) GetProfile(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Code:    401,
			Message: "未授权",
		})
		return
	}

	profile, err := u.UserService.GetProfile(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Code:    200,
		Message: "获取成功",
		Data:    profile,
	})
}

// UpdateProfile 更新个人信息
// @Summary 更新个人信息
// @Description 更新用户个人信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body request.UpdateProfileRequest true "更新信息请求"
// @Success 200 {object} response.SuccessResponse "更新成功"
// @Failure 400 {object} response.ErrorResponse "参数错误"
// @Failure 401 {object} response.ErrorResponse "未授权"
// @Failure 500 {object} response.ErrorResponse "服务器错误"
// @Router /api/user/profile [put]
func (u *User) UpdateProfile(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Code:    401,
			Message: "未授权",
		})
		return
	}

	var req request.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: "参数错误",
		})
		return
	}

	if err := u.UserService.UpdateProfile(userID, &req); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Code:    200,
		Message: "更新成功",
	})
}

// ApplyCertification 申请认证
// @Summary 申请认证
// @Description 用户申请专业认证
// @Tags 认证管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body request.CertificationRequest true "认证申请请求"
// @Success 200 {object} response.SuccessResponse "申请成功"
// @Failure 400 {object} response.ErrorResponse "参数错误"
// @Failure 401 {object} response.ErrorResponse "未授权"
// @Failure 500 {object} response.ErrorResponse "服务器错误"
// @Router /api/user/certification [post]
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

	if err := u.UserService.ApplyCertification(userID, req); err != nil {
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



// CreateAICompanion 创建AI陪伴
// @Summary 创建AI陪伴
// @Description 创建AI陪伴角色
// @Tags AI陪伴
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body request.AICompanionRequest true "AI陪伴创建请求"
// @Success 200 {object} response.SuccessResponse{data=DAO.AICompanion} "创建成功"
// @Failure 400 {object} response.ErrorResponse "参数错误"
// @Failure 401 {object} response.ErrorResponse "未授权"
// @Failure 500 {object} response.ErrorResponse "服务器错误"
// @Router /api/user/ai-companion [post]
func (u *User) CreateAICompanion(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Code:    401,
			Message: "未授权",
		})
		return
	}

	var req request.AICompanionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: "参数错误",
		})
		return
	}

	companion, err := u.UserService.CreateAICompanion(userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Code:    200,
		Message: "创建成功",
		Data:    companion,
	})
}

// GetAICompanions 获取AI陪伴列表
// @Summary 获取AI陪伴列表
// @Description 获取用户的AI陪伴列表
// @Tags AI陪伴
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.SuccessResponse{data=[]DAO.AICompanion} "获取成功"
// @Failure 401 {object} response.ErrorResponse "未授权"
// @Failure 500 {object} response.ErrorResponse "服务器错误"
// @Router /api/user/ai-companions [get]
func (u *User) GetAICompanions(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Code:    401,
			Message: "未授权",
		})
		return
	}

	companions, err := u.UserService.GetAICompanions(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Code:    200,
		Message: "获取成功",
		Data:    companions,
	})
}

// CreateVirtualTherapist 创建虚拟疗愈导师
// @Summary 创建虚拟疗愈导师
// @Description 创建虚拟疗愈导师
// @Tags 虚拟疗愈导师
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body request.VirtualTherapistRequest true "虚拟疗愈导师创建请求"
// @Success 200 {object} response.SuccessResponse{data=DAO.VirtualTherapist} "创建成功"
// @Failure 400 {object} response.ErrorResponse "参数错误"
// @Failure 401 {object} response.ErrorResponse "未授权"
// @Failure 500 {object} response.ErrorResponse "服务器错误"
// @Router /api/user/virtual-therapist [post]
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

// GetVirtualTherapists 获取虚拟疗愈导师列表
// @Summary 获取虚拟疗愈导师列表
// @Description 获取用户的虚拟疗愈导师列表
// @Tags 虚拟疗愈导师
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.SuccessResponse{data=[]DAO.VirtualTherapist} "获取成功"
// @Failure 401 {object} response.ErrorResponse "未授权"
// @Failure 500 {object} response.ErrorResponse "服务器错误"
// @Router /api/user/virtual-therapists [get]
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

// CreateChildArchive 创建儿童档案
// @Summary 创建儿童档案
// @Description 创建儿童档案
// @Tags 儿童档案管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body request.ChildArchiveRequest true "儿童档案创建请求"
// @Success 200 {object} response.SuccessResponse{data=DAO.ChildArchive} "创建成功"
// @Failure 400 {object} response.ErrorResponse "参数错误"
// @Failure 401 {object} response.ErrorResponse "未授权"
// @Failure 500 {object} response.ErrorResponse "服务器错误"
// @Router /api/user/child-archive [post]
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

// GetChildArchives 获取儿童档案列表
// @Summary 获取儿童档案列表
// @Description 获取用户的儿童档案列表
// @Tags 儿童档案管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.SuccessResponse{data=[]DAO.ChildArchive} "获取成功"
// @Failure 401 {object} response.ErrorResponse "未授权"
// @Failure 500 {object} response.ErrorResponse "服务器错误"
// @Router /api/user/child-archives [get]
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

// UpdateChildArchive 更新儿童档案
// @Summary 更新儿童档案
// @Description 更新儿童档案信息
// @Tags 儿童档案管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "档案ID"
// @Param request body request.ChildArchiveRequest true "儿童档案更新请求"
// @Success 200 {object} response.SuccessResponse "更新成功"
// @Failure 400 {object} response.ErrorResponse "参数错误"
// @Failure 401 {object} response.ErrorResponse "未授权"
// @Failure 500 {object} response.ErrorResponse "服务器错误"
// @Router /api/user/child-archive/{id} [put]
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

// DeleteChildArchive 删除儿童档案
// @Summary 删除儿童档案
// @Description 删除儿童档案
// @Tags 儿童档案管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "档案ID"
// @Success 200 {object} response.SuccessResponse "删除成功"
// @Failure 400 {object} response.ErrorResponse "参数错误"
// @Failure 401 {object} response.ErrorResponse "未授权"
// @Failure 500 {object} response.ErrorResponse "服务器错误"
// @Router /api/user/child-archive/{id} [delete]
func (u *User) DeleteChildArchive(c *gin.Context) {
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

// AddFavorite 添加收藏
// @Summary 添加收藏
// @Description 添加内容到收藏夹
// @Tags 收藏管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body request.FavoriteRequest true "收藏请求"
// @Success 200 {object} response.SuccessResponse "添加成功"
// @Failure 400 {object} response.ErrorResponse "参数错误"
// @Failure 401 {object} response.ErrorResponse "未授权"
// @Failure 500 {object} response.ErrorResponse "服务器错误"
// @Router /api/user/favorite [post]
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

// GetFavorites 获取收藏列表
// @Summary 获取收藏列表
// @Description 获取用户的收藏列表
// @Tags 收藏管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.SuccessResponse{data=[]DAO.UserFavorite} "获取成功"
// @Failure 401 {object} response.ErrorResponse "未授权"
// @Failure 500 {object} response.ErrorResponse "服务器错误"
// @Router /api/user/favorites [get]
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

// RemoveFavorite 移除收藏
// @Summary 移除收藏
// @Description 从收藏夹移除内容
// @Tags 收藏管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body request.FavoriteRequest true "收藏移除请求"
// @Success 200 {object} response.SuccessResponse "移除成功"
// @Failure 400 {object} response.ErrorResponse "参数错误"
// @Failure 401 {object} response.ErrorResponse "未授权"
// @Failure 500 {object} response.ErrorResponse "服务器错误"
// @Router /api/user/favorite [delete]
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
		Message: "移除成功",
	})
}









// GetCertificationStatus 获取认证状态
// @Summary 获取认证状态
// @Description 获取用户认证状态
// @Tags 认证管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.SuccessResponse{data=DAO.Certification} "获取成功"
// @Failure 401 {object} response.ErrorResponse "未授权"
// @Failure 500 {object} response.ErrorResponse "服务器错误"
// @Router /api/user/certification [get]
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

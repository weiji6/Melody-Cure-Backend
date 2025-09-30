package controller

import (
	"melody_cure/DAO"
	"melody_cure/api/response"
	"melody_cure/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChildArchiveController struct {
	userService *service.User
}

func NewChildArchiveController(userService *service.User) *ChildArchiveController {
	return &ChildArchiveController{userService: userService}
}

// GetChildProfile 获取儿童个人信息
// @Summary 获取儿童个人信息
// @Description 获取儿童的个人信息，包括照片、姓名、年龄、性别、诊断结果、已疗愈天数等
// @Tags 儿童档案
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param archive_id path string true "儿童档案ID"
// @Success 200 {object} object{code=int,data=response.ChildArchiveResponse} "获取成功"
// @Failure 400 {object} response.ErrorResponse "参数错误"
// @Failure 401 {object} response.ErrorResponse "未认证"
// @Failure 404 {object} response.ErrorResponse "档案不存在"
// @Failure 500 {object} response.ErrorResponse "获取失败"
// @Router /api/child-archive/{archive_id}/profile [get]
func (c *ChildArchiveController) GetChildProfile(ctx *gin.Context) {
	archiveID := ctx.Param("archive_id")
	if archiveID == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Code: http.StatusBadRequest, Message: "档案ID不能为空"})
		return
	}

	// 从JWT获取用户ID
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{Code: http.StatusUnauthorized, Message: "未认证"})
		return
	}

	// 获取用户的所有儿童档案
	archives, err := c.userService.GetChildArchives(userID.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Code: http.StatusInternalServerError, Message: "获取档案失败: " + err.Error()})
		return
	}

	// 查找指定的档案
	var targetArchive *DAO.ChildArchive
	for _, archive := range archives {
		if archive.ID == archiveID {
			targetArchive = &archive
			break
		}
	}

	if targetArchive == nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{Code: http.StatusNotFound, Message: "档案不存在"})
		return
	}

	// 转换为响应格式
	profileResponse := response.ToChildArchiveResponse(targetArchive)

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": profileResponse})
}

// GetChildArchives 获取用户的所有儿童档案列表
// @Summary 获取儿童档案列表
// @Description 获取当前用户的所有儿童档案列表
// @Tags 儿童档案
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} object{code=int,data=[]response.ChildArchiveResponse} "获取成功"
// @Failure 401 {object} response.ErrorResponse "未认证"
// @Failure 500 {object} response.ErrorResponse "获取失败"
// @Router /api/child-archive [get]
func (c *ChildArchiveController) GetChildArchives(ctx *gin.Context) {
	// 从JWT获取用户ID
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{Code: http.StatusUnauthorized, Message: "未认证"})
		return
	}

	// 获取用户的所有儿童档案
	archives, err := c.userService.GetChildArchives(userID.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Code: http.StatusInternalServerError, Message: "获取档案失败: " + err.Error()})
		return
	}

	// 转换为响应格式
	var archiveResponses []response.ChildArchiveResponse
	for _, archive := range archives {
		archiveResponses = append(archiveResponses, response.ToChildArchiveResponse(&archive))
	}

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": archiveResponses})
}
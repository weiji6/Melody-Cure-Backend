package controller

import (
	"melody_cure/api/response"
	"melody_cure/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageController struct {
	imageService *service.ImageService
}

// NewImageController 创建ImageController实例
func NewImageController(imageService *service.ImageService) *ImageController {
	return &ImageController{
		imageService: imageService,
	}
}

// GetQiniuUploadToken 获取七牛云上传token
// @Summary 获取七牛云上传token
// @Description 获取七牛云图床上传凭证，用于前端直接上传文件到七牛云
// @Tags 图床管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} object{code=int,message=string,data=object{token=string,domain=string,bucket=string,expires_at=int64,use_https=bool}} "获取上传token成功"
// @Failure 500 {object} response.ErrorResponse "生成token失败"
// @Router /api/image/qiniu/token [get]
func (c *ImageController) GetQiniuUploadToken(ctx *gin.Context) {
	// 生成七牛云上传token
	tokenInfo, err := c.imageService.GenerateQiniuUploadToken()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "生成上传token失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "获取上传token成功",
		"data":    tokenInfo,
	})
}
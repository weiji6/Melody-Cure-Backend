package routes

import (
	"melody_cure/controller"
	"melody_cure/middleware"

	"github.com/gin-gonic/gin"
)

// SetupImageRoutes 设置图床相关路由
func SetupImageRoutes(router *gin.Engine, imageController *controller.ImageController, jwtMiddleware *middleware.JwtClient) {
	// 图床API路由组
	imageGroup := router.Group("/api/image")
	
	// 需要认证的路由
	protected := imageGroup.Group("")
	protected.Use(jwtMiddleware.AuthMiddleware())
	{
		// 获取七牛云上传token
		protected.GET("/qiniu/token", imageController.GetQiniuUploadToken)
	}
}
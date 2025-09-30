package routes

import (
	"melody_cure/controller"
	"melody_cure/middleware"

	"github.com/gin-gonic/gin"
)

// 设置儿童档案相关路由
func SetupChildArchiveRoutes(router *gin.Engine, childArchiveController *controller.ChildArchiveController, jwtMiddleware *middleware.JwtClient) {
	// 儿童档案路由组，需要JWT认证
	childArchiveGroup := router.Group("/api/child-archive")
	childArchiveGroup.Use(jwtMiddleware.AuthMiddleware())
	{
		// 获取儿童个人信息
		childArchiveGroup.GET("/:archiveId/profile", childArchiveController.GetChildProfile)
		
		// 获取用户的所有儿童档案列表
		childArchiveGroup.GET("/list", childArchiveController.GetChildArchives)
	}
}
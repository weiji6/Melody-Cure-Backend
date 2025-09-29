package routes

import (
	"melody_cure/controller"
	"melody_cure/middleware"

	"github.com/gin-gonic/gin"
)

// SetupHealingLogRoutes 设置疗愈日志相关路由
func SetupHealingLogRoutes(router *gin.Engine, healingLogController *controller.HealingLogController, jwtMiddleware *middleware.JwtClient) {
	healingLogGroup := router.Group("/api/healing-log")

	// 需要认证的路由
	protected := healingLogGroup.Group("")
	protected.Use(jwtMiddleware.AuthMiddleware())
	{
		protected.POST("", healingLogController.CreateHealingLog)
		protected.GET("/child/:child_id", healingLogController.GetHealingLogsByChildID)
		protected.GET("/:log_id", healingLogController.GetHealingLogByID)
		protected.DELETE("/:log_id", healingLogController.DeleteHealingLog)
	}
}
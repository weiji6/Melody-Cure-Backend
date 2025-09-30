package routes

import (
	"melody_cure/controller"
	"melody_cure/middleware"

	"github.com/gin-gonic/gin"
)

// SetupAIReportRoutes 设置AI报告相关路由
func SetupAIReportRoutes(router *gin.Engine, aiReportController *controller.AIReportController, jwtMiddleware *middleware.JwtClient) {
	// AI报告相关路由组
	aiReportGroup := router.Group("/api/ai-reports")
	aiReportGroup.Use(jwtMiddleware.AuthMiddleware())
	{
		// 生成AI报告
		aiReportGroup.POST("/generate", aiReportController.GenerateReport)
		
		// 获取AI报告
		aiReportGroup.GET("", aiReportController.GetReport)
		
		// 更新AI报告内容
		aiReportGroup.PUT("/:id", aiReportController.UpdateReportContent)
	}
}
package routes

import (
	"melody_cure/controller"
	"melody_cure/middleware"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, userController *controller.User, jwtMiddleware *middleware.JwtClient) {
	// 公开路由（无需认证）
	public := r.Group("/api/user")
	{
		public.POST("/register", userController.Register)
		public.POST("/login", userController.Login)
	}

	// 需要认证的路由
	protected := r.Group("/api/user")
	protected.Use(jwtMiddleware.AuthMiddleware())
	{
		// 个人信息管理
		protected.GET("/profile", userController.GetProfile)
		protected.PUT("/profile", userController.UpdateProfile)
		protected.POST("/logout", userController.Logout)
		protected.PUT("/password", userController.ChangePassword)

		// 认证相关
		protected.POST("/certification/apply", userController.ApplyCertification)
		protected.GET("/certification/status", userController.GetCertificationStatus)

		// AI陪伴功能
		protected.POST("/ai-companion", userController.CreateAICompanion)
		protected.GET("/ai-companions", userController.GetAICompanions)

		// 虚拟疗愈导师
		protected.POST("/virtual-therapist", userController.CreateVirtualTherapist)
		protected.GET("/virtual-therapists", userController.GetVirtualTherapists)

		// 儿童档案管理
		protected.POST("/child-archive", userController.CreateChildArchive)
		protected.GET("/child-archives", userController.GetChildArchives)
		protected.PUT("/child-archive/:id", userController.UpdateChildArchive)
		protected.DELETE("/child-archive/:id", userController.DeleteChildArchive)

		// 收藏功能
		protected.POST("/favorite", userController.AddFavorite)
		protected.GET("/favorites", userController.GetFavorites)
		protected.DELETE("/favorite", userController.RemoveFavorite)

		// 课程和游戏（只读）
		protected.GET("/courses", userController.GetCourses)
		protected.GET("/course/:id", userController.GetCourse)
		protected.GET("/games", userController.GetGames)
		protected.GET("/game/:id", userController.GetGame)
	}
}
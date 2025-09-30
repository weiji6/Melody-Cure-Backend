//go:build wireinject
// +build wireinject

package main

import (
	"melody_cure/DAO"
	"melody_cure/config"
	"melody_cure/controller"
	"melody_cure/middleware"
	"melody_cure/routes"
	"melody_cure/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type App struct {
	Engine *gin.Engine
}

var ProviderSet = wire.NewSet(
	DAO.NewDB,
	DAO.NewUserDAO,
	DAO.NewHealingLogDAO,
	DAO.NewGeneratedReportDAO,
	service.NewUser,
	service.NewHealingLogService,
	service.NewImageService,
	service.NewOtherService,
	service.NewAIReportService,
	controller.NewUserController,
	controller.NewHealingLogController,
	controller.NewChildArchiveController,
	controller.NewAIReportController,
	NewJwtClient,
	NewEngine,
	wire.Struct(new(App), "Engine"),
	wire.Bind(new(service.UserService), new(*service.User)),
)

func NewJwtClient() *middleware.JwtClient {
	return &middleware.JwtClient{SecretKey: config.GetJWTConfig().SecretKey}
}

func NewEngine(
	userController *controller.User,
	healingLogController *controller.HealingLogController,
	childArchiveController *controller.ChildArchiveController,
	aiReportController *controller.AIReportController,
	jwtClient *middleware.JwtClient,
) *gin.Engine {
	r := gin.Default()
	
	// 设置用户路由
	routes.SetupUserRoutes(r, userController, jwtClient)
	
	// 设置疗愈日志路由
	routes.SetupHealingLogRoutes(r, healingLogController, jwtClient)
	
	// 设置儿童档案路由
	routes.SetupChildArchiveRoutes(r, childArchiveController, jwtClient)
	
	// 设置AI报告路由
	routes.SetupAIReportRoutes(r, aiReportController, jwtClient)
	
	return r
}

func InitializeApp() (*App, error) {
	wire.Build(ProviderSet)
	return &App{}, nil
}
//go:build wireinject
// +build wireinject

package main

import (
	"melody_cure/DAO"
	"melody_cure/controller"
	"melody_cure/middleware"
	"melody_cure/routes"
	"melody_cure/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

type App struct {
	Engine *gin.Engine
}

func ProvideJwtClient() *middleware.JwtClient {
	return &middleware.JwtClient{SecretKey: viper.GetString("JWT.secretKey")}
}

func ProvideUserService(jwt *middleware.JwtClient) service.UserService {
	return service.NewUser(DAO.NewUserDAO(DAO.DB), jwt)
}

func ProvideUserController(svc service.UserService) *controller.User {
	return controller.NewUserController(svc)
}

func ProvideImageService() *service.ImageService {
	return service.NewImageService()
}

func ProvideImageController(imageService *service.ImageService) *controller.ImageController {
	return controller.NewImageController(imageService)
}

func ProvideHealingLogDAO() *DAO.HealingLogDAO {
	return DAO.NewHealingLogDAO(DAO.DB)
}

func ProvideHealingLogService(healingLogDAO *DAO.HealingLogDAO) *service.HealingLogService {
	return service.NewHealingLogService(healingLogDAO)
}

func ProvideHealingLogController(healingLogService *service.HealingLogService) *controller.HealingLogController {
	return controller.NewHealingLogController(healingLogService)
}

func ProvideEngine(uc *controller.User, ic *controller.ImageController, hlc *controller.HealingLogController, jwt *middleware.JwtClient) *gin.Engine {
	r := gin.Default()

	// 设置用户路由
	routes.SetupUserRoutes(r, uc, jwt)
	
	// 设置图床路由
	routes.SetupImageRoutes(r, ic, jwt)

	// 设置疗愈日志路由
	routes.SetupHealingLogRoutes(r, hlc, jwt)

	return r
}

func InitializeApp() (*App, error) {
	wire.Build(
		ProvideJwtClient,
		ProvideUserService,
		ProvideUserController,
		ProvideImageService,
		ProvideImageController,
		ProvideHealingLogDAO,
		ProvideHealingLogService,
		ProvideHealingLogController,
		ProvideEngine,
		wire.Struct(new(App), "Engine"),
	)
	return &App{}, nil
}
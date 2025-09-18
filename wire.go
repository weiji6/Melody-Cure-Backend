//go:build wireinject
// +build wireinject

package main

import (
	"melody_cure/DAO"
	"melody_cure/controller"
	"melody_cure/middleware"
	"melody_cure/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

type App struct {
	Engine *gin.Engine
}

func ProvideJwtClient() middleware.JwtClient {
	return middleware.JwtClient{SecretKey: viper.GetString("JWT.secretKey")}
}

func ProvideUserService(jwt middleware.JwtClient) service.UserService {
	return service.NewUser(DAO.User{}, jwt)
}

func ProvideUserController(svc service.UserService) *controller.User {
	return controller.NewUserController(svc)
}

func ProvideEngine(uc *controller.User) *gin.Engine {
	r := gin.Default()

	// 用户登录注册路由
	r.POST("/api/user/register", uc.Register)
	r.POST("/api/user/login", uc.Login)

	return r
}

func InitializeApp() (*App, error) {
	wire.Build(
		ProvideJwtClient,
		ProvideUserService,
		ProvideUserController,
		ProvideEngine,
		wire.Struct(new(App), "Engine"),
	)
	return &App{}, nil
}
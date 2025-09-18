package main

import (
	"fmt"
	"melody_cure/DAO"
	"melody_cure/config"
)

func main() {

	config.InitConfig()

	DAO.NewData()

	app, err := InitializeApp()
	if err != nil {
		panic(fmt.Sprintf("初始化应用失败: %v", err))
	}

	if err := app.Engine.Run(":8080"); err != nil {
		panic(fmt.Sprintf("服务启动失败: %v", err))
	}
}

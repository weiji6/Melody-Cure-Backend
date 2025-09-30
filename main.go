package main

import (
	"fmt"
	"melody_cure/DAO"
	"melody_cure/config"
)

func main() {
	config.InitConfig()

	// 初始化数据库连接和全局变量
	db, err := DAO.NewDB()
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}
	DAO.DB = db

	// 自动迁移数据库表
	if err := DAO.AutoMigrate(db); err != nil {
		panic("数据库迁移失败: " + err.Error())
	}

	// 初始化Redis连接
	DAO.RDB = DAO.NewRedis()

	app, err := InitializeApp()
	if err != nil {
		panic(fmt.Sprintf("初始化应用失败: %v", err))
	}

	if err := app.Engine.Run(":8080"); err != nil {
		panic(fmt.Sprintf("服务启动失败: %v", err))
	}
}

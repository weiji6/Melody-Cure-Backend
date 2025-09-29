package DAO

import (
	"context"
	"fmt"
	"melody_cure/config"
	"melody_cure/model"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 全局实例，供其他包直接使用
var (
	DB  *gorm.DB
	RDB *redis.Client
)

func NewData() {
	config.InitConfig()

	db, err := NewDB()
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}
	DB = db

	// 自动迁移数据库表
	if err := AutoMigrate(db); err != nil {
		panic("数据库迁移失败: " + err.Error())
	}

	RDB = NewRedis()
}

// AutoMigrate 自动迁移数据库表
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&Certification{},
		&AICompanion{},
		&VirtualTherapist{},
		&ChildArchive{},
		&UserFavorite{},
		&Course{},
		&Game{},
		&model.ImageToken{}, // 图床token表
		&model.HealingLog{}, // 疗愈日志表
		&model.LogMedia{},   // 日志媒体表
	)
}

func NewDB() (*gorm.DB, error) {
	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	user := viper.GetString("db.username")
	pass := viper.GetString("db.password")
	name := viper.GetString("db.dbname")
	charset := viper.GetString("db.charset")
	parseTime := viper.GetString("db.parseTime")
	loc := viper.GetString("db.loc")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		user, pass, host, port, name, charset, parseTime, loc,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	return db, nil
}

func NewRedis() *redis.Client {
	addr := fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port"))

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})

	// 启动时校验连接
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil {
		panic("Redis 连接失败: " + err.Error())
	}

	return rdb
}

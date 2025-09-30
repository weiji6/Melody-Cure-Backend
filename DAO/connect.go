package DAO

import (
	"context"
	"fmt"
	"melody_cure/config"
	"melody_cure/model"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 全局实例，供其他包直接使用
var (
	DB  *gorm.DB
	RDB *redis.Client
)

func NewData() (*gorm.DB, *redis.Client) {
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

	rdb := NewRedis()
	RDB = rdb
	
	return db, rdb
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
		&model.HealingLog{},
		&model.LogMedia{},
		&model.ImageToken{},
		&model.GeneratedReport{},
	)
}

func NewDB() (*gorm.DB, error) {
	dbConfig := config.GetDatabaseConfig()
	
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, 
		dbConfig.DBName, dbConfig.Charset, dbConfig.ParseTime, dbConfig.Loc,
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
	redisConfig := config.GetRedisConfig()
	
	addr := fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	// 启动时校验连接
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil {
		panic("Redis 连接失败: " + err.Error())
	}

	return rdb
}

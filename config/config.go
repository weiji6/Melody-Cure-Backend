package config

import "github.com/spf13/viper"

type Config struct {
	Database DatabaseConfig
	redis    RedisConfig
}

type DatabaseConfig struct {
	Host      string
	Port      string
	Username  string
	Password  string
	DBName    string
	Charset   string
	ParseTime string
	Loc       string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type JWTConfig struct {
	SecretKey string
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		panic("读取配置失败：" + err.Error())
	}
}

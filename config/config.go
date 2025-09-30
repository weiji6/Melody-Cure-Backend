package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig
	Redis    RedisConfig
	JWT      JWTConfig
	Email    EmailConfig
	Qiniu    QiniuConfig
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

type EmailConfig struct {
	Email string
	Key   string
}

type QiniuConfig struct {
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
	Bucket    string `mapstructure:"bucket"`
	Domain    string `mapstructure:"domain"`
	Zone      string `mapstructure:"zone"`
	UseHTTPS  bool   `mapstructure:"use_https"`
	Expires   int64  `mapstructure:"expires"`
}

var GlobalConfig Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	
	// 支持环境变量覆盖配置文件
	viper.AutomaticEnv()
	
	// 设置环境变量前缀
	viper.SetEnvPrefix("MELODY")
	
	// 环境变量映射
	viper.BindEnv("db.host", "DB_HOST")
	viper.BindEnv("db.port", "DB_PORT")
	viper.BindEnv("db.username", "DB_USER")
	viper.BindEnv("db.password", "DB_PASSWORD")
	viper.BindEnv("db.dbname", "DB_NAME")
	
	viper.BindEnv("redis.host", "REDIS_HOST")
	viper.BindEnv("redis.port", "REDIS_PORT")
	viper.BindEnv("redis.password", "REDIS_PASSWORD")
	viper.BindEnv("redis.db", "REDIS_DB")
	
	viper.BindEnv("jwt.secretKey", "JWT_SECRET")
	
	viper.BindEnv("qiniu.access_key", "QINIU_ACCESS_KEY")
	viper.BindEnv("qiniu.secret_key", "QINIU_SECRET_KEY")
	viper.BindEnv("qiniu.bucket", "QINIU_BUCKET")
	viper.BindEnv("qiniu.domain", "QINIU_DOMAIN")
	
	viper.BindEnv("email.email", "EMAIL_USERNAME")
	viper.BindEnv("email.key", "EMAIL_PASSWORD")
	
	// 设置默认值
	setDefaults()
	
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		// 如果配置文件不存在，检查是否有环境变量
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件不存在，使用环境变量
			if os.Getenv("DB_HOST") == "" {
				panic("配置文件不存在且未设置环境变量")
			}
		} else {
			panic("读取配置失败：" + err.Error())
		}
	}
	
	// 将配置解析到结构体
	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		panic("配置解析失败：" + err.Error())
	}
}

func setDefaults() {
	// 数据库默认配置
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "3306")
	viper.SetDefault("db.username", "melody_cure")
	viper.SetDefault("db.dbname", "melody_cure")
	viper.SetDefault("db.charset", "utf8mb4")
	viper.SetDefault("db.parseTime", "True")
	viper.SetDefault("db.loc", "Local")
	
	// Redis默认配置
	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", "6379")
	viper.SetDefault("redis.db", 0)
	
	// 七牛云默认配置
	viper.SetDefault("qiniu.zone", "Zone_z0")
	viper.SetDefault("qiniu.use_https", true)
	viper.SetDefault("qiniu.expires", 3600)
	
	// 邮箱默认配置
	viper.SetDefault("email.host", "smtp.gmail.com")
	viper.SetDefault("email.port", 587)
}

// GetConfig 获取全局配置
func GetConfig() Config {
	return GlobalConfig
}

// GetDatabaseConfig 获取数据库配置
func GetDatabaseConfig() DatabaseConfig {
	return GlobalConfig.Database
}

// GetRedisConfig 获取Redis配置
func GetRedisConfig() RedisConfig {
	return GlobalConfig.Redis
}

// GetJWTConfig 获取JWT配置
func GetJWTConfig() JWTConfig {
	return GlobalConfig.JWT
}

// GetEmailConfig 获取邮箱配置
func GetEmailConfig() EmailConfig {
	return GlobalConfig.Email
}

// GetQiniuConfig 获取七牛云配置
func GetQiniuConfig() QiniuConfig {
	return GlobalConfig.Qiniu
}

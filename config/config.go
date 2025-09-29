package config

import "github.com/spf13/viper"

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

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		panic("读取配置失败：" + err.Error())
	}
}

// GetQiniuConfig 获取七牛云配置
func GetQiniuConfig() QiniuConfig {
	return QiniuConfig{
		AccessKey: viper.GetString("qiniu.access_key"),
		SecretKey: viper.GetString("qiniu.secret_key"),
		Bucket:    viper.GetString("qiniu.bucket"),
		Domain:    viper.GetString("qiniu.domain"),
		Zone:      viper.GetString("qiniu.zone"),
		UseHTTPS:  viper.GetBool("qiniu.use_https"),
		Expires:   viper.GetInt64("qiniu.expires"),
	}
}

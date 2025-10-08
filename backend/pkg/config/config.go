package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	JWT      JWTConfig
	Log      LogConfig
	RabbitMQ RabbitMQConfig
	Wechat   WechatConfig
	Alipay   AlipayConfig
	OSS      OSSConfig
	Hardware HardwareConfig
}

type ServerConfig struct {
	Port int
	Mode string
}

type DatabaseConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	DBName          string
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
	PoolSize int
}

type JWTConfig struct {
	Secret             string
	ExpireHours        int
	RefreshExpireHours int
}

type LogConfig struct {
	Level    string
	FilePath string
}

type RabbitMQConfig struct {
	URL      string
	Exchange string
	Queues   struct {
		Order     string
		Payment   string
		Inventory string
	}
}

type WechatConfig struct {
	AppID     string
	AppSecret string
	MchID     string
	APIKey    string
	NotifyURL string
}

type AlipayConfig struct {
	AppID      string
	PrivateKey string
	PublicKey  string
	NotifyURL  string
}

type OSSConfig struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	Bucket    string
}

type HardwareConfig struct {
	Hikvision struct {
		Host     string
		Port     int
		Username string
		Password string
	}
	ESL struct {
		APIURL string
		APIKey string
	}
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	// 设置默认值
	setDefaults()

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// 读取环境变量覆盖配置
	loadEnvVars()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func setDefaults() {
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.mode", "debug")
	viper.SetDefault("database.max_open_conns", 100)
	viper.SetDefault("database.max_idle_conns", 10)
	viper.SetDefault("database.conn_max_lifetime", 3600)
	viper.SetDefault("redis.pool_size", 100)
	viper.SetDefault("jwt.expire_hours", 24)
	viper.SetDefault("jwt.refresh_expire_hours", 168)
	viper.SetDefault("log.level", "debug")
}

func loadEnvVars() {
	// 数据库配置
	if host := os.Getenv("DATABASE_HOST"); host != "" {
		viper.Set("database.host", host)
	}
	if port := os.Getenv("DATABASE_PORT"); port != "" {
		viper.Set("database.port", port)
	}
	if user := os.Getenv("DATABASE_USER"); user != "" {
		viper.Set("database.user", user)
	}
	if password := os.Getenv("DATABASE_PASSWORD"); password != "" {
		viper.Set("database.password", password)
	}
	if dbname := os.Getenv("DATABASE_NAME"); dbname != "" {
		viper.Set("database.dbname", dbname)
	}

	// Redis配置
	if host := os.Getenv("REDIS_HOST"); host != "" {
		viper.Set("redis.host", host)
	}
	if port := os.Getenv("REDIS_PORT"); port != "" {
		viper.Set("redis.port", port)
	}

	// RabbitMQ配置
	if url := os.Getenv("RABBITMQ_URL"); url != "" {
		viper.Set("rabbitmq.url", url)
	}

	// 服务器配置
	if port := os.Getenv("SERVER_PORT"); port != "" {
		viper.Set("server.port", port)
	}
	if mode := os.Getenv("SERVER_MODE"); mode != "" {
		viper.Set("server.mode", mode)
	}

	// JWT配置
	if secret := os.Getenv("JWT_SECRET"); secret != "" {
		viper.Set("jwt.secret", secret)
	}
}

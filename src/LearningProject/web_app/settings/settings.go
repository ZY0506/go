package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// APPConfig 应用配置结构体
type APPConfig struct {
	APP    *APP    `mapstructure:"app"`
	Logger *Logger `mapstructure:"log"`
	MySQL  *MySQL  `mapstructure:"mysql"`
	Redis  *Redis  `mapstructure:"redis"`
}

// APP 应用相关配置
type APP struct {
	Name      string `mapstructure:"name"`
	Port      int    `mapstructure:"port"`
	Mode      string `mapstructure:"mode"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
}

// Logger 日志相关配置
type Logger struct {
	Level      string `mapstructure:"level"`
	Output     string `mapstructure:"output"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`     // 最大保存天数
	MaxBackups int    `mapstructure:"max_backups"` // 最多文件个数
	Compress   bool   `mapstructure:"compress"`    // 是否压缩
}

// MySQL MySQL数据库相关配置
type MySQL struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"` // 最大连接数
	MaxIdleConns int    `mapstructure:"max_idle_conns"` // 最大空闲连接数
}

// Redis Redis相关配置
type Redis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

// Config 全局配置实例
var Config = new(APPConfig)

// Init 初始化配置
func Init() error {
	// 设置配置文件路径
	viper.SetConfigFile("./settings/config.yaml")
	viper.AddConfigPath("./")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	// 解析配置文件到结构体
	if err := viper.Unmarshal(Config); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file change:")
		if err := viper.Unmarshal(Config); err != nil {
			fmt.Printf("unmarshal config failed, err: %v\n", err)
		}
	})

	return nil
}

package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {	// 在序列化和反序列化的时候，结构体的tag必须是mapstructure，而不是yaml、json。。。
	Host        string      `mapstructure:"host"`
	Version     string      `mapstructure:"version"`
	MySQLConfig MySQLConfig `mapstructure:"mysql"`
}
type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

func main() {
	viper.SetDefault("fileDir", "./")
	// 读取配置文件
	viper.SetConfigFile("./config.yaml") // 指定配置文件路径
	viper.SetConfigName("config")        // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")          // 如果配置文件的名称中没有扩展名，则需要配置此项
	// viper.SetConfigFile("./config.yaml")
	viper.AddConfigPath("/etc/appname/")  // 查找配置文件所在的路径
	viper.AddConfigPath("$HOME/.appname") // 多次调用以添加多个搜索路径
	viper.AddConfigPath(".")              // 还可以在工作目录中查找配置
	err := viper.ReadInConfig()           // 查找并读取配置文件
	if err != nil {                       // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) { // 回调函数
		fmt.Println("config file changed:", e.Name)
	})

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		fmt.Printf("unmarshal err: %v\n", err)
		return
	} else {
		fmt.Printf("%#v\n", c)
	}

	// r := gin.Default()
	// r.GET("/version", func(c *gin.Context) {
	// 	c.String(http.StatusOK, viper.GetString("version"))
	// })
	// r.Run(":8080")
}

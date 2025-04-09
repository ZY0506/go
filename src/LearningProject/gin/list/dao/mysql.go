package dao

import (
	"fmt"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 数据库信息结构体
type MysqlConfig struct {
	Username string `ini:"username"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	DB       string `ini:"db"`
}

// 加载配置文件

func LoadMySQLConfig() *MysqlConfig {
	mysqlconfig := new(MysqlConfig)
	// 加载配置文件
	err := ini.MapTo(mysqlconfig, "./config/mysql.ini")
	if err != nil {
		fmt.Println("加载配置文件失败")
		panic(err)
	}
	fmt.Printf("%#v", mysqlconfig)
	return mysqlconfig
}

// 连接数据库
func Connect() (*gorm.DB, error) {
	mysqlconfig := LoadMySQLConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlconfig.Username, mysqlconfig.Password, mysqlconfig.Host, mysqlconfig.Port, mysqlconfig.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true, //关闭小写转换
		},
	})
	return db, err
}

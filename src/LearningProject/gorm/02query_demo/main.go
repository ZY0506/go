package main

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB    *gorm.DB
	sqlDB *sql.DB
)

func init() {
	username := "root"
	password := "zy20050608"
	host := "localhost"
	port := 3306
	dbname := "golearningtest"
	timeout := "10s"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, dbname, timeout)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// SingularTable: true, //禁用复数=>表的名字不会是复数
			NoLowerCase: true, //关闭小写转换
		},
	})
	if err != nil {
		panic(err)
	}
	DB = db
	sqlDB, err = DB.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
}

func main() {

}

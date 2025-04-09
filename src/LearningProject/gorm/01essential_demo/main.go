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

type Student struct {
	ID   uint   `gorm:"type:int(11) unsigned;autoIncrement:true;primaryKey"`
	Name string `gorm:"type:varchar(255);not null;default:'小王子'"`
	Age  sql.NullInt16    `gorm:"type:int(11) unsigned;size:3;default:0"`
}

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
			NoLowerCase:   true, //关闭小写转换
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
	// // 创建表
	// DB.AutoMigrate(&Student{})
	// //创建数据行
	// stu1 := Student{
	// 	Name: "章鱼小王子",
	// 	Age:  18,
	// }
	// DB.Create(&stu1)
	DB.Create(&Student{
		// Name: "大王子",
		Age: sql.NullInt16{Int16:18,Valid: true},	
	})
	DB.Create(&Student{
		Age:sql.NullInt16{Int16:0,Valid: false},
	})
	//查询第一行的数据
	var stu Student
	DB.First(&stu)
	fmt.Printf("%#v", stu)

	// //更新数据
	// var stu Student
	// DB.Take(&stu,"id=?",3)
	// stu.Age = 20
	// DB.Save(&stu)
	// DB.Model(&Student{}).Where("age=?",18).Update("age", 20)

	// var studentList []Student
	// DB.Find(&studentList, "age = ?", 20).Update("age", "18")

	//删除
	// DB.Delete(&stu)

	if err := sqlDB.Close(); err != nil {
		panic(err)
	}
}

package main

//sqlx demo
import (
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" //只调用 init()
	"github.com/jmoiron/sqlx"
)

// 配置信息结构体
type DBConfig struct {
	DBUsername string `json:"MySQL_USERNAME"`
	DBPassword string `json:"MySQL_PASSWORD"`
	DBIP       string `json:"MySQL_IP"`
	DBPort     string `json:"MySQL_PORT"`
	DBName     string
}

//Go连接MySQL数据库

var (
	db *sqlx.DB //数据库的连接池
)

type user struct {
	Id   int
	Name string
	Age  int
}

// 初始化
func initDB() (err error) {
	//数据库信息
	ret, err := os.ReadFile("MySQL.json")
	if err != nil {
		//fmt.Printf("read config failed,err:%v\n", err)
		return
	}
	var dbconfig DBConfig
	err = json.Unmarshal(ret, &dbconfig)
	if err != nil {
		//fmt.Printf("Unmasrshal failed,err%v\n", err)
		return
	}
	dbconfig.DBName = "golearningtest"
	//用户名:密码@tcp(IP:3306)/databaseName
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbconfig.DBUsername, dbconfig.DBPassword, dbconfig.DBIP, dbconfig.DBPort, dbconfig.DBName)
	fmt.Println(dsn)
	//连接数据库
	db, err = sqlx.Connect("mysql", dsn) //不会校验用户和密码
	if err != nil {                      //dsn格式错误的时候会报错
		//fmt.Printf("dsn: %s invalid,err:%v\n", dsn, err)
		return
	}
	err = db.Ping() //尝试连接数据库
	if err != nil {
		//fmt.Printf("open %s failed,err:%v\n", dsn, err)
		return
	}
	//fmt.Println("连接数据库成功!")
	db.SetMaxOpenConns(10) //设置数据库连接池数量,请求连接多于这个值，会一直等待
	db.SetMaxIdleConns(5)  //设置最大的闲置连接数
	return
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("initDB failed,err:", err)
		return
	}
	defer db.Close()
	sqlStr := "select name,age from user where id=1"
	var u user
	db.Get(&u, sqlStr)
	fmt.Printf("u:%#v\n", u)

	var userList = make([]user, 0, 10)
	sqlStr2 := "select id,name,age from user"
	err = db.Select(&userList, sqlStr2)
	if err != nil {
		fmt.Println("select failed,err:", err)
		return
	}
	fmt.Printf("userList:%#v\n", userList)
}

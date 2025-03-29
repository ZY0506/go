package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" //只调用 init()
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
	db *sql.DB //数据库的连接池
)

type user struct {
	id   int
	name string
	age  int
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
	db, err = sql.Open("mysql", dsn) //不会校验用户和密码
	if err != nil {                  //dsn格式错误的时候会报错
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

func queryrow() {
	//查询单条记录

	//1、查询单条记录的sql语句
	sqlStr := "select id,name,age from user where id=?" //?是占位符，QueryRow的第二个参数就是值
	//2、执行
	rowObj := db.QueryRow(sqlStr, 2) //从连接池里拿一个出来去数据库查询单条记录
	//3、拿到结果
	var u1 user                            //结构体数据是值类型，必须传指针
	rowObj.Scan(&u1.id, &u1.name, &u1.age) //rowObj必须调用Scan方法，可以自动释放数据库的连接
	//打印结果
	fmt.Printf("u1:%#v\n", u1)
}

func querymore(n int) {
	//查询多条记录

	//1、SQL语句
	sqlStr := "select id,name,age from user where id<?"
	//2、执行
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exec %s query failed,err:%v\n", sqlStr, err)
		return
	}
	//3、一定要关闭rows
	defer rows.Close()
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
		}
		fmt.Printf("u1:%#v\n", u1)
	}
}

// 插入
func insert() {
	//sql语句
	sqlStr := `insert into user(name,age) values("李四",28)`
	//exec
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	//如果是插入数据操作，能够拿到插入数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Println("id:", id)
}

// 更新操作
func update(age, id int) {
	//sql语句
	sqlStr := `update user set age=? where id=?`
	ret, err := db.Exec(sqlStr, age, id)
	if err != nil {
		fmt.Printf("update failed,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get rowsfailed,err:%v\n", err)
		return
	}
	fmt.Printf("更新了%d行\n", n)
}

// 删除操作
func delete(id int) {
	//sql语句
	sqlStr := `delete from user where id=?`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get rowsfailed,err:%v\n", err)
		return
	}
	fmt.Printf("删除了%d行\n", n)
}

// 预处理实现多条插入
func prepareInsert() {
	sqlStr := `insert into user(name,age) values(?,?)`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed,err:%v\n", err)
		return
	}
	//关闭预处理语句
	defer stmt.Close()
	var m = map[string]int{
		"张三": 18,
		"李四": 28,
		"王五": 38,
		"赵六": 48,
		"钱七": 58,
	}
	for name, age := range m {
		stmt.Exec(name, age)
	}
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("initDB failed,err:%v\n", err)
	} else {
		fmt.Println("连接数据库成功！")
	}
	defer db.Close()
	//查询单条记录
	queryrow()

	//查询多条记录
	//querymore(3)

	//插入
	//insert()

	//更新
	// update(29,3)

	//删除
	// delete(3)

	//预处理
	prepareInsert()
}

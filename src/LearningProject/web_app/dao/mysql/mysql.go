package mysql

import (
	"LearningProject/web_app/settings"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init(mysqlconfig *settings.MySQL) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		// 读取数据库信息
		mysqlconfig.User,
		mysqlconfig.Password,
		mysqlconfig.Host,
		mysqlconfig.Port,
		mysqlconfig.DBName,
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("Connect DB failed", zap.Error(err))
		return
	}
	// 设置数据库最大连接数以及最大空闲连接数
	db.SetMaxOpenConns(mysqlconfig.MaxOpenConns)
	db.SetMaxIdleConns(mysqlconfig.MaxIdleConns)
	return
}

func Close() {
	_ = db.Close()
}

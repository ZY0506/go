package mysql

import (
	"LearningProject/web_app/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

var secret = "ZhangYuXiaoWangZi"

var (
	ErrorUserExist       = errors.New("用户名已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)

// CheckUserExist 判断用户是否存在
func CheckUserExist(username string) error {
	// 执行sql语句
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return nil
}

// InsertUser 向数据库中插入一条用户数据
func InsertUser(user *models.User) (err error) {
	// 密码加密
	user.Password = encryptPassword(user.Password)
	// 执行sql语句
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id,username,password from user where username = ?`
	err = db.Get(user, sqlStr, user.Username)
	// 用户不存在
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	// 查询错误
	if err != nil {
		return
	}
	// 判断密码是否正确
	if encryptPassword(oPassword) != user.Password {
		return ErrorInvalidPassword
	}
	return
}

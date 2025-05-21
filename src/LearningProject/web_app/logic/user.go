package logic

import (
	"LearningProject/web_app/dao/mysql"
	"LearningProject/web_app/models"
	"LearningProject/web_app/pkg/jwt"
	"LearningProject/web_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 1、检查用户是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 2、生成UID
	userID := snowflake.GenID()
	// 构造一个用户实例
	u := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 3、保存到数据库
	return mysql.InsertUser(u)
}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err = mysql.Login(user); err != nil {
		return "", err
	}
	// 生成jwt
	return jwt.GenToken(user.UserID, user.Username)
}

package controller

import (
	"LearningProject/web_app/dao/mysql"
	"LearningProject/web_app/logic"
	"LearningProject/web_app/models"
	"LearningProject/web_app/pkg/jwt"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strings"
)

// SignUpHandler 注册
func SignUpHandler(c *gin.Context) {
	// 1、获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断是不是validator.ValidationError类型
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2、业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		// 用户存在
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		// 数据库或者其他错误，返回系统繁忙
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3、返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 登录
func LoginHandler(c *gin.Context) {
	// 1、获取参数和参数校验
	var p = new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断是不是validator.ValidationError类型
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs) // 这种错误一般是参数绑定时出现的错误
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2、业务逻辑处理
	aToken, rToken, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		ResponseError(c, CodeServerBusy)
	}
	// 3、返回响应
	ResponseSuccess(c, gin.H{
		"accessToken":  aToken,
		"refreshToken": rToken,
		"username":     p.Username,
	})
}

// RefreshTokenHandler 刷新token
func RefreshTokenHandler(c *gin.Context) {
	rt := c.Query("refresh_token")
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		ResponseErrorWithMsg(c, CodeInvalidToken, "请求头缺少Auth Token")
		c.Abort()
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		ResponseErrorWithMsg(c, CodeInvalidToken, "请求头中Auth格式有误")
		c.Abort()
		return
	}
	aToken, rToken, err := jwt.RefreshToken(parts[1], rt)
	fmt.Println(err)
	ResponseSuccess(c, gin.H{
		"access_token":  aToken,
		"refresh_token": rToken,
	})
}

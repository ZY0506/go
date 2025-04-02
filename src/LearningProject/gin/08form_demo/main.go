package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//获取form表单提交的参数

func main() {
	r := gin.Default()
	//加载模板文件
	r.LoadHTMLFiles("login.html", "index.html")
	r.GET("/login", func(ctx *gin.Context) { //get请求
		ctx.HTML(http.StatusOK, "login.html", nil) //返回响应
	})
	//login post 登录时发送post请求
	r.POST("/login", func(ctx *gin.Context) {
		// username := ctx.PostForm("username")
		// password := ctx.PostForm("password")
		username := ctx.DefaultPostForm("username", "somebody")
		password := ctx.DefaultPostForm("password", "***")
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})
	r.Run(":9090")
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//querystring

func main() {
	r := gin.Default()
	//GET请求
	r.GET("/web", func(ctx *gin.Context) {
		//获取浏览器那边发请求携带的 querystring 参数
		name := ctx.Query("query") //通过Query方法获取参数
		age := ctx.Query("age")
		// name:=ctx.DefaultQuery("query","somebody")	//取不到就用指定的默认值
		// name,ok:=ctx.GetQuery("query")
		// if!ok{
		// 	name="somebody"	//取不到就用指定的默认值
		// }
		ctx.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.Run(":9090")
}

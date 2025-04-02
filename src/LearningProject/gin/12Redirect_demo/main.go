package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/index1", func(ctx *gin.Context) {
		// ctx.JSON(http.StatusOK, gin.H{
		// 	"status": "ok",
		// })
		ctx.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	r.GET("/a", func(ctx *gin.Context) {
		//跳转到 /b 对应的路由处理函数
		ctx.Request.URL.Path = "/b" //把请求的URL修改
		r.HandleContext(ctx)        //继续处理这个请求
	})
	r.GET("/b", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "b",
		})
	})
	r.Run(":9090")
}

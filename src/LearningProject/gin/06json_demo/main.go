package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//方法一：map
	r.GET("/json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"name":    "小王子",
			"message": "hello world",
			"age":     18,
		})
	})
	//方法二：结构体
	type msg struct {
		Name    string `json:"name"`
		Message string `json:"message"`
		Age     int    `json:"age"`
	}
	r.GET("/json2", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, msg{Name: "小王子", Message: "hello world", Age: 18})
	})
	r.Run(":9090")
}

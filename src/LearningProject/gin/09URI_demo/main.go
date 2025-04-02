package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.GET("/:name/:age", func(ctx *gin.Context) {	//这样写，和下面在同一份代码中并不安全
	r.GET("user/:name/:age",func(ctx *gin.Context) {	//安全写法
		//获取路径参数
		name := ctx.Param("name")
		age := ctx.Param("age")
		ctx.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.GET("/blog/:year/:month",func(ctx *gin.Context){
		year := ctx.Param("year")
		month := ctx.Param("month")
		ctx.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
		})
	})
	r.Run(":9090")
}

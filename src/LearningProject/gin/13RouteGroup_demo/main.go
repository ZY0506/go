package main

//路由组

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.GET("/video/index",func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK,gin.H{"msg":"/video/index"})
	// })
	// r.GET("/video/xx",func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK,gin.H{"msg":"/video/xx"})
	// })
	// r.GET("/video/oo",func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK,gin.H{"msg":"/video/oo"})
	// })
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
		})
		videoGroup.GET("/xx", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "/video/xx"})
		})
		videoGroup.GET("/oo", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "/video/oo"})
		})
	}
	r.Run(":9090")
}

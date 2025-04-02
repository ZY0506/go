package main

//中间件

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	fmt.Println("indexHandler in......")
	value, exists := c.Get("name")
	if !exists {
		value = "匿名用户"
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "index page",
		"name":    value,
	})
}
func m1(c *gin.Context) {
	fmt.Println("m1 in......")
	start := time.Now()
	c.Next() //调用后续的处理函数
	// c.Abort()///终止后续的处理函数
	fmt.Println("m1 out......")
	cost := time.Since(start)
	fmt.Println("indexHandler cost:", cost)
}

func m2(c *gin.Context) {
	fmt.Println("m2 in......")
	c.Set("name", "章鱼小王子")
	c.Next() //调用后续的处理函数
	fmt.Println("m2 out......")
}

// 检查登录 一般的中间件
func checkLogin(doCheck bool) gin.HandlerFunc {
	/*
		可以做的其他操作：连接数据库、其他准备工作
	*/
	return func(c *gin.Context) {
		if doCheck {
			//具体检查登录的操作
			// if 是登录用户
			c.Next()
			// else 不是
			// 	c.Abort()
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default()
	r.Use(m1, m2, checkLogin(true)) //全局注册中间件
	r.GET("/index", indexHandler)

	r.GET("/shop", func(ctx *gin.Context) {
		fmt.Println("shopHandler in......")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "shop page",
		})
	})

	r.GET("/user", func(ctx *gin.Context) {
		fmt.Println("userHandler in......")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "user page",
		})
	})

	// //路由组注册中间件 方法1：
	// videoGroup := r.Group("/video", checkLogin(true))
	// {
	// 	videoGroup.GET("/index", func(ctx *gin.Context) {
	// 		ctx.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
	// 	})
	// 	videoGroup.GET("/xx", func(ctx *gin.Context) {
	// 		ctx.JSON(http.StatusOK, gin.H{"msg": "/video/xx"})
	// 	})
	// 	videoGroup.GET("/oo", func(ctx *gin.Context) {
	// 		ctx.JSON(http.StatusOK, gin.H{"msg": "/video/oo"})
	// 	})
	// }
	// //路由组注册中间件 方法2：
	// videoGroup2 := r.Group("/video")
	// videoGroup2.Use(checkLogin(true))
	// {
	// 	videoGroup2.GET("/index", func(ctx *gin.Context) {
	// 		ctx.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
	// 	})
	// 	videoGroup2.GET("/xx", func(ctx *gin.Context) {
	// 		ctx.JSON(http.StatusOK, gin.H{"msg": "/video/xx"})
	// 	})
	// 	videoGroup2.GET("/oo", func(ctx *gin.Context) {
	// 		ctx.JSON(http.StatusOK, gin.H{"msg": "/video/oo"})
	// 	})
	// }

	r.Run(":9090")
}

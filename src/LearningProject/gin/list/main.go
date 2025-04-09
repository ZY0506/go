package main

import (
	"LearningProject/gin/list/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 注册路由
	r := gin.Default()
	// 设置值静态文件
	r.Static("/static", "static")
	// 加载模板
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	// 代办事项
	v1Group := r.Group("/v1")
	{
		// 添加
		v1Group.POST("/todo", controller.AddTodo)
		// 查看所有待办事项
		v1Group.GET("/todo", controller.FindAllTodos)
		// 查看某一个待办事项
		v1Group.GET("/todo/:id", controller.FindATodos)
		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	r.Run(":9090")
}

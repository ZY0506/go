package controller

import (
	"LearningProject/gin/list/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID     int    `json:"id"`
	Titile string `json:"title"`
	Status bool   `json:"status"`
}

// 添加一个待办事项
func AddTodo(ctx *gin.Context) {
	// 获取参数
	var todo Todo
	// 参数绑定
	if err := ctx.BindJSON(&todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	// 存入数据库
	// 返回响应
	if err := dao.DB.Create(&todo).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"data": todo})
	}
}

// 查找所有待办事项
func FindAllTodos(ctx *gin.Context) {
	// 在数据库中查找全部待办事项
	var todoList []Todo
	// 返回响应
	if err := dao.DB.Find(&todoList).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todoList)
	}
}

// 查找一个待办事项
func FindATodos(ctx *gin.Context) {
	// 获取参数
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "id is not found",
		})
		return
	}
	// 在数据库中查找
	var todo Todo
	// 返回响应
	if err := dao.DB.First(&todo, id).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
}

// 更新一个待办事项
func UpdateTodo(ctx *gin.Context) {
	// 获取参数
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "id is not found",
		})
		return
	}
	// 在数据库中修改
	// 返回响应
	var todo Todo
	if err := dao.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	ctx.BindJSON(&todo)
	if err := dao.DB.Save(&todo).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
}

// 删除一个待办事项
func DeleteTodo(ctx *gin.Context) {
	// 获取参数
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "id is not found",
		})
		return
	}
	// 在数据库中删除
	// 返回响应
	if err := dao.DB.Where("id = ?", id).Delete(Todo{}).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"msg": "delete success"})
	}
}

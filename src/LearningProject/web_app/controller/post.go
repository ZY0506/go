package controller

import (
	"LearningProject/web_app/logic"
	"LearningProject/web_app/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CreatePostHandler 创建帖子
func CreatePostHandler(c *gin.Context) {
	// 获取参数和参数校验
	post := new(models.Post)
	if err := c.ShouldBindJSON(post); err != nil {
		zap.L().Error("create post with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 从 c 取到当前用户的authID
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	post.AuthorID = userID
	// 创建帖子
	if err := logic.CreatePost(post); err != nil {
		zap.L().Error("logic.CreatePost(post) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 获取帖子详情
func CetPostDetailHandler(c *gin.Context) {
	// 获取帖子id
	pidStr := c.Param("id")
	// 类型转换
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 根据id取出帖子数据
	data, err := logic.GetPostDetail(pid)
	if err != nil {
		zap.L().Error("logic.GetPostDetail(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回数据
	ResponseSuccess(c, data)
}

// CetPostListHandler 获取帖子列表
func CetPostListHandler(c *gin.Context) {
	page, size := getPageInfo(c)
	// 获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)
}

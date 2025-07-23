package controller

import (
	"LearningProject/web_app/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// --- 社区相关代码 ---

// CommunityHandler 获取社区列表
func CommunityHandler(c *gin.Context) {
	// 获取社区列表（community_id,community_name）
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler 获取社区详情
func CommunityDetailHandler(c *gin.Context) {
	// 获取社区id
	idStr := c.Param("id")
	// 类型转换
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("get community detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

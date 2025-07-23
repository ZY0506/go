package controller

import (
	"LearningProject/web_app/logic"
	"LearningProject/web_app/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// PostVoteHandler 投票
func PostVoteHandler(c *gin.Context) {
	// 参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("PostVoteHandler with invalid param", zap.Error(err))
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		// 翻译错误，返回响应
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 获取当前请求用户的ID
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	// 业务逻辑
	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

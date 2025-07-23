package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// ResponseError 返回错误信息
func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

// ResponseErrorWithMsg 带错误信息的返回
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// ResponseSuccess 返回成功信息
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}

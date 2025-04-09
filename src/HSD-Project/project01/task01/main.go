package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 计算正方体、长方体的体积
// 正方体请求结构体
type CubeRequest struct {
	Num float64 `json:"num"`
}

// 长方体请求结构体
type CuboidRequest struct {
	Length float64 `json:"Length"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

// 响应结构体
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 计算正方体体积
func calculateCubeVolume(c *gin.Context) {
	var request CubeRequest
	var response Response
	// 绑定请求参数
	if err := c.ShouldBind(&request); err != nil {
		response = Response{
			Code: 0,
			Msg:  "请求参数错误：" + err.Error(),
			Data: nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	} else if request.Num < 0 { //参数必须大于等于0
		response = Response{
			Code: 0,
			Msg:  "请求参数错误：边长必须大于等于0",
			Data: nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response = Response{
		Code: 1,
		Msg:  "success",
		Data: request.Num * request.Num * request.Num,
	}
	c.JSON(http.StatusOK, response) //返回响应
}

// 计算长方体体积
func calculateCuboidVolume(c *gin.Context) {
	var request CuboidRequest
	var response Response
	if err := c.ShouldBind(&request); err != nil { //绑定请求参数
		response = Response{
			Code: 0,
			Msg:  "请求参数错误：" + err.Error(),
			Data: nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	} else if request.Length < 0 || request.Width < 0 || request.Height < 0 { //判断参数是否小于等于0
		response = Response{
			Code: 0,
			Msg:  "请求参数错误：长、宽、高必须大于等于0",
			Data: nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response = Response{
		Code: 1,
		Msg:  "success",
		Data: request.Length * request.Width * request.Height,
	}
	c.JSON(http.StatusOK, response) //返回响应
}
func main() {
	r := gin.Default()
	r.POST("/calculate/volume/cube", calculateCubeVolume)
	r.POST("/calculate/volume/cuboid", calculateCuboidVolume)
	r.Run(":9090")
}
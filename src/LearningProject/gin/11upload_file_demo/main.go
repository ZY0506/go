package main

import (
	"net/http"
	"path"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("index.html")
	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/upload", func(ctx *gin.Context) {
		//从请求中读取文件
		file, err := ctx.FormFile("f1")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "文件上传失败",
			})
		} else {
			//拿到文件保存到本地（服务端）
			// dst:=fmt.Sprintf("./%s",file.Filename)
			dst := path.Join("./", file.Filename)
			ctx.SaveUploadedFile(file, dst)
			ctx.JSON(http.StatusOK, gin.H{
				"status":"OK",
				"msg": "文件上传成功",
			})
		}
	})
	r.Run(":9090")
}

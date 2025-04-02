package main

//参数绑定

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func main() {
	r := gin.Default()
	r.POST("/user", func(ctx *gin.Context) {
		// username := ctx.Query("username")
		// password := ctx.Query("password")
		// u := UserInfo{
		// 	username: username,
		// 	password: password,
		// }
		var u UserInfo            //声明一个结构体变量
		err := ctx.ShouldBind(&u) //绑定，传递指针
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "参数绑定失败",
			})
		} else {
			fmt.Printf("%#v\n", u)
			ctx.JSON(http.StatusOK, gin.H{
				"message": "ok",
				"username":u.Username,
				"password":u.Password,
			})
		}
		// fmt.Printf("%#v\n", u)
		// ctx.JSON(http.StatusOK, gin.H{
		// 	"message": "ok",
		// })
	})
	r.Run(":9090")
}

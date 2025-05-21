package middlewares

import (
	"LearningProject/web_app/controller"
	"LearningProject/web_app/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设token放在Header的Authorization中，并使用Bearer开头
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的用户信息放在请求上下文中
		c.Set(controller.CtxUserIDKey, mc.UserID)
		c.Next()
	}
}

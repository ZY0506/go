package routes

import (
	"LearningProject/web_app/controller"
	"LearningProject/web_app/logger"
	"LearningProject/web_app/middlewares"
	"go.uber.org/zap"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //gin设置成发布模式
	}
	// 初始化gin框架内置的检验器使用的翻译器
	if err := controller.InitTranc("zh"); err != nil {
		zap.L().Error("init validator translator failed", zap.Error(err))
		return nil
	}
	// 注册路由
	r := gin.New()
	// 注册全局中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")
	{
		v1.POST("/signup", controller.SignUpHandler)
		v1.POST("/login", controller.LoginHandler)
		v1.GET("/refresh_token", controller.RefreshTokenHandler)
	}

	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.CetPostDetailHandler)
		v1.GET("/posts/", controller.CetPostListHandler)
		v1.POST("vote", controller.PostVoteHandler)
	}
	r.GET("ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		// 测试，如果是登录的用户，判断请求头中是否有 有效的token
		c.JSON(http.StatusOK, "pong")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}

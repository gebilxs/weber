package routes

import (
	"fmt"
	"net/http"
	"syscall"
	"time"
	"weber/controllers"
	"weber/logger"
	"weber/middlewares"
	"weber/setting"

	_ "weber/docs"

	"github.com/gin-gonic/gin"
	//"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Setup(mode string) *gin.Engine {

	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置成发布模式
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(2*time.Second, 1))

	//r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	v1 := r.Group("api/v1")
	//注册业务路由
	v1.POST("/signup", controllers.SignUpHandler)

	//注册登陆业务
	v1.POST("/login", controllers.LoginHandler)

	//应用JWT认证的中间件
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)

		v1.POST("/post", controllers.CommunityPostHandler)
		v1.GET("/post/:id", controllers.GetPostDetailHandler)
		v1.GET("/posts", controllers.GetPostListHandler)

		v1.POST("/vote", controllers.PostVoteHandler)
	}

	//登陆的用户才可以访问
	//把认证的操作放到中间件里面
	//v1.GET("/Ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
	//	//如果用户登陆了可以走这一条路由,判断请求头中是否有有效的jwt
	//	c.String(http.StatusOK, "pong")
	//	////如果用户没有登陆那么用户
	//	//c.String(http.StatusOK, "请登录")
	//})
	//手动关闭接口
	v1.GET("/end", func(c *gin.Context) {
		fmt.Println("手动调接口关闭")
		setting.Quit <- syscall.Signal(10000000)
		c.String(http.StatusOK, "OK")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}

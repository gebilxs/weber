package routes

import (
	"fmt"
	"net/http"
	"syscall"
	"weber/controllers"
	"weber/logger"
	"weber/setting"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {

	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置成发布模式
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	//注册业务路由
	r.POST("/signup", controllers.SignUpHandler)

	//注册登陆业务
	r.POST("/login", controllers.LoginHandler)

	//手动关闭接口
	r.GET("/end", func(c *gin.Context) {
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

package controllers

import (
	"net/http"
	"weber/logic"

	"github.com/gin-gonic/gin"
)

//SignUpHandler处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	//1.获取参数和参数校验

	//2.业务处理
	logic.SignUp()
	//3.返回相应
	c.JSON(http.StatusOK, "OK")
}

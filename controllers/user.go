package controllers

import (
	"fmt"
	"net/http"
	"weber/logic"
	"weber/models"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

//SignUpHandler处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	//1.获取参数和参数校验
	//use shouldBindJSON
	var p models.ParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		//如果请求参数有错误，则直接返回相应
		zap.L().Error("SignUp with invalid params", zap.Error(err)) // zap.String("xx", "vv"),
		//判断err是不是validator,validatorErrors的错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			//"msg": "请求参数有错误",
			"msg": removeTopStruct(errs.Translate(trans)), //翻译错误
		})
		return
	}
	//手动对请求参数进行判断要求返回不能为空
	//手动进行业务规则的校验
	if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.RePassword != p.Password {
		zap.L().Error("SignUp with invalid params") // zap.String("xx", "vv"),
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有错误",
		})
		return
	}
	fmt.Println(p)
	//2.业务处理
	logic.SignUp()
	//3.返回相应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

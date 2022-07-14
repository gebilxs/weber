package controllers

import (
	"net/http"
	"weber/logic"
	"weber/models"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoginHandler(c *gin.Context) {
	//获取请求参数及参数校验
	var p models.ParamLogin
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Login with invalid params", zap.Error(err))
		//判断err是不是validator,validatorErrors的错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			//use translations part
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	//业务逻辑处理
	if err := logic.Login(&p); err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户名或者密码错误",
		})
		return
	}
	//返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "登陆成功",
	})
}
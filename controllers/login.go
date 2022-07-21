package controllers

import (
	"errors"
	"fmt"
	"weber/dao/mysql"
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
			//c.JSON(http.StatusOK, gin.H{
			//	"msg": err.Error(),
			//})
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		//c.JSON(http.StatusOK, gin.H{
		//	//use translations part
		//	"msg": removeTopStruct(errs.Translate(trans)),
		//})
		return
	}
	//业务逻辑处理
	user, err := logic.Login(&p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		//可以先进行判断用户是不是存在
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": "用户名或者密码错误",
		//})
		return
	}
	//返回响应
	//c.JSON(http.StatusOK, gin.H{
	//	"msg": "登陆成功",
	//})
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID), //返回用户的UserID
		"user_name": user.Username,                  //返回用户的Username
		"token":     user.Token,                     //返回用户的token
	})
}

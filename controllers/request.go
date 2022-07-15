package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var ContextUserIDKey = "UserID"

var ErrorUserNotLogin = errors.New("用户未登录")

//getCurrentUser 获取当前登陆用户的id
func getCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

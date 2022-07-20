package controllers

import (
	"errors"
	"strconv"

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

func getOffsetInfo(c *gin.Context) (int64, int64) {
	offsetStr := c.Query("offset")
	limitStr := c.Query("limit")
	//获取分页参数
	var (
		limit  int64
		offset int64
		err    error
	)

	offset, err = strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = 1
	}
	limit, err = strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = 10
	}
	return offset, limit
}

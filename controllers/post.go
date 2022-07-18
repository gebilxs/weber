package controllers

import (
	"weber/logic"
	"weber/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func CommunityPostHandler(c *gin.Context) {
	//1.获取参数及参数的校验

	//c.ShouldBindJSON(http.StatusOK)
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.shouldBindJSON(p),error", zap.Any("err", err))
		zap.L().Error("create post with invalid params")
		ResponseError(c, CodeInvalidParam)
		return
	}
	//从c取到当前发请求的用户的ID
	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
	}
	p.AuthorID = userID
	//2.创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回相应
	ResponseSuccess(c, nil)
}

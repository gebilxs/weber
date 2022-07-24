package controllers

import (
	"strconv"
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

// GetCommunityDetails 获取帖子详情的处理函数
func GetPostDetailHandler(c *gin.Context) {
	//1.获取参数（从URL中获取帖子的id）
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//2.根据ID取出帖子数据（查数据库）
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Debug("c.GetPostById(p),error", zap.Any("err", err))
		zap.L().Error("logic.GetPostById (pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, data)
}

//GetPostListHandler 获取帖子列表的接口
func GetPostListHandler(c *gin.Context) {
	offset, limit := getOffsetInfo(c)
	//获取数据
	data, err := logic.GetPostList(offset, limit)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, data)
}

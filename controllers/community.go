package controllers

import (
	"strconv"
	"weber/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//----------跟社区相关的----------

func CommunityHandler(c *gin.Context) {
	//查询到所有的社区（community_id,community_name)以列表的形式
	data, err := logic.GetCommunityList()
	if err != nil {
		//打印日志
		zap.L().Error("logic.GetCommunity() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) //不轻易把服务端报错给外界
		return
	}
	ResponseSuccess(c, data)
}

func CommunityDetailHandler(c *gin.Context) {
	//	 1.获取社区的详情id
	idStr := c.Param("id")
	//做一个数的判断
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 获取社区的详情
	data, err := logic.GetCommunityDetails(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetails failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) //不要把服务器的错误暴露给外面
		return
	}
	ResponseSuccess(c, data)
}

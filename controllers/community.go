package controllers

import (
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

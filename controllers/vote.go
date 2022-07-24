package controllers

import (
	"weber/logic"
	"weber/models"

	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

//投票
//type VoteData struct {
//	//UserID
//	PostID    int64 `json:"post_id,string"`   //帖子id
//	Direction int   `json:"direction,string"` // 赞成票还是反对票
//}

func PostVoteHandler(c *gin.Context) {
	//参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) //进行规范化
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans)) //翻译并去错误提示中的结构体
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return

	}
	//获取当前用户的id
	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	//业务逻辑
	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, nil)
}

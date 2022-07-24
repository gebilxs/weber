package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
{
	"code":10001;//程序中出现的错误码
	"msg":xx     //提示信息
	"data":{}	 //数据
*/

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResponseError(c *gin.Context, code ResCode) {
	//rd := &ResponseData{
	//	Code: code,
	//	Msg:  code.getMsg(),
	//	Data: nil,
	//}
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.getMsg(),
		Data: nil,
	})
}
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	//rd := &ResponseData{
	//	Code: CodeSuccess,
	//	Msg:  CodeSuccess.getMsg(),
	//	Data: data,
	//}
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.getMsg(),
		Data: data,
	})
}

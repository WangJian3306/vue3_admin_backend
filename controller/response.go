package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
{
	"code": "10001", // 程序中的错误码
	"msg": "xxx",    // 提示信息
	"data": {},      // 数据
}
*/

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"message"`
	Data interface{} `json:"data"`
	OK   bool        `json:"ok"`
}

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
		OK:   false,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
		OK:   true,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
		OK:   false,
	})
}

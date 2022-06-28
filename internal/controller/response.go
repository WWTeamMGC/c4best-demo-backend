package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}
type TimeAndCount struct {
	Time  string `json:"time"`
	Count string `json:"value"`
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
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
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}

func ResponseTimeAndCount(c *gin.Context, timestr, count string) {

	c.JSON(http.StatusOK, gin.H{
		"code":  CodeSuccess,
		"msg":   CodeSuccess.Msg(),
		"time":  timestr,
		"count": count,
	})
}

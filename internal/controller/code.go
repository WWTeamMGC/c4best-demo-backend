package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExsist
	CodeUserNameShort
	CodePasswordShort
	CodeServerBusy
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:       "成功",
	CodeInvalidParam:  "参数错误",
	CodeUserExsist:    "用户已存在",
	CodeUserNameShort: "用户名太短",
	CodePasswordShort: "密码太短啦",
	CodeServerBusy:    "服务器忙",
}

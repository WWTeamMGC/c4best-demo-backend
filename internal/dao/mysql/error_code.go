package mysql

import "errors"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
	ErrorBusy            = errors.New("服务忙")
)

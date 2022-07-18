package mysql

import "errors"

var (
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("密码错误")
	ErrorUserExist       = errors.New("用户已经存在")
	ErrorInvalidID       = errors.New("无效的id")
)

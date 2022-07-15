package mysql

import (
	"database/sql"
	"errors"
	"weber/models"
)

var (
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("密码错误")
)

func Login(user *models.User) (err error) {
	oPassword := user.Password //用户登陆的密码
	sqlStr := `select user_id,username,password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserExist
	}
	if err != nil {
		//数据库查询失败
		return err
	}
	//判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}

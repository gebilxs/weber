package logic

import (
	"weber/dao/mysql"
	"weber/models"
	"weber/pkg/jwt"
)

func Login(p *models.ParamLogin) (token string, err error) {
	//登陆
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递的是指针,就能够拿到user.UserID
	if err := mysql.Login(user); err != nil {
		return " ", err
	}
	// user.UserID 生成jwt
	return jwt.GenToken(user.UserID, user.Username)
}

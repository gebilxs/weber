package logic

import (
	"weber/dao/mysql"
	"weber/models"
	"weber/pkg/jwt"
)

func Login(p *models.ParamLogin) (user *models.User, err error) {
	//登陆
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递的是指针,就能够拿到user.UserID
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	// user.UserID 生成jwt
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}

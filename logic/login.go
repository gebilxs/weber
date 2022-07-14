package logic

import (
	"weber/dao/mysql"
	"weber/models"
)

func Login(p *models.ParamLogin) (err error) {
	//登陆
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.Login(user)
}

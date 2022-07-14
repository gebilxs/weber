package logic

import (
	"weber/dao/mysql"
	"weber/models"
	"weber/pkg/snowflake"
)

//存在业务逻辑的代码
func SignUp(p *models.ParamSignUp) (err error) {
	//1.判断用户存不存在
	if err := mysql.CheckUserExists(p.Username); err != nil {
		//数据库的查询出现错误
		return err
	}
	//2.生成UID
	userID := snowflake.GenID()
	//构造一个user的实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//3.密码进行加密
	//4.保存进数据库
	return mysql.InsertUser(user)
}

package logic

import (
	"weber/dao/mysql"
	"weber/pkg/snowflake"
)

//存在业务逻辑的代码
func SignUp() {
	//1.判断用户存不存在
	mysql.QueryUserByUsername()
	//2.生成UID
	snowflake.GenID()
	//3.密码进行加密
	//4.保存进数据库
	mysql.InsertUser()
}

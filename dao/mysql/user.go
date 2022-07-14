package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"weber/models"
)

//每一步数据库操作封装成函数
//待logic层根据业务需求调用

const secret = "xckxckxck"

// CheckUserExists检查指定用户名的用户是否存在
func CheckUserExists(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	//小写说明只在dao层生效
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已经存在")
	}
	return
}

//InsertUser 往数据库中插入一条新的用户数据
func InsertUser(user *models.User) (err error) {
	//对密码进行加密
	user.Password = encryptPassword(user.Password)
	//执行SQL 语句入库
	sqlStr := `insert into user (user_id,username,password) values (?,?,?)`
	db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))

	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

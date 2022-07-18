package logic

import (
	"weber/dao/mysql"
	"weber/models"
	"weber/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	//1.生成post id
	p.ID = snowflake.GenID()
	//2.保存到数据库
	return mysql.CreatePost(p)
	//3.返回
}

package logic

import (
	"weber/dao/mysql"
	"weber/models"
)

func GetCommunityList() ([]*models.Community, error) {
	//查找数据库 查找到所有的community 并返回
	return mysql.GetCommunityList()
}
func GetCommunityDetails(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailsByID(id)
}

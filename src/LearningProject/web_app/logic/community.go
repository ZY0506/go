package logic

import (
	"LearningProject/web_app/dao/mysql"
	"LearningProject/web_app/models"
)

// GetCommunityList 获取社区列表
func GetCommunityList() (communityList []models.Community, err error) {
	return mysql.GetCommunityList()
}

// GetCommunityDetail 获取社区详情
func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}

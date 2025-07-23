package mysql

import (
	"LearningProject/web_app/models"
	"database/sql"
	"errors"
	"go.uber.org/zap"
)

// GetCommunityList 获取社区列表
func GetCommunityList() (communityList []models.Community, err error) {
	// 查询数据库community表中的所有数据
	sqlStr := "select community_id,community_name from community"
	err = db.Select(&communityList, sqlStr)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

// GetCommunityDetailByID 根据id查询社区详情
func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select community_id,community_name,introduction,create_time
			from community
			where community_id = ?
				`
	if err = db.Get(community, sqlStr, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrorInvalidID
		}
	}
	return community, err
}

package logic

import (
	"LearningProject/web_app/dao/mysql"
	"LearningProject/web_app/dao/redis"
	"LearningProject/web_app/models"
	"LearningProject/web_app/pkg/snowflake"
	"go.uber.org/zap"
)

// CreatePost 创建帖子
func CreatePost(post *models.Post) error {
	// 根据雪花算法生成postID
	post.ID = snowflake.GenID()
	// 插入数据库
	err := mysql.CreatePost(post)
	if err != nil {
		return err
	}
	err = redis.CreatePost(post.ID)
	return err
	// 返回错误
}

// GetPostDetail 根据id查询帖子详情
func GetPostDetail(pid int64) (data *models.ApiPostDetail, err error) {
	// 查询帖子数据
	post, err := mysql.GetPostByID(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(pid) failed", zap.Int64("pid", pid), zap.Error(err))
		return
	}
	// 查询作者信息
	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID(post.AuthorID) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
		return
	}
	// 查询 社区id 查询 community 信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed", zap.Int64("community_id", post.CommunityID), zap.Error(err))
		return
	}
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return data, nil
}

// GetPostList 获取帖子列表
func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		zap.L().Error("mysql.GetPostList() failed", zap.Error(err))
		return
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		user := new(models.User)
		// 查询作者信息
		user, err = mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID(post.AuthorID) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
			return
		}
		// 查询 社区id 查询 community 信息
		community := new(models.CommunityDetail)
		community, err = mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed", zap.Int64("community_id", post.CommunityID), zap.Error(err))
			return
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

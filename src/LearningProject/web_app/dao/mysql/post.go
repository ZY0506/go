package mysql

import "LearningProject/web_app/models"

// CreatePost 创建帖子
func CreatePost(post *models.Post) error {
	sqlStr := `insert into post(post_id,title,content,author_id,community_id) values (?,?,?,?,?)`
	_, err := db.Exec(sqlStr, post.ID, post.Title, post.Content, post.AuthorID, post.CommunityID)
	return err
}

// GetPostByID 根据id查询单篇帖子数据
func GetPostByID(pid int64) (*models.Post, error) {
	post := new(models.Post)
	sqlStr := `select post_id,title,content,author_id,community_id,create_time from post where id=?`
	err := db.Get(post, sqlStr, pid)
	return post, err
}

// GetPostList 查询帖子列表
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,community_id,create_time from post limit ?,?`
	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}

package models

import "time"

// 模型在设计的时候，记得要内存对齐

// Post 帖子结构体
type Post struct {
	ID          int64     `json:"id,string" db:"post_id"`
	AuthorID    int64     `json:"author_id,string" db:"author_id"`
	CommunityID int64     `json:"community_id,string" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

// ApiPostDetail 帖子详情
type ApiPostDetail struct {
	AuthorName       string                    `json:"author_name"`
	*Post                                      // 帖子结构体
	*CommunityDetail `json:"community_detail"` // 社区结构体
}

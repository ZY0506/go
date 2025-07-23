package models

// 存放请求的结构体

// 注册参数结构体
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// 登录参数结构体
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 投票数据结构体
type ParamVoteData struct {
	//UserID 从需求这种获取当前用户
	PostID    string `json:"post_id" binding:"required"`              // 帖子id
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` // 赞成票(1)还是反对票(-1)取消投票(0)
}

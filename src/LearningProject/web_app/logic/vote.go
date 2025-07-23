package logic

import (
	"LearningProject/web_app/dao/redis"
	"LearningProject/web_app/models"
	"go.uber.org/zap"
	"strconv"
)

// 本项目使用简化版的投票策略
// 投1 票就加432分 86400 / 200 => 200张赞成票就续一天

/* PostVote 为帖子投票
投票分为四种情况：1.投赞成票 2.投反对票 3.取消投票 4.反转投票

记录文章参与投票的人
更新文章分数：赞成票要加分；反对票减分

direction=1时，有两种情况
	1.之前没投过票，现在要投赞成票
	2.之前投过反对票，现在要改为赞成票
direction=0时，有两种情况
	1.之前投过赞成票，现在要取消
	2.之前投过反对票，现在要取消
direction=-1时，有两种情况
	1.之前没投过票，现在要投反对票
	2.之前投过赞成票，现在要改为反对票

投票限制：
每个帖子自发表之日起，一个星期之内允许用户投票，超过一个星期就不允许再投票了
	1、到期之后，将redis中保存的赞成票数和反对票数保存到mysql中
	2、到期之后，删除redis中保存的赞成票数和反对票数 -> KeyPostVotedZSetPrefix
*/

// PostVote 帖子投票
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	zap.L().Debug(
		"VoteForPost", zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("value", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))

}

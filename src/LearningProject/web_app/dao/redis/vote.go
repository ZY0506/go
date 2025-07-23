package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"time"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
)

// CreatePost 在redis中创建帖子
func CreatePost(id int64) error {
	pipeline := rdb.TxPipeline() // 事务
	// 帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: id,
	})

	// 帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  0,
		Member: id,
	})
	_, err := pipeline.Exec()
	return err
}

// VoteForPost 投票功能
func VoteForPost(userID, postID string, value float64) error {
	// 1.判断投票限制
	// 去redis取帖子发布时间
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	// 2、更新分数
	// 先查当前用户给当前帖子的投票记录
	ov := rdb.ZScore(getRedisKey(KeyPostVotedZSetPrefix+postID), userID).Val()
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value)
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID)
	// 3、记录用户为该帖子投票
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPrefix+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPrefix+postID), redis.Z{
			Score:  value,
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}

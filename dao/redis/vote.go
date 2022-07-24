package redis

import (
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis"
)

/*
投票功能
1.用户投票的数据
2.使用简单的投票分数算法
	direction =1 有两种情况：
	①之前没有投过票，现在投赞成票 --》 更新分数和投票记录 差值的绝对值：1
	②之前投反对票，现在投赞成票	 --》 更新分数和投票记录 差值的绝对值：2

	direction =0 有两种情况：
	①之前投赞成票，现在要取消投票
	②之前投反对票，现在要取消投票

	direction =-1 有两种情况：
	①之前没有投过票，现在投反对票
	②之前投赞成票，现在投反对票

投票的限制：
每个帖子自发表一个星期之内允许用户投票，超过一个星期就不允许用户再投票
	1.到期之后将redis 中保存的赞成票数及反对票数存储到mysql表中
	2.到期之后删除的那个 KeyPostVotedZSetPF
*/
const (
	oneWeekInSecond = 7 * 24 * 3600
	scorePerVote    = 432
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已经过了")
)

func CreatePost(postID int64) error {
	//保证在同一个事物中进行
	pipeline := rdb.TxPipeline()
	//帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	//帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	_, err := pipeline.Exec()
	return err
}
func VoteForPost(userID, postID string, value float64) error {
	//1.判断投票限制

	//去redis取帖子发布的时间
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSecond {
		return ErrVoteTimeExpire
	}
	//2.更新帖子分数
	//先查之前的投票记录
	ovalue := rdb.ZScore(getRedisKey(KeyPostVotedPrefix+postID), userID).Val()
	var dir float64
	if value > ovalue {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(ovalue - value) //计算两次投票的差值
	//放在同一个事物中
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), dir*diff*scorePerVote, postID)
	//if ErrVoteTimeExpire != nil {
	//	return err
	//}
	//3.记录用户为该帖子投票的数据
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedPrefix+postID), postID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedPrefix+postID), redis.Z{
			Score:  value, //赞成票还是反对票
			Member: nil,
		})
	}
	_, err := pipeline.Exec()
	return err
}

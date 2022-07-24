package redis

//redis key

//redis key 注意使用命名空间的方式,方便查询和拆分

const (
	KeyPrefix          = "weber"
	KeyPostTimeZSet    = "post:time"   //zset;帖子发帖时间
	KeyPostScoreZSet   = "post:score"  //zset;帖子及投票的分数
	KeyPostVotedPrefix = "post:voted:" //zset;参数是记录用户及投票的类型;参数是post_id
)

//给redis key 加上前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}

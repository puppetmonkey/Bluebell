package redis

// redis key

// redis key注意使用命名空间的方式,方便查询和拆分

const (
	Prefix             = "bluebell:"   // 项目key前缀
	KeyPostTimeZSet    = "post:time"   // zset;贴子及发帖时间 发布后面固定 key postid  value time.Now().unix()
	KeyPostScoreZSet   = "post:score"  // zset;贴子及投票的分数 key postid  value time.Now().unix()+-票数*432。   昨天需要获得200票才能和今天一样分数。
	KeyPostVotedZSetPF = "post:voted:" // zset;记录用户及投票类型;标题是post id 参数 key useid ；value direction(-1,0,1)

	KeyCommunitySetPF = "community:" // set;保存每个分区下帖子的id 01社区有1，2，3 ;03社区有2，3
)

// 给redis key加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}

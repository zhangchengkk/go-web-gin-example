package redis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go-web-gin-example/config"
	"strconv"
	"time"
)

const period = 5
const accessCount int64 = 50
/**
 * 限流方法（滑动时间算法）
 * @param key      限流标识
 * @param period   限流时间范围（单位：秒）
 * @param maxCount 最大运行访问次数
 * @return
 */
func IsPeriodLimiting(c *gin.Context, key string) bool {
	now := time.Now()
	duration, _ := time.ParseDuration("-"+ strconv.Itoa(period) +"s")
	before := now.Add(duration).UnixNano()
	config.RedisClient.ZRemRangeByScore(config.RedisContext, key, "0" , strconv.FormatInt(before, 10))
	//时间段的所有
	count := config.RedisClient.ZCard(config.RedisContext, key)

	if count.Val() >= accessCount { //并发数量控制
		c.JSON(403, map[string]int{"Res": 10000})
		c.Abort()
		return false
	}
	config.RedisClient.ZAdd(config.RedisContext, key, &redis.Z{
		Score:   float64(now.UnixNano()) ,
		Member: "request_time_" + strconv.FormatInt(now.UnixNano(), 10),
	})
	return true
}

package cache

import (
	"fmt"

	"gin_server/conf/setting"

	"git.shiyou.kingsoft.com/server/util/redis"
	r "github.com/gomodule/redigo/redis"
)

func GetRedis(redisConf setting.Redis) (rd *r.Pool) {
	c := redis.NewRedisPool(redisConf.REDIS_HOST, redisConf.REDIS_PORT, redisConf.REDIS_PASSWORD, 10, 10, 5)
	// redis.Dial("tcp", redisConf.REDIS_HOST+":"+redisConf.REDIS_PORT, redis.DialDatabase(redisConf.REDIS_DB), redis.DialPassword(redisConf.REDIS_PASSWORD))

	fmt.Println("connect redis success: ", redisConf.REDIS_HOST, ":", redisConf.REDIS_PORT)
	return c
}

package lib

import (
	"fmt"

	"gin_server/conf/setting"

	"git.shiyou.kingsoft.com/server/util/redis"
	r "github.com/gomodule/redigo/redis"
)

func GetRedis() (rd *r.Pool) {
	conf, _ := setting.Conf()

	c := redis.NewRedisPool(conf.Redis.REDIS_HOST, conf.Redis.REDIS_PORT, conf.Redis.REDIS_PASSWORD, 10, 10, 5)
	// redis.Dial("tcp", conf.Redis.REDIS_HOST+":"+conf.Redis.REDIS_PORT, redis.DialDatabase(conf.Redis.REDIS_DB), redis.DialPassword(conf.Redis.REDIS_PASSWORD))

	fmt.Println("connect redis success: ", conf.Redis.REDIS_HOST, ":", conf.Redis.REDIS_PORT)
	return c
}

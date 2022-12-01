package util

import (
	"fmt"

	"gin_server/conf/setting"

	"github.com/gomodule/redigo/redis"
)

func GetRedis() (rd redis.Conn) {
	conf := setting.Conf()

	c, err := redis.Dial("tcp", conf.Redis.REDIS_HOST+":"+conf.Redis.REDIS_PORT, redis.DialDatabase(conf.Redis.REDIS_DB), redis.DialPassword(conf.Redis.REDIS_PASSWORD))
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	fmt.Println("connect redis success: ", conf.Redis.REDIS_HOST+":"+conf.Redis.REDIS_PORT)
	return c
}

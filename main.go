package main

import (
	"fmt"

	"gin_server/app/captcha"
	"gin_server/app/user"
	"gin_server/conf/setting"
	"gin_server/lib"
	"gin_server/routers"
)

func main() {
	// 加载多个 APP 的路由配置
	routers.Include(user.Users, captcha.Captchas)
	// 加载环境变量配置
	setting.InitSetting()
	conf, _ := setting.Conf()
	// MySQL 连接
	lib.InitDB()
	// Redis 连接
	lib.GetRedis()
	// 初始化路由
	r := routers.Init()
	if err := r.Run(conf.Project.APP_PORT); err != nil {
		fmt.Println("startup service failed, err: %v\n", err)
	}
}

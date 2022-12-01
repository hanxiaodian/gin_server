package main

import (
	"fmt"

	"gin_server/app/captcha"
	"gin_server/app/user"
	"gin_server/conf/setting"
	"gin_server/routers"
	"gin_server/util"
)

func main() {
	// 加载多个APP的路由配置
	routers.Include(user.Users, captcha.Captchas)
	// 加载环境变量配置
	setting.InitSetting()
	conf := setting.Conf()
	// MySQL连接
	util.InitDB()
	// Redis连接
	util.GetRedis()
	// 初始化路由
	r := routers.Init()
	if err := r.Run(conf.Project.APP_PORT); err != nil {
		fmt.Println("startup service failed, err: %v\n", err)
	}
}

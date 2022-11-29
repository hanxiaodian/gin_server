package main

import (
	"fmt"

	"github.com/hanxiaodian/gin_server/app/captcha"
	"github.com/hanxiaodian/gin_server/app/user"
	"github.com/hanxiaodian/gin_server/conf/setting"
	"github.com/hanxiaodian/gin_server/routers"
)

func main() {
	// 加载多个APP的路由配置
	routers.Include(user.Users, captcha.Captchas)
	// 加载环境变量配置
	setting.InitSetting()
	conf := setting.Conf()
	// 初始化路由
	r := routers.Init()
	if err := r.Run(conf.Project.APP_PORT); err != nil {
		fmt.Println("startup service failed, err: %v\n", err)
	}
}

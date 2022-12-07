package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"gin_server/conf/setting"
	"gin_server/lib/cache"
	"gin_server/lib/db"
	"gin_server/middleware"
	"gin_server/routers"

	"git.shiyou.kingsoft.com/go/graceful"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 中间件加载
	r.Use(middleware.HealthCheck)

	// 部分项目功能初始化
	conf := serverInit(r)

	server := &http.Server{
		Addr:    conf.Project.APP_PORT,
		Handler: r,
	}

	go func() {
		// service connections
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("startup service failed, err: %s\n", err)
		}
	}()

	log.Printf("golang zt server, project %s is running in %s\n", conf.Project.PROJECT_NAME, conf.Project.APP_PORT)

	gracefulExit(server)
}

func serverInit(r *gin.Engine) (conf *setting.Setting) {
	// 加载环境变量配置
	setting.InitSetting()
	config, _ := setting.Conf()
	// 初始化 MySQL
	db.InitDB(config.DataBase)
	// 初始化 Redis
	cache.GetRedis(config.Redis)
	// 初始化路由加载
	group := r.Group(config.Project.PROJECT_NAME)
	routers.InitRouter(group, config.Project.PROJECT_NAME)

	return config
}

func gracefulExit(server *http.Server) {
	// 优雅退出
	wait := graceful.Shutdown(
		context.Background(),
		2*time.Second,
		[]graceful.Operation{
			func(ctx context.Context) {
				server.Shutdown(ctx)
			},
		},
	)

	<-wait

	log.Println("App exit, Good bye ...")
}

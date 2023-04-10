package main

import (
	"demo-project/app/internal/config"
	"demo-project/app/internal/handler"
	"demo-project/app/internal/middleware"
	"demo-project/app/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/app-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config

	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	//解决跨域 如果单个域名可以这么写：
	//srv := rest.MustNewServer(c, rest.WithCors("http://example.com"))
	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	redisConf := redis.RedisConf{Host: ctx.Config.Redis.Host, Pass: ctx.Config.Redis.Pass, Type: ctx.Config.Redis.Type}
	//全局中间件 限流
	server.Use(middleware.NewTokenLimiterMiddleware(redisConf).Handle)

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

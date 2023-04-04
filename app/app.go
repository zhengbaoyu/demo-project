package main

import (
	"demo-project/app/internal/config"
	"demo-project/app/internal/handler"
	"demo-project/app/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
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

	//全局中间件 解决跨域
	//server.Use(func(next http.HandlerFunc) http.HandlerFunc {
	//	return middleware.CorsHandle(next)
	//})

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

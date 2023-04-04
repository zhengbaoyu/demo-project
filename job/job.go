package main

import (
	"flag"
	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"

	"demo-project/job/internal/config"
	"demo-project/job/internal/service"
	"demo-project/job/internal/svc"
)

var configFile = flag.String("f", "etc/job.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)

	crontab := cron.New(cron.WithSeconds()) //精确到秒

	//全局s日志配置
	logConf := logx.LogConf{
		Mode:     ctx.Config.Log.Mode,
		Encoding: ctx.Config.Log.Encoding,
		Path:     ctx.Config.Log.Path,
		Level:    ctx.Config.Log.Level,
	}

	service.CronInit(ctx, crontab, logConf)
	// 启动定时器
	crontab.Start()
	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer crontab.Stop()
	// 定时任务是另起协程执行的,这里使用 select 简答阻塞.实际开发中需要
	// 根据实际情况进行控制
	select {} //阻塞主线程停止
}

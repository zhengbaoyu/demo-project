package service

import (
	"demo-project/job/internal/service/user"
	"demo-project/job/internal/svc"
	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/logx"
)

func CronInit(ctx *svc.ServiceContext, cron *cron.Cron, logConf logx.LogConf) {
	logx.MustSetup(logConf)
	var err error
	userNew := user.NewUser(ctx, logConf)

	_, err = cron.AddFunc(ctx.Config.CronConf.User, userNew.UserLogic)
	if err != nil {
		logx.Error("user crontab err", err)
	}
}

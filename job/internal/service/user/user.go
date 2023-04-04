package user

import (
	"demo-project/job/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type user struct {
	SvcCtx *svc.ServiceContext
}

func NewUser(svcCtx *svc.ServiceContext, logConf logx.LogConf) *user {
	logx.MustSetup(logConf)
	return &user{
		SvcCtx: svcCtx,
	}
}

func (t *user) UserLogic() {
	logx.Info("user start.....")

}

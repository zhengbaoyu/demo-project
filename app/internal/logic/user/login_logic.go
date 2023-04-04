package user

import (
	"context"
	"demo-project/common/jwt"
	"time"

	"demo-project/app/internal/svc"
	"demo-project/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginUserReq) (resp *types.LoginUserResp, err error) {
	userInfo, err := l.svcCtx.DemoUserModel.GetInfoByUsername(l.ctx, l.svcCtx.Orm, req.UserName)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	resp = &types.LoginUserResp{}
	payloads := make(map[string]any)
	payloads["uid"] = userInfo.Id
	payloads["username"] = userInfo.UserName
	accessToken, tokenErr := jwt.GetToken(time.Now().Unix(), l.svcCtx.Config.JwtAuth.AccessSecret, payloads, l.svcCtx.Config.JwtAuth.AccessExpire)
	if tokenErr != nil {
		l.Logger.Error(tokenErr)
		return resp, tokenErr
	}
	resp.Expire = time.Now().Unix() + l.svcCtx.Config.JwtAuth.AccessExpire
	resp.Token = accessToken
	return
}

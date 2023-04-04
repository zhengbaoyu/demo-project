package user

import (
	"context"
	"demo-project/app/internal/svc"
	"demo-project/app/internal/types"
	"demo-project/common/xerr"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info(req *types.UserInfoReq) (resp *types.UserResp, err error) {
	//token加密内的值
	fmt.Println("uid: %v", l.ctx.Value("uid"))
	fmt.Println("username: %v", l.ctx.Value("username"))
	//获取用户
	userInfo, err := l.svcCtx.DemoUserModel.GetInfoById(l.ctx, l.svcCtx.Orm, req.Uid)
	if err != nil {
		if err != sqlx.ErrNotFound {
			return nil, xerr.NewErrCode(xerr.DataNoExistError)
		}
		return nil, err
	}
	resp = &types.UserResp{
		UserName:  userInfo.UserName,
		Email:     userInfo.Email,
		Avatar:    userInfo.Avatar,
		CreatedAt: userInfo.CreatedAt.Unix(),
		UpdatedAt: userInfo.UpdatedAt.Unix(),
	}
	return
}

package user

import (
	"demo-project/common/result"
	"net/http"

	"demo-project/app/internal/logic/user"
	"demo-project/app/internal/svc"
	"demo-project/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func InfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewInfoLogic(r.Context(), svcCtx)
		resp, err := l.Info(&req)
		result.HttpResult(r, w, resp, err)
	}
}

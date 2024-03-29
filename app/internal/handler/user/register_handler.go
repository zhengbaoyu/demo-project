package user

import (
	"net/http"

	"demo-project/app/internal/logic/user"
	"demo-project/app/internal/svc"
	"demo-project/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewRegisterLogic(r.Context(), svcCtx)
		err := l.Register(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

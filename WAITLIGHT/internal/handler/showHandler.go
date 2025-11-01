// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"ztt1/internal/logic"
	"ztt1/internal/svc"
	"ztt1/internal/types"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShowRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		//参数规则校验
		//客户端可以看到什么由http.err决定
		if err:=validator.New().StructCtx(r.Context(),&req); err!=nil{
			logx.Error("validator check failed",logx.LogField{Key:"err",Value:err.Error()})
			httpx.ErrorCtx(r.Context(),w,err)
			return 
		}

		l := logic.NewShowLogic(r.Context(), svcCtx)
		resp, err := l.Show(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			// httpx.OkJsonCtx(r.Context(), w, resp)
			//返回重定向响应
			http.Redirect(w,r,resp.LongUrl,http.StatusFound)
		}
	}
}

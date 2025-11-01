// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"ztt1/internal/svc"
	"ztt1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortLogic {
	return &ShortLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortLogic) Short(req *types.ShowRequest) (resp *types.ShowResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"ztt1/internal/svc"
	"ztt1/internal/types"

	// "github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/logx"
	// "github.com/zeromicro/go-zero/core/stores/redis"
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.ShowRequest) (resp *types.ShowResponse, err error) {
	// todo: add your logic here and delete this line
	//查看短连接 /abcd->重定向到真实链接
	//req.ShortUrl
	//1.根据短连接查询长连接
	//优化 写缓存
	//自己写缓存 surl->lurl

	//查缓存之前使用布隆过滤器
	//a.基于内存版本
	//b.基于redis版本
	exit,err:=l.svcCtx.Filter.Exists([]byte(req.ShortUrl))
	if err!=nil{
		logx.Errorw("Filter failed",logx.LogField{Value: err.Error(),Key: "err"})
	}
	//不存在
	if !exit{
		return nil,errors.New("404 bloom failed")
	}
	fmt.Println("数据存在 开始查询缓存和DB...")
	//go-zero surl->数据行
	//go-zero缓存本身支持signfilght
	u,err:=l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx,sql.NullString{Valid: true,String: req.ShortUrl})
	if err!=nil{
		if err ==sql.ErrNoRows{
			return nil,errors.New("404")
		}
		logx.Errorw("ShortUrlModel.FindOneBySurl failed",logx.LogField{Value: err.Error(),Key: "err"})
		return nil,err
	}
	
	//2.返回重定向响应 返回查询到的长连接，在调用handler层返回重定向响应
	return &types.ShowResponse{LongUrl: u.Lurl.String},nil
}

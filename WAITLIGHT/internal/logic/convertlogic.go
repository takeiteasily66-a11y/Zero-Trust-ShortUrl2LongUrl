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
	"ztt1/model"
	"ztt1/pkg/base62"
	"ztt1/pkg/connect"
	"ztt1/pkg/md5"
	"ztt1/pkg/urltool"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

)

type ConvertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//输入长连接转为短连接
func (l *ConvertLogic) Convert(req *types.ConvertRequest) (resp *types.ConvertResponse, err error) {
	// todo: add your logic here and delete this line
	//1.首先接收长连接
	//检验
	//1.1数据不能为空
	// if req.LongUrl==0{
	// 	return
	// }
	//handler vilate设置了
	//1.2输入数据为ping通
	// connect.Get()
	if ok :=connect.Get(req.LongUrl);!ok{
		return nil,errors.New("无效的链接")
	}
	//1.3判断数据库中是否存在
	//1.3.1
	md5Value:=md5.Sum([]byte(req.LongUrl))//
	u, errr:=l.svcCtx.ShortUrlModel.FindOneByMd5(l.ctx,sql.NullString{String: md5Value,Valid: true})
	if errr!=sqlx.ErrNotFound{
		if errr==nil{
			return nil,fmt.Errorf("链接已经存在数据库当中%s",u.Surl.String)
		}
		logx.Errorw("ShortUrlModel.FindOneByMd5",logx.LogField{Key: "err",Value: errr.Error()})
		return nil,errr
	}

	//1.4避免循环转链 输入的需要时短连接
	//输入的是一个完整url 把后一截拿去判断
	basePath,err:=urltool.GetBasePath(req.LongUrl)
	if err!=nil {
		logx.Errorw("url,parse failed",logx.LogField{Key: "lurl",Value: req.LongUrl})
		return nil,err
	}
	_,err=l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx,sql.NullString{String: basePath,Valid: true})
	if err!=sqlx.ErrNotFound{
		if err==nil{
			return nil,errors.New("该链接已经是短链")
		}
		logx.Errorw("ShortUrlModel.FindOneByurl",logx.LogField{Key: "err",Value: errr.Error()})
		return nil,err
	}
	var short string
	//黑名单所以循环
	for{
		//2.转换-取号之类的
		//每来一个转链请求使用replace into 在mysql sequence整数据并取出主键id
		seq,err:=l.svcCtx.Sequence.Next()
		if err!=nil{
			logx.Errorw("Sequence.Next() failed",logx.LogField{Key: "err",Value: err.Error()})
			return nil,err
		}
		
		fmt.Println(seq)
		//3.存入数据库
		//3.1号码转换
		
		short=base62.Int2String(seq)
		fmt.Println(short)
		//安全性：1En 黑客可能一直遍历请求各种页面 可以把62base62str进制完全打乱
			//直接在yam文件中打乱顺序
		//黑名单机制 避免不文明或者特殊词汇出现 //这里下去可以list转为map
		if _,ok:=l.svcCtx.ShortUrlBlackList[short];!ok{
			break//生成ok就跳出来
		}
	}

	//4.存长短映射关系
	if _,err:=l.svcCtx.ShortUrlModel.Insert(
		l.ctx,
		&model.ShortUrlMap{
			Lurl: sql.NullString{String: req.LongUrl,Valid: true},
			Md5: sql.NullString{String: md5Value,Valid: true},
			Surl: sql.NullString{String: short,Valid: true},
		},
	);err!=nil{
		logx.Errorw("ShortUrlModel.Insert failed",logx.LogField{Key: "err",Value: err.Error()})
			return nil,err
	}
	//将生成的短连接加到布隆过滤器中
	if err:=l.svcCtx.Filter.Add([]byte(short));err!=nil{
		logx.Errorw("Bloom.Add failed",logx.LogField{Key: "err",Value: err.Error()})
	}
	//5.返回响应
	//5.1返回短域名+短连接..
	shortUrl:=l.svcCtx.Config.ShortDomain+"/"+short
	return &types.ConvertResponse{
		ShortUrl: shortUrl},nil
}


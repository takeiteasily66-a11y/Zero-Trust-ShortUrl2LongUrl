// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"ztt1/internal/config"
	"ztt1/model"
	"ztt1/sequence"

	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	//加数据库
	ShortUrlModel model.ShortUrlMapModel
	// Sequence  *sequence.MySQL
	// Sequence  *sequence.Redis
	Sequence sequence.Sequence
	//map黑名单
	ShortUrlBlackList map[string]struct{}
	//布隆过滤器
	//a.内存版本 重启后就没了 每次重启要加载已有连接
	Filter *bloom.Filter
	//b.基于redis版本，go-zero自带布隆过滤器



}

func NewServiceContext(c config.Config) *ServiceContext {
	//初始化连接
	conn:=sqlx.NewMysql(c.ShortUrlDB.DSN)
	m:=make(map[string]struct{},len(c.ShortUrlBlackList))
	//加载map
	for _,v :=range c.ShortUrlBlackList{
		m[v]=struct{}{}
	}
	//初始化布隆过滤器
	store:=redis.New(c.CacheRedis[0].Host,func(r *redis.Redis) {
		r.Type=redis.NodeType
	})
	filter:=bloom.New(store,"bloom_filter",20*(1<<20))

	//加载已有的短连接数据
	return &ServiceContext{
		Config: c,
		ShortUrlModel: model.NewShortUrlMapModel(conn,c.CacheRedis),
		// Sequence: sequence.NewMySQL(c.Sequence.DSN),
		Sequence: sequence.NewRedis(c.Sequence.DSN),
		ShortUrlBlackList: m,
		Filter: filter,
	}
	//上面也是初始化连接
	// return &ServiceContext{
	// 	Config: c,
	// }
}
//加载已有的数据至布隆过滤器之前的数据加载
func loadDataToBloomFilter(){
	
}
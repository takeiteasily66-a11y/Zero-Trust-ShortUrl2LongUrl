package sequence

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Redis struct{
	//redis链接
	rdb *redis.Redis
}

func NewRedis(redisAddr string) Sequence{
	return &Redis{
		rdb: redis.New(redisAddr),
	}
}

func (r *Redis) Next() (seq uint64,err error){
	//redis实现发号器
	//incr
	val,err:=r.rdb.Incr("sequence:incr")
	if err!=nil{
		logx.Errorw("redis incr failed", logx.LogField{Key: "err", Value: err.Error()})
		return 0, err
	}
	return uint64(val),nil
}
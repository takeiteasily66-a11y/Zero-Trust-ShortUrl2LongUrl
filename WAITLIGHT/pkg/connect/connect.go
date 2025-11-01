package connect

import (
	"net/http"
	// "net/url"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

//client 全局客户端
var client = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
	Timeout: 2*time.Second,
}

//判断URL是否ping通
func Get(url string) bool{
	resp,err:=client.Get(url)
	if err!=nil{
		logx.Errorw("connect client.Get failed",logx.LogField{Key: "err",Value: err.Error()})
		return false
	}
	resp.Body.Close()
	return resp.StatusCode==http.StatusOK//跳转连接也不算301码
}
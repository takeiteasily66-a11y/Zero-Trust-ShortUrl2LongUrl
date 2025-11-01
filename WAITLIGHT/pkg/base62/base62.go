package base62

import (
	"math"
	"strings"
)

//62进制转换模块
//0123456789[a-zA-Z]
//0-9:0-9
//a-z:10-35
//A-Z:36-61
//例子
//0-0 10-a 11-b 61-z 62-10 63-11
var(
	base62str string
)
//使用base62包必须调用该方法完成初始化
func MustInit(bs string){
	if len(bs)==0{
		panic("need base string")
	}
	base62str=bs


}
//十进制转62进制
func Int2String(seq uint64) string{
	if seq ==0{
		return  string(base62str[0])
	}
	bl:=[]byte{}
	for seq>0{
		mod:=seq%62
		div:=seq/62
		bl=append(bl,base62str[mod])
		seq=div
	}
	//还需要翻转
	return  string(reverse(bl))
}

//62进制转十进制
func String2Int(s string) (seq uint64){
	bl:=[]byte(s)
	bl=reverse(bl)
	//从右往左遍历
	for idx,b :=range bl{
		base:=math.Pow(62,float64(idx))
		seq+=uint64(strings.Index(base62str,string(b)))*uint64(base)
	}
	return seq



}
func reverse(s []byte) []byte {
	for i,j:=0,len(s)-1; i < len(s)/2; i,j=i+1,j-1 {
		s[i],s[j]=s[j],s[i]
	}
	return s

}
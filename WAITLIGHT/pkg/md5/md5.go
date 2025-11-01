package md5

import (
	"crypto/md5"
	"encoding/hex"
)

//对传入参数求md5值
func Sum(data []byte) string{
	h:=md5.New()
	h.Write(data)
	return  hex.EncodeToString(h.Sum(nil))//32位16进制数

}
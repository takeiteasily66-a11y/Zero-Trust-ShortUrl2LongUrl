package urltool

import (
	"errors"
	"net/url"
	"path"
)

func GetBasePath(targeturl string)(string,error){
	myurl,err:=url.Parse(targeturl)
	if err!=nil {
		return "",err
	}
	if len(myurl.Host)==0{
		return "",errors.New("无host")
	}
	return  path.Base(myurl.Path),nil
}
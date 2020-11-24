package api

import (
	"time"
)

type ApiData struct {
	AppCode    string
	StreamCode string
	LiveUrl    string
	NotifyUrl  string
	DomainName string
	IsMobile   int
}

var loc, _ = time.LoadLocation("Local")
var fDate = "2006-01-02 15:04:05"
var fDay = "2006-01-02"

//数据缓存的key集合
var SessionRedisKey = "Session_Redis_Key"
var httpHeader map[string]string

func init() {
	httpHeader = make(map[string]string)
	httpHeader["content-type"] = "application/json; charset=gb2312"
	httpHeader["accept"] = "application/json, text/plain, */*"
	httpHeader["cache-control"] = "no-cache"
	httpHeader["sec-fetch-mode"] = "cors"
	httpHeader["sec-fetch-site"] = "cross-site"
	httpHeader["user-agent"] = "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1"
	httpHeader["Accept-Language"] = "zh-CN,zh;q=0.9"
}

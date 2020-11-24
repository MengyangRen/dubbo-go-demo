package middleware

import (
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

/**
* sql注入过滤判断
* @param	str	需要做判断的字符串
* return	bool	如果包含注入关键词输出true，否则false
 */
func IsAllow(str string) bool {
	hRes := true
	str = strings.ToLower(str)
	dangerKeys := `(?:')|(?:--)|(\b(select|update|delete|insert|trancate|char|into|substr|ascii|declare|exec|count|master|into|drop|execute|script)\b)`
	re, _ := regexp.Compile(dangerKeys)
	if re.MatchString(str) {
		hRes = false
	}
	return hRes
}

/**
*  正则判断
*  reg  正则表达式
*  str  需要判断的字符串
 */
func authReg(str, reg string) bool {
	res, _ := regexp.MatchString(reg, str)
	return res
}

func AuthInput(keyArr []string, ctx *gin.Context) (int, string, map[string]string) {
	hCode := 100
	hMsg := "参数不合法"
	valMap := map[string]string{}
	if len(keyArr) < 1 {
		hCode = 200
		hMsg = "success"
		return hCode, hMsg, valMap
	}

	for _, kVal := range keyArr {
		val := ctx.PostForm(kVal)
		if val == "" {
			continue
		}
		if !IsAllow(val) {
			return hCode, hMsg, valMap
		}
		valMap[kVal] = val
	}
	hCode = 200
	hMsg = "success"
	return hCode, hMsg, valMap
}

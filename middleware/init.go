package middleware

import (
	"encoding/json"

	"gitlab.stagingvip.net/publicGroup/public/common"
)

var logPath string

func init() {
	confByte, err := common.ReadFile("./conf/conf.json")

	if err != nil {
		panic(err)
	}
	var jsonConf map[string]string
	//解析json格式r
	err = json.Unmarshal(confByte, &jsonConf)
	if err != nil {
		panic(err)
	}
	logPath = jsonConf["logPath"]
}

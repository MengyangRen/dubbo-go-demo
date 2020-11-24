package main

import (
	"cloudlive/router"
	"encoding/json"

	"gitlab.stagingvip.net/publicGroup/public/common"
)

func main() {
	confByte, err := common.ReadFile("./conf/conf.json")
	if err != nil {
		panic(err)
	}
	var jsonConf map[string]string
	//解析json格式
	err = json.Unmarshal(confByte, &jsonConf)
	if err != nil {
		panic(err)
	}
	_ = router.Router.Run(jsonConf["port"])
}

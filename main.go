package main

import (
	"encoding/json"

	"gin-dubbogo-consumer/router"

	"gitlab.stagingvip.net/publicGroup/public/common"
)

func main() {
	confByte, err := common.ReadFile("./conf/conf.json")
	if err != nil {
		panic(err)
	}

	var jsonConf map[string]string
	err = json.Unmarshal(confByte, &jsonConf)
	if err != nil {
		panic(err)
	}
	_ = router.Router.Run(jsonConf["port"])
}

package dnc

import (
	"cloudlive/util"

	hessian "github.com/apache/dubbo-go-hessian2"
	"github.com/apache/dubbo-go/config"
)

var (
	survivalTimeout int = 10e9
	UserPvder           = new(UserProvider)
)

const (
	// dev,test,release
	ENV = "dev"
	//base path
	DNC_BASE_PATH = "./conf/dubbo/nacos"
)

func init() {
	config.SetConsumerService(UserPvder)
	
	hessian.RegisterPOJO(&User{})

	hessian.RegisterPOJO(&UserPaginationQ{})
	hessian.RegisterPOJO(&UserStoreResult{})
	_confConFile := new(util.DubboCtConfiguration).SetEnv(ENV).SetBasePath(DNC_BASE_PATH).GetRealPath()
	config.ConsumerInit(_confConFile)
	config.Load()
}

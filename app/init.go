package main

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 *  全局初始化
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package init
 * @description
 *
 * 说明:
 *  1.初始化 生产者向注册中心注册
 *  2.初始化 mysql
 *  3.初始化 redis
 *  4.初始化 Signal  (daemon)
 */

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	hessian "github.com/apache/dubbo-go-hessian2"
	_ "github.com/apache/dubbo-go/cluster/cluster_impl"
	_ "github.com/apache/dubbo-go/cluster/loadbalance"
	"github.com/apache/dubbo-go/common/logger"
	_ "github.com/apache/dubbo-go/common/proxy/proxy_factory"
	"github.com/apache/dubbo-go/config"
	_ "github.com/apache/dubbo-go/filter/filter_impl"
	_ "github.com/apache/dubbo-go/protocol/dubbo"
	_ "github.com/apache/dubbo-go/registry/nacos"
	_ "github.com/apache/dubbo-go/registry/protocol"
	_ "github.com/go-sql-driver/mysql"

	"dev-dubbo-producer/app/processor"
	"dev-dubbo-producer/app/util"
)

var (
	survivalTimeout = int(3e9)
)

const (
	// dev,test,release
	ENV = "dev"
	//base path
	DNC_BASE_PATH = "../profiles"
)

type Init struct{}

func NewInit() *Init {
	return &Init{}
}

func (this *Init) RegisterProvider(isDebug bool) {

	config.SetProviderService(new(processor.UserProvider))
	// ------for hessian2------
	hessian.RegisterPOJO(&processor.User{})
	hessian.RegisterPOJO(&processor.UserPaginationQ{})
	hessian.RegisterPOJO(&processor.UserStoreResult{})

	//服务端调试模式下 使用自定义配置
	if isDebug {
		config.ProviderInit(new(util.DubboCtConfiguration).SetEnv(ENV).SetBasePath(DNC_BASE_PATH).GetRealPath())
	}
	config.Load()
}

func (this *Init) Signal() {
	signals := make(chan os.Signal, 1)
	// It is not possible to block SIGKILL or syscall.SIGSTOP
	signal.Notify(signals, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-signals
		logger.Infof("get signal %s", sig.String())
		switch sig {
		case syscall.SIGHUP:
			// reload()
		default:
			time.AfterFunc(time.Duration(survivalTimeout), func() {
				logger.Warnf("app exit now by force...")
				os.Exit(1)
			})

			// The program exits normally or timeout forcibly exits.
			fmt.Println("provider app exit now...")
			return
		}
	}
}

package util

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 *  Dubbo 生产端自定义文件类
 * @author  m.y <j01428@kok.work>
 * @date
 * @package controllers
 * @description
 *
 * 说明:
 *  //获取配置文件路径
 * new(DubboCtConfiguration).SetEnv("dev").SetBasePath("../profiles").SetConfStream().GetRealPath()
 * //获取配置文件内容
 * new(DubboCtConfiguration).SetEnv("dev").SetBasePath("../profiles").SetConfStream().GetConf()
 * //获取配置文件流
 * new(DubboCtConfiguration).SetEnv("dev").SetBasePath("../profiles").SetConfStream().GetConfStream()
 */
import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type DubboCtConfiguration struct {
	execPath       string
	env            string
	basePath       string
	RealPath       string
	fileStream     []byte
	readYamlStream func(confPath string) ([]byte, error)
}

//support  dev,test,release
func (this *DubboCtConfiguration) SetEnv(env string) *DubboCtConfiguration {
	this.env = env
	return this
}

// ./profiles/dev/client.yaml
func (this *DubboCtConfiguration) SetBasePath(bp string) *DubboCtConfiguration {
	this.basePath = bp + "/" + this.env
	tmp := this.basePath + "/server.yml"
	this.RealPath, _ = filepath.Abs(tmp)
	return this
}

// real-path
func (this *DubboCtConfiguration) GetRealPath() string {
	return this.RealPath
}

// get conf file stream
func (this *DubboCtConfiguration) SetConfStream() *DubboCtConfiguration {
	f, err := os.Open(this.RealPath)
	if err != nil {
		log.Fatal("[Error] 检查本地消费者配置文件是否有效")
		//注意：该函数 目前熟悉度读不够 是退出子进程还是主进程 ，那么生产环境是有风险
		os.Exit(0)
	}
	this.fileStream, _ = ioutil.ReadAll(f)
	return this
}

func (this *DubboCtConfiguration) GetConfStream() []byte {
	return this.fileStream
}

func (this *DubboCtConfiguration) GetConf() string {
	return string(this.fileStream)
}

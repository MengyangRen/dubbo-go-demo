package model

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"gitlab.stagingvip.net/publicGroup/public/common"
)

var RedisCli *redis.Client
var G Gorm

type Gorm struct {
	DB *gorm.DB
}

type Init struct{}

func NewInit() *Init {
	return &Init{}
}

func (this *Init) MySql(fileName string) Gorm {
	if fileName == "" {
		fileName = "./conf/database.json"
	}
	confByte, err := common.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var jsonConf map[string]string
	//解析json格式
	err = json.Unmarshal(confByte, &jsonConf)
	if err != nil {
		panic(err)
	}
	lifeTime, _ := time.ParseDuration(jsonConf["life_time"])
	maxOpen, _ := strconv.Atoi(jsonConf["max_open"])
	if maxOpen < 1 {
		maxOpen = 40
	}
	maxIdle, _ := strconv.Atoi(jsonConf["max_idle"])
	if maxIdle < 1 {
		maxIdle = 10
	}
	connStr := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8mb4", jsonConf["user"], jsonConf["pwd"], jsonConf["network"],
		jsonConf["host"], jsonConf["port"], jsonConf["db_name"])
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		//_, file, line, _ := runtime.Caller(0)
		//logs.SaveLogs(file, line, err.Error(), 4)
		fmt.Println("connStr->", connStr)
		panic(err)
	}

	//最大生命周期
	db.DB().SetConnMaxLifetime(lifeTime)
	//连接池的最大打开连接数
	db.DB().SetMaxOpenConns(maxOpen)
	//连接池的最大空闲连接数
	db.DB().SetMaxIdleConns(maxIdle)
	db.SingularTable(true)
	//启用Logger，显示详细日志
	db.LogMode(true)
	// 禁用日志记录器，不显示任何日志
	//db.LogMode(false)
	return Gorm{DB: db}
}

//for hunter
func (this *Init) Redis() *redis.Client {
	confByte, err := common.ReadFile("./conf/redis.json")
	var jsonConf map[string]string
	//解析json格式
	err = json.Unmarshal(confByte, &jsonConf)
	if err != nil {
		//_, file, line, _ := runtime.Caller(0)
		//logs.SaveLogs(file, line, err.Error(), 4)
		panic(err)
	}

	addr := jsonConf["addr"]         //Redis的IP地址端口号
	password := jsonConf["password"] //Redis密码
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		//_, file, line, _ := runtime.Caller(0)
		//logs.SaveLogs(file, line, err.Error(), 4)
		return nil
	}

	client.Ping()
	//_, file, line, _ := runtime.Caller(0)
	//logs.SaveLogs(file, line, "Redis初始化状态:"+pong, 2)
	fmt.Println(pong)
	return client
}

func init() {
	G = NewInit().MySql("")
	RedisCli = NewInit().Redis()
}

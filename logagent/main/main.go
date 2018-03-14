package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"HelloGo/logagent/tail"
	"time"
	"os"
	"HelloGo/logagent/etcd"
)

func main() {
	fmt.Println(os.Getwd())
	err := loadConfig("ini", "./logagent/config/log.conf")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("load log.conf succcess", appConfig)
	initLogger()

	err=etcd.InitEtcd(appConfig.Endpoint)
	if err != nil {
		panic(err.Error())
	}

	go func() {
		var count int
		for {
			count++
			logs.Debug("COUNT:", count)
			time.Sleep(time.Second)
		}
	}()

	logs.Debug("init logger success")
	tail.InitTail(appConfig.CollectConf, appConfig.chanSize)


}

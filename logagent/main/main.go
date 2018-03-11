package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"HelloGo/logagent/tail"
	"time"
)

func main() {

	err := loadConfig("ini", "./logagent/config/log.conf")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("load log.conf succcess", appConfig)
	initLogger()

	go func() {
		var count int
		for {
			count++
			logs.Debug("COUNT:", count)
			time.Sleep(time.Millisecond * 100)
		}
	}()

	logs.Debug("init logger success")
	tail.InitTail(appConfig.CollectConf)

}

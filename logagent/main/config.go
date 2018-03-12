package main

import (
	"github.com/astaxie/beego/config"
	"fmt"
	"HelloGo/logagent/tail"
	"github.com/astaxie/beego/logs"
	"os"
)

type Config struct {
	LogLevel    string
	LogPath     string
	CollectConf []tail.CollectConf

	chanSize int
	Endpoint []string
}

var (
	appConfig *Config
)

func loadConfig(adapterName, filename string) (err error) {

	config, err := config.NewConfig(adapterName, filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	appConfig = &Config{}

	logLevel := config.String("logs::log_level")
	if len(logLevel) == 0 {
		logLevel = "debug"
	}

	appConfig.LogLevel = logLevel
	logPath := config.String("logs::log_path")

	if len(logPath) == 0 {
		fmt.Println("invald logpath")
		return
	}
	appConfig.LogPath = logPath

	cc := tail.CollectConf{}
	tialPath := config.String("collect::log_path")
	if len(tialPath) == 0 {
		logs.Error("tial file empty")
		panic("tail file empty")
		return
	}
	cc.LogPath = tialPath
	tialTopic := config.String("collect::topic")
	if len(tialTopic) == 0 {
		tialTopic = "nginx_log"
	}
	cc.Topic = tialTopic
	appConfig.CollectConf = append(appConfig.CollectConf, cc)

	chanSize, err := config.Int("collect::chanSize")
	if os.IsExist(err) {
		chanSize = 100
	}
	appConfig.chanSize = chanSize

	endPoint1 := config.String("etcd::Endpoint1")
	appConfig.Endpoint = append(appConfig.Endpoint, endPoint1)
	endPoint2 := config.String("etcd::Endpoint2")
	appConfig.Endpoint = append(appConfig.Endpoint, endPoint2)
	endPoint3 := config.String("etcd::Endpoint3")
	appConfig.Endpoint = append(appConfig.Endpoint, endPoint3)

	return
}

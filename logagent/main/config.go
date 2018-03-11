package main

import (
	"github.com/astaxie/beego/config"
	"fmt"
	"HelloGo/logagent/tail"
	"github.com/astaxie/beego/logs"
)

type Config struct {
	LogLevel    string
	LogPath     string
	CollectConf []tail.CollectConf
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
		panic("tial file empty")

		return
	}
	cc.LogPath = tialPath
	tialTopic := config.String("collect::topic")
	if len(tialTopic) == 0 {
		tialTopic = "nginx_log"
	}
	cc.Topic = tialTopic
	appConfig.CollectConf = append(appConfig.CollectConf, cc)




	return
}

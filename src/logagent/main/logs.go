package main

import (
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"fmt"
)

/**
   LevelEmergency = iota
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInformational
	LevelDebug
 */
func convertLogLevel(loglevel string) int {
	switch loglevel {
	case "debug":
		return logs.LevelDebug
	case "trace":
		return logs.LevelTrace
	case "info":
		return logs.LevelInfo
	}
	return logs.LevelDebug
}
func initLogger() {
	config := make(map[string]interface{})
	config["filename"] = appConfig.LogPath
	config["level"] = convertLogLevel(appConfig.LogLevel)
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}
	logs.SetLogger(logs.AdapterFile, string(configStr))
}

package main

import (
	"bandwidth_summary/conf"
	logs "bandwidth_summary/core/log"
	"bandwidth_summary/core/scheduled"
)

func main() {
	config, err := conf.LoadConf("./conf/config.yaml")
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}
	logs.InitLog(config.Base.Log.Dir, config.Base.Log.MaxSize, config.Base.Log.MaxBackups, config.Base.Log.MaxAge, config.Base.Log.Compress)
	scheduled.InitScheduled(config)
}

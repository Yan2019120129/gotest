package main

import (
	"embed"
	"operate/cmd"
	"operate/conf"
	"operate/core"
)

//go:embed conf/*
var fs embed.FS // Go 1.16 版本之后提供的将静态资源打包的方法，写法固定，可以将目录也打包

func main() {
	err := conf.InitConf(fs)
	if err != nil {
		panic(err)
	}

	core.InitLog(conf.Conf.Log.Dir, conf.Conf.Log.MaxSize, conf.Conf.Log.MaxBackups, conf.Conf.Log.MaxAge, conf.Conf.Log.Compress)
	cmd.Execute()
}

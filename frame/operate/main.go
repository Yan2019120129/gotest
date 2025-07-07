package main

import (
	"embed"
	"gopkg.in/yaml.v3"
	"operate/cmd"
	"operate/conf"
)

//go:embed conf/*
var fs embed.FS // Go 1.16 版本之后提供的将静态资源打包的方法，写法固定，可以将目录也打包

func main() {
	file, err := fs.ReadFile("conf/config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &conf.Conf)
	if err != nil {
		panic(err)
	}

	cmd.Execute()
}

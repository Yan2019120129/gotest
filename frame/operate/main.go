package main

import (
	"operate/cmd"
	"operate/conf"
)

func main() {
	if err := conf.InitConf("./conf/config.yaml"); err != nil {
		panic(err)
	}
	cmd.Execute()
}

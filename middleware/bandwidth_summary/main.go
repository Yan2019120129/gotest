package main

import (
	"bandwidth_summary/conf"
	"bandwidth_summary/core/scheduled"
	"context"
)

func main() {
	ctx := context.Background()
	config, err := conf.LoadConf()
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}
	ctx = context.WithValue(ctx, "checkInterval", config.Base.CheckInterval)
	ctx = context.WithValue(ctx, "addrs", config.Client.Address)
	scheduled.InitScheduled(ctx)
}

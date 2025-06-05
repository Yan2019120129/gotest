package scheduled

import (
	"context"
	"time"
)

// InitScheduled 初始化定时任务
func InitScheduled(ctx context.Context) {
	internal, ok := ctx.Value("checkInterval").(int)
	if !ok {
		panic("checkInterval not found in context")
	}

	addrs, ok := ctx.Value("addrs").([]string)
	if !ok {
		panic("checkInterval not found in context")
	}
	// 定时任务
	ticker := time.NewTicker(time.Duration(internal) * time.Second)
	for {
		// 获取客户端带宽数据
		GetClentBwSummary(addrs)
		<-ticker.C
	}
	ctx.Done()
}

package time_t

import (
	"fmt"
	"go.uber.org/zap"
	"gotest/common/module/logs"
	"time"
)

// TimeCreate 测试time.unix 创建时间是否会不一致
func TimeCreate() {
	nowTime := time.Now()
	fmt.Println("time1:", nowTime.Unix())
	fmt.Println("time2:", nowTime.Unix())
}

// TimeTicker 测试时间表
func TimeTicker(second time.Duration) {
	ch := time.NewTicker(second)

	for {
		logs.Logger.Info("time", zap.String("second", second.String()))
		logs.Logger.Info("time", zap.Reflect("ch", ch.C))
		<-ch.C
	}

}

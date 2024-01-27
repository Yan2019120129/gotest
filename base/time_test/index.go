package time_test

import (
	"fmt"
	"go.uber.org/zap"
	"gotest/common/module/logs"
	"time"
)

// TestTimeCreate 测试time.unix 创建时间是否会不一致
func TestTimeCreate() {
	nowTime := time.Now()
	fmt.Println("time1:", nowTime.Unix())
	fmt.Println("time2:", nowTime.Unix())
}

// TestTimeTicker 测试时间表
func TestTimeTicker(second time.Duration) {
	ch := time.NewTicker(second)

	for {
		logs.Logger.Info("time", zap.String("second", second.String()))
		logs.Logger.Info("time", zap.Reflect("ch", ch.C))
		<-ch.C
	}

}

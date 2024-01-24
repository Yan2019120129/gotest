package time_test

import (
	"fmt"
	"time"
)

// TestTimeCreate 测试time.unix 创建时间是否会不一致
func TestTimeCreate() {
	nowTime := time.Now()
	fmt.Println("time1:", nowTime.Unix())
	fmt.Println("time2:", nowTime.Unix())
}

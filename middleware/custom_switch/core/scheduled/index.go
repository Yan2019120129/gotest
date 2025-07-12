package scheduled

import (
	"bandwidth_summary/conf"
	"log"
	"time"
)

// InitScheduled 初始化定时任务
func InitScheduled(c *conf.Config) {
	if c.Base.CheckInterval == 0 {
		log.Fatal("checkInterval cannot be zero")
	}

	if len(c.Client.Address) == 0 {
		log.Fatal("address cannot be zero")
	}

	if c.Base.KgGroup == "" {
		log.Fatal("kg_group cannot be empty")
	}

	log.Println("custom_switch start success")
	// 定时任务
	ticker := time.NewTicker(time.Duration(c.Base.CheckInterval) * time.Second)
	for {
		// 获取客户端带宽数据
		reportTcInfo := GetClentBwSummary(c.Base.KgGroup, c.Client.Address)

		log.Println("reportTcInfo:", reportTcInfo)

		// 上报带宽
		err := ReportTcInfo(c.Base.TargetServer, reportTcInfo)
		if err != nil {
			log.Println("reportTcInfo error:", reportTcInfo)
		}
		<-ticker.C
	}
}

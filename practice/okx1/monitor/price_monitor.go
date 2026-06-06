package monitor

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var DefaultConfigs = map[string]MonitorConfig{
	"BTC-USDT": {
		OneMinute: AlertRule{
			Threshold: 0.8,
			Cooldown:  10 * time.Minute,
		},
		FiveMinute: AlertRule{
			Threshold: 1.5,
			Cooldown:  20 * time.Minute,
		},
		FifteenMin: AlertRule{
			Threshold: 3,
			Cooldown:  30 * time.Minute,
		},
		OneHour: AlertRule{
			Threshold: 4,
			Cooldown:  time.Hour,
		},
		SixHour: AlertRule{
			Threshold: 8,
			Cooldown:  2 * time.Hour,
		},
	},

	"ETH-USDT": {
		OneMinute: AlertRule{
			Threshold: 1,
			Cooldown:  10 * time.Minute,
		},
		FiveMinute: AlertRule{
			Threshold: 2,
			Cooldown:  20 * time.Minute,
		},
		FifteenMin: AlertRule{
			Threshold: 4,
			Cooldown:  30 * time.Minute,
		},
		OneHour: AlertRule{
			Threshold: 6,
			Cooldown:  time.Hour,
		},
		SixHour: AlertRule{
			Threshold: 10,
			Cooldown:  2 * time.Hour,
		},
	},

	"DOGE-USDT": {
		OneMinute: AlertRule{
			Threshold: 2,
			Cooldown:  5 * time.Minute,
		},
		FiveMinute: AlertRule{
			Threshold: 4,
			Cooldown:  15 * time.Minute,
		},
		FifteenMin: AlertRule{
			Threshold: 8,
			Cooldown:  30 * time.Minute,
		},
		OneHour: AlertRule{
			Threshold: 12,
			Cooldown:  time.Hour,
		},
		SixHour: AlertRule{
			Threshold: 20,
			Cooldown:  2 * time.Hour,
		},
	},

	"SOL-USDT": {
		OneMinute: AlertRule{
			Threshold: 1.5,
			Cooldown:  5 * time.Minute,
		},
		FiveMinute: AlertRule{
			Threshold: 3,
			Cooldown:  15 * time.Minute,
		},
		FifteenMin: AlertRule{
			Threshold: 6,
			Cooldown:  30 * time.Minute,
		},
		OneHour: AlertRule{
			Threshold: 10,
			Cooldown:  time.Hour,
		},
		SixHour: AlertRule{
			Threshold: 15,
			Cooldown:  2 * time.Hour,
		},
	},
	"ALLO-USDT": {
		OneMinute: AlertRule{
			Threshold: 2.5,
			Cooldown:  3 * time.Minute,
		},

		FiveMinute: AlertRule{
			Threshold: 5,
			Cooldown:  10 * time.Minute,
		},

		FifteenMin: AlertRule{
			Threshold: 10,
			Cooldown:  20 * time.Minute,
		},

		OneHour: AlertRule{
			Threshold: 18,
			Cooldown:  40 * time.Minute,
		},

		SixHour: AlertRule{
			Threshold: 30,
			Cooldown:  2 * time.Hour,
		},
	},
}

type AlertRule struct {
	Threshold float64       // 到达百分比预警
	Cooldown  time.Duration // 邮件冷却时间
}

type MonitorConfig struct {
	OneMinute  AlertRule
	FiveMinute AlertRule
	FifteenMin AlertRule
	OneHour    AlertRule
	SixHour    AlertRule
}

type PricePoint struct {
	Price float64
	Time  time.Time
}

type EmailSender interface {
	Send(
		to []string,
		subject string,
		body string,
	) error
}

type PriceMonitor struct {
	mu sync.RWMutex

	symbol string

	config MonitorConfig

	points []PricePoint

	maxPoints int

	last1mAlarm  time.Time
	last5mAlarm  time.Time
	last15mAlarm time.Time
	last1hAlarm  time.Time
	last6hAlarm  time.Time

	email EmailSender

	receivers []string
}

func NewPriceMonitor(
	symbol string,
	cfg MonitorConfig,
	email EmailSender,
	receivers []string,
) *PriceMonitor {
	return &PriceMonitor{
		symbol:    symbol,
		config:    cfg,
		maxPoints: 360,
		email:     email,
		receivers: receivers,
	}
}

func (m *PriceMonitor) Add(price float64) {

	m.mu.Lock()
	defer m.mu.Unlock()

	m.points = append(m.points, PricePoint{
		Price: price,
		Time:  time.Now(),
	})

	if len(m.points) > m.maxPoints {
		m.points = m.points[1:]
	}

	m.check()
}

func (m *PriceMonitor) check() {
	current := m.points[len(m.points)-1].Price

	m.check1m(current)
	m.check5m(current)
	m.check15m(current)
	m.check1h(current)
	m.check6h(current)
}

func (m *PriceMonitor) check1m(current float64) {

	if len(m.points) < 2 {
		return
	}

	old := m.points[len(m.points)-2].Price

	change := (current - old) / old * 100

	rule := m.config.OneMinute

	if change >= rule.Threshold ||
		change <= -rule.Threshold {

		if time.Since(m.last1mAlarm) < 10*time.Minute {
			return
		}

		m.last1mAlarm = time.Now()
		log.Println("1分钟异动: ", m.symbol, current, old, change)
		m.sendAlert(
			"1分钟异动",
			current,
			old,
			change,
		)
	}
}

func (m *PriceMonitor) check5m(current float64) {

	if len(m.points) < 5 {
		return
	}

	old := m.points[len(m.points)-5].Price

	change := (current - old) / old * 100

	rule := m.config.FiveMinute

	if change >= rule.Threshold ||
		change <= -rule.Threshold {

		if time.Since(m.last5mAlarm) < 30*time.Minute {
			return
		}

		m.last5mAlarm = time.Now()
		log.Println("5分钟异动: ", m.symbol, current, old, change)
		m.sendAlert(
			"5分钟异动",
			current,
			old,
			change,
		)
	}
}

func (m *PriceMonitor) check15m(current float64) {

	if len(m.points) < 15 {
		return
	}

	old := m.points[len(m.points)-15].Price

	change := (current - old) / old * 100

	rule := m.config.FifteenMin

	if change >= rule.Threshold ||
		change <= -rule.Threshold {

		if time.Since(m.last15mAlarm) < time.Hour {
			return
		}

		m.last15mAlarm = time.Now()

		log.Println("15分钟异动: ", m.symbol, current, old, change)
		m.sendAlert(
			"15分钟异动",
			current,
			old,
			change,
		)
	}
}

func (m *PriceMonitor) check1h(current float64) {

	if len(m.points) < 60 {
		return
	}

	avg := average(
		m.points[len(m.points)-60:],
	)

	change := (current - avg) / avg * 100

	rule := m.config.OneHour

	if change >= rule.Threshold ||
		change <= -rule.Threshold {

		if time.Since(m.last1hAlarm) < time.Hour {
			return
		}

		m.last1hAlarm = time.Now()
		log.Println("1小时均价偏离: ", m.symbol, current, avg, change)
		m.sendAvgAlert(
			"1小时均价偏离",
			current,
			avg,
			change,
		)
	}
}

func (m *PriceMonitor) check6h(current float64) {

	if len(m.points) < 360 {
		return
	}

	avg := average(m.points)

	change := (current - avg) / avg * 100

	rule := m.config.SixHour
	if change >= rule.Threshold ||
		change <= -rule.Threshold {

		if time.Since(m.last6hAlarm) < 2*time.Hour {
			return
		}

		m.last6hAlarm = time.Now()

		log.Println("6小时均价偏离: ", m.symbol, current, avg, change)
		m.sendAvgAlert(
			"6小时均价偏离",
			current,
			avg,
			change,
		)
	}
}

func average(points []PricePoint) float64 {

	var total float64

	for _, p := range points {
		total += p.Price
	}

	return total / float64(len(points))
}

func (m *PriceMonitor) sendAlert(
	alertType string,
	current float64,
	base float64,
	change float64,
) {

	subject := fmt.Sprintf(
		"[%s] %s",
		m.symbol,
		alertType,
	)

	body := fmt.Sprintf(
		`币种: %s

当前价格: %.4f

基准价格: %.4f

涨跌幅: %.2f%%

时间: %s`,
		m.symbol,
		current,
		base,
		change,
		time.Now().Format("2006-01-02 15:04:05"),
	)

	_ = m.email.Send(
		m.receivers,
		subject,
		body,
	)
}

func (m *PriceMonitor) sendAvgAlert(
	alertType string,
	current float64,
	avg float64,
	change float64,
) {

	subject := fmt.Sprintf(
		"[%s] %s",
		m.symbol,
		alertType,
	)

	body := fmt.Sprintf(
		`币种: %s

当前价格: %.4f

均价: %.4f

偏离幅度: %.2f%%

时间: %s`,
		m.symbol,
		current,
		avg,
		change,
		time.Now().Format("2006-01-02 15:04:05"),
	)

	_ = m.email.Send(
		m.receivers,
		subject,
		body,
	)
}

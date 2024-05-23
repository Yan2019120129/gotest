package main

import (
	"github.com/goccy/go-json"
	"github.com/gocolly/colly/v2"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"gotest/common/module/cache"
	"gotest/common/module/logs"
)

const (
	serverAddr = "https://www.amazon.com/"
	addr       = "https://www.amazon.com/Levis-Womens-Oxnard-Choice-Medium-Indigo/dp/B096SZ2X2R?ref_=Oct_DLandingS_D_ea28edf4_2&th=1"
)

func main() {
	rdsConn := cache.RdsPool.Get()
	defer rdsConn.Close()

	c := colly.NewCollector()

	// 模拟浏览器
	c.OnRequest(func(r *colly.Request) {
		value, err := redis.String(rdsConn.Do("HGET", "headers", "amazon"))
		_ = json.Unmarshal([]byte(value), r.Headers)
		if err != nil {
			logs.Logger.Error("OnRequest", zap.Error(err))
			r.Headers.Set("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
			r.Headers.Set("Accept", "*/*")
			r.Headers.Set("Host", "www.amazon.com")
			r.Headers.Set("Connection", "keep-alive")
			r.Headers.Set("Cookie", "session-id=134-7385922-0600243; session-id-time=2082787201l; i18n-prefs=USD; ubid-main=135-9468614-9654608; session-token=1nFruRr508lxlVmlE9ZwlHdQFmYVZDCStksKi8cGKybHPoqVLOsoVs1PBPMh4y7aM6L/loWH2rZRCc68mhu7grENU2TfRfjtXnpqW8thFpE8MUBS51gsj6Iy/9ReSaN0ha4Xh3+jvyqCRqJjXBcHnzTDn4wWSrB5i5YJwxgmfGM/dIacUrbWmQxLXfwCQ7vRFR2+LC3zf6nCjCaBr2jh5hDzPArkzzLeqNKmAjsK1IhnBho4YAxEGmHgnpfMozWC6Cz0xm4b7yFFYWHw5hpyPqFOO3GMNmNtLLjDZr2o7Xo04IKgo5JciWZv9z3Erz55cMQGAwDOWtxHedhmqUVwig9dkZvaWW+f; skin=noskin")
			r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")
			r.Headers.Set("Origin", "https://www.amazon.sg")
			r.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.120 Safari/537.36")
		}

		c.Cookies(r.URL.String())
	})

	c.OnResponse(func(response *colly.Response) {
		value, err := json.Marshal(response.Headers)
		if err != nil {
			logs.Logger.Error("OnResponse", zap.Error(err))
			return
		}
		err = rdsConn.Send("HSET", "headers", "amazon", value)
		if err != nil {
			logs.Logger.Error("OnResponse", zap.Error(err))
		}
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEach("#title", func(i int, element *colly.HTMLElement) {
			title := element.Text
			logs.Logger.Info("OnHTML", zap.Bool("title", element.Text != ""), zap.String("title", title))
		})

		e.ForEach("#corePrice_desktop > div > table > tbody > tr > td.a-span12 > span.a-price-range > span:nth-child(1)", func(i int, element *colly.HTMLElement) {
			price := element.Text
			logs.Logger.Info("OnHTML", zap.Bool("price", element.Text != ""), zap.String("price", price))
		})

		e.ForEach("#a-popover-3 > div > div > ul > li", func(i int, element *colly.HTMLElement) {
			size := element.Text
			logs.Logger.Info("OnHTML", zap.Bool("size", element.Text != ""), zap.String("size", size))
		})

		e.ForEach("#variation_color_name > ul > li", func(i int, element *colly.HTMLElement) {
			color := element.ChildAttr("img", "src")
			logs.Logger.Info("OnHTML", zap.Bool("color", element.Text != ""), zap.String("color", color))
		})

		e.ForEach("#altImages > ul > li", func(i int, element *colly.HTMLElement) {
			img := element.ChildAttr("img", "src")
			logs.Logger.Info("OnHTML", zap.Bool("color", img != ""), zap.String("img", img))
		})
	})

	// 响应错误
	c.OnError(func(rsp *colly.Response, err error) {
		logs.Logger.Error("OnError", zap.ByteString("body", rsp.Body), zap.Error(err))
	})

	// 在所有OnHTML后执行
	c.OnScraped(func(r *colly.Response) {
		logs.Logger.Info("OnScraped")
	})

	// 发送请求
	if err := c.Visit(addr); err != nil {
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
	}

	// 等待渲染完成
	c.Wait()
}

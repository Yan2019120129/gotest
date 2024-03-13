package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
	"go.uber.org/zap"
	"gotest/common/module/logs"
	"log"
	"net/http"
)

// const addr = "https://www.amazon.com/"
const addr = "http://www.dhgate.com"

//const addr = "https://gorm.io/zh_CN/docs/create.html"

func main() {

	c := colly.NewCollector()

	//// Instantiate default collector
	//c := colly.NewCollector(colly.AllowURLRevisit())

	// 表示抓取时异步的
	c.Async = true

	// Rotate two socks5 proxies
	rp, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:1337", "socks5://127.0.0.1:1338")
	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)

	c.WithTransport(&http.Transport{
		DisableKeepAlives: true,
	})
	// 模拟浏览器
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")
		r.Headers.Set("Origin", "https://www.dhgate.com")
		r.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.120 Safari/537.36") // 模拟浏览器访问
	})

	// 设置Cookie
	if err := c.SetCookies(addr, []*http.Cookie{
		{Name: "cookieName", Value: "cookieValue"},
	}); err != nil {
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
		return
	}
	c.Cookies(addr)

	c.OnHTML("body", func(e *colly.HTMLElement) {
		fmt.Println("body:", e.Text)
	})

	c.OnError(func(rsp *colly.Response, err error) {
		log.Println("rsp", string(rsp.Body))
		log.Println("Something went wrong:", err)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// 结束
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	if err := c.Visit(addr); err != nil {
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
		return
	}
	//param := url.Values{"dspm": {"pcen.ranking.moreranking_sale.toprank-1.lMwG5TfgPl67FMwOPDop"}, "resource_id": {"926012915"}}
	//err := c.Request("Get", addr, strings.NewReader(param.Encode()), colly.NewContext(), nil)
	//if err != nil {
	//	logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
	//}

	// 采集等待结束
	c.Wait()
}

package main

import (
	"github.com/goccy/go-json"
	"github.com/gocolly/colly/v2"
	"github.com/gomodule/redigo/redis"
	"gotest/common/module/cache"
	"log"
)

const (
	serverAddr = "https://www.amazon.com/"
	addr       = "https://www.amazon.com/SAMSUNG-Smartphone-Unlocked-Android-Processor/dp/B0CMDPRN7M/ref=sr_1_7?crid=18UMAKQ6QSB0V&dib=eyJ2IjoiMSJ9.THp_WLIyFX2-n1VQorYio1deNbWqTeg2EwxjBOqjecRZ08NJmUJDoW-QxnEmw2g8J-xDVG-AW_FVkvH3S2RW8i70jOZGEdUoYX8Jf5BZ-oDlbxVE4mg3ceH4h-u1Dfx6HNvqasykovjQGHHCHxPtTvoWZv4sd23mbtqBO2pMHZ0djbWDlyzHnXFliBzG5eBg3A3H6877Pz-YnSNO4NSIbu0-Gzctv0KcWer0heJ-jm8.gLDA4ma_UGO5TrVK2e_OuDNlrH40VwVrA9nTj5sY9QQ&dib_tag=se&keywords=phone&qid=1716515618&sprefix=phon%2Caps%2C721&sr=8-7&th=1"
)

func main() {
	rdsConn := cache.RdsPool.Get()
	defer rdsConn.Close()

	c := colly.NewCollector()

	// 设置请求头
	c.OnRequest(func(r *colly.Request) {
		value, err := redis.String(rdsConn.Do("HGET", "headers", "amazon"))
		_ = json.Unmarshal([]byte(value), r.Headers)
		cookie := r.Headers.Get("Set-Cookie")
		if cookie != "" {
			r.Headers.Set("Cookie", cookie)
		}
		if err != nil {
			log.Println("OnRequest", err)
			r.Headers.Set("Accept", "*/*")
			r.Headers.Set("Host", "www.amazon.com")
			r.Headers.Set("Connection", "keep-alive")
			r.Headers.Set("Cookie", "session-id=134-7385922-0600243; session-id-time=2082787201l; i18n-prefs=USD; ubid-main=135-9468614-9654608; session-token=1nFruRr508lxlVmlE9ZwlHdQFmYVZDCStksKi8cGKybHPoqVLOsoVs1PBPMh4y7aM6L/loWH2rZRCc68mhu7grENU2TfRfjtXnpqW8thFpE8MUBS51gsj6Iy/9ReSaN0ha4Xh3+jvyqCRqJjXBcHnzTDn4wWSrB5i5YJwxgmfGM/dIacUrbWmQxLXfwCQ7vRFR2+LC3zf6nCjCaBr2jh5hDzPArkzzLeqNKmAjsK1IhnBho4YAxEGmHgnpfMozWC6Cz0xm4b7yFFYWHw5hpyPqFOO3GMNmNtLLjDZr2o7Xo04IKgo5JciWZv9z3Erz55cMQGAwDOWtxHedhmqUVwig9dkZvaWW+f; skin=noskin")
			r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")
			r.Headers.Set("Origin", "https://www.amazon.com")
			r.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.120 Safari/537.36")
		} else {
			log.Println("Cookie", r.Headers.Get("Cookie"))
		}
	})
	c.OnResponse(func(response *colly.Response) {
		log.Println("Cookie", response.Headers.Get("Cookie"))
		value, err := json.Marshal(response.Headers)
		if err != nil {
			log.Println("OnResponse", err)
			return
		}
		err = rdsConn.Send("HSET", "headers", "amazon", value)
		if err != nil {
			log.Println("OnResponse", err)
		}
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEach("#title", func(i int, element *colly.HTMLElement) {
			title := element.Text
			log.Println("OnHTML", "title", title)
		})

		e.ForEach("#corePrice_desktop > div > table > tbody > tr > td.a-span12 > span.a-price-range > span:nth-child(1)", func(i int, element *colly.HTMLElement) {
			price := element.Text
			log.Println("OnHTML", "price", price)
		})

		e.ForEach("#a-popover-3 > div > div > ul > li", func(i int, element *colly.HTMLElement) {
			size := element.Text
			log.Println("OnHTML", "size", size)
		})

		e.ForEach("#variation_color_name > ul > li", func(i int, element *colly.HTMLElement) {
			color := element.ChildAttr("img", "src")
			log.Println("OnHTML", "color", color)
		})

		e.ForEach("#altImages > ul > li", func(i int, element *colly.HTMLElement) {
			img := element.ChildAttr("img", "src")
			log.Println("OnHTML", "img", img)
		})
	})

	// 响应错误
	c.OnError(func(rsp *colly.Response, err error) {
		log.Println("OnError", err, "body", string(rsp.Body), "ERROR", err)
	})

	// 在所有OnHTML后执行
	c.OnScraped(func(r *colly.Response) {
		log.Println("OnScraped")
	})

	// 发送请求
	if err := c.Visit(addr); err != nil {
		log.Println("Visit", err)
	}

	// 等待渲染完成
	c.Wait()
}

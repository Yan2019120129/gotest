package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

const url = "https://kyfw.12306.cn/otn/leftTicket/init?linktypeid=dc&fs=%E9%BE%99%E5%B2%A9,LYS&ts=%E7%8E%89%E6%9E%97,YLZ&date=2024-01-23&flag=N,N,Y"

func main() {
	c := colly.NewCollector()

	c.OnHTML("body", func(e *colly.HTMLElement) {
		fmt.Println("body:", e.Text)
	})
	// Find and visit all links
	c.OnHTML("#float > th:nth-child(11)", func(e *colly.HTMLElement) {
		fmt.Println("text:", e.Text)
		//e.Request.Visit()
	})

	//fmt.Println("html:", c.OnResponse())
	//c.OnRequest(func(r *colly.Request) {
	//	fmt.Println("Visiting", r.)
	//})

	if err := c.Visit(url); err != nil {
		return
	}
}

package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
	"gotest/common/module/logs"
	"net/http"
	"net/url"
	"strings"
)

const (
	// 日志信息标识
	logMsg = "crawling"

	// ObtainTypeOne 获取单一产品
	ObtainTypeOne = 1

	// ObtainTypeSearch 获取搜索产品
	ObtainTypeSearch = 2
)

var addrs = []string{
	"https://www.amazon.com/-/es/Wrangler-Authentics-Pantalones-el%C3%A1sticos-polvorienta/dp/B07YP9LTYF/?_encoding=UTF8&pd_rd_w=dV7JV&content-id=amzn1.sym.9929d3ab-edb7-4ef5-a232-26d90f828fa5&pf_rd_p=9929d3ab-edb7-4ef5-a232-26d90f828fa5&pf_rd_r=ET1V3PPQ0DE8ZJV0PAWR&pd_rd_wg=NZmEe&pd_rd_r=9b6df6db-31d4-4ee1-981a-c2c59a9f8e3f&ref_=pd_hp_d_btf_crs_zg_bs_7141123011",
	"https://www.amazon.com/-/en/interactivo-mascotas-electr%C3%B3nicas-virtuales-reaccionan/dp/B0BW1B31SR/ref=sr_1_5?_encoding=UTF8&content-id=amzn1.sym.44da4965-9668-4613-bec2-a3a75f0c2ad4&dib=eyJ2IjoiMSJ9.b8n9r6_6RGcDrh2OyrQ8rq9KqGl45pWLWgfqQLxww4g6NMQf59WWhrc3edi2tgBpAzI-KqtZSctfqBqqgIlRsx2SErh1Y4n2mWaQKdREXS2sHn_RNFIbIl78zmd6k05Cx9nmTtGag-_eCJNZvINdx7wdXC_MPz9eIDaZMbNMH5jdldvLyvtB74dNGB2XTdqFXcNg1m7kIZWw7KR4RbLOJsTaflpTZGpB6vWP5rAHFr9Yhs2lGm34m-7dysAz-18FHTXWikcUmHPLrBfRbLvyt5IkNVeLvk1V5r9r6yR2Vc8.iKFYnrMBMKtFMx0lTFIGSeUQIv8wdIEkTIs-IoDIW8E&dib_tag=se&keywords=toys&pd_rd_r=cfbeebeb-5643-4e4f-acf8-452d0bc8a5bf&pd_rd_w=xb5Ae&pd_rd_wg=GxofG&pf_rd_p=44da4965-9668-4613-bec2-a3a75f0c2ad4&pf_rd_r=WREZYKEXD9249MDRSWDR&qid=1717228485&refinements=p_36%3A-2500&sr=8-5",
	"https://www.amazon.com/-/en/dp/B0BJZ8WJ86/ref=sr_1_1?dib=eyJ2IjoiMSJ9.h4wcHhGqJDYWDNbwqcxOaz0HfEV_he2cxauBO-GVLAupv6jgmKgPWbg2enOIDIbnnzBsAJrhqMu6chMxmP9cVv30njqzaRfCJwWO3Xk4jqdvKCVCSuAB4IqtpCZ-Gj5poN3P9maYDkf0GoOemTEln_d7aeoPYZV4dFLvSj94fIL-EleIHnDQUmGOrAdL0YyEZXPqxiYr-UIz2qQi8eFPOSrRFgbaprQP0tySywNmUa75W-G9E87mOeTxZA2uBcm3ykg8e48UKLc7OPtlsZGGu6n-92DrJxkBwzRtrUq3cjQ.q7Y70ZhT_g7iM-DL4xs4KSxtwigzIdKsqEmJmSUPhBI&dib_tag=se&qid=1716542042&s=fashion-mens-intl-ship&sr=1-1",
	"https://www.amazon.com/SAMSUNG-Smartphone-Unlocked-Android-Processor/dp/B0CMDPRN7M/ref=sr_1_7?crid=18UMAKQ6QSB0V&dib=eyJ2IjoiMSJ9.THp_WLIyFX2-n1VQorYio1deNbWqTeg2EwxjBOqjecRZ08NJmUJDoW-QxnEmw2g8J-xDVG-AW_FVkvH3S2RW8i70jOZGEdUoYX8Jf5BZ-oDlbxVE4mg3ceH4h-u1Dfx6HNvqasykovjQGHHCHxPtTvoWZv4sd23mbtqBO2pMHZ0djbWDlyzHnXFliBzG5eBg3A3H6877Pz-YnSNO4NSIbu0-Gzctv0KcWer0heJ-jm8.gLDA4ma_UGO5TrVK2e_OuDNlrH40VwVrA9nTj5sY9QQ&dib_tag=se&keywords=phone&qid=1716515618&sprefix=phon%2Caps%2C721&sr=8-7&th=1",
	"https://www.amazon.com/-/en/dp/B0CHN7H2SL/ref=sr_1_1?dib=eyJ2IjoiMSJ9.mr4o2X0z_1tgbQNezvkZPK5kcImaQnJIikB2oylCs2bsxGnFxs8kDOkIHIjnw0M3_CRMN3kJNKZnrz5_5vlvGVx7Fo1oMJrjN7EOHzla2Gg-93ucVwaOs-qAQZ6gHlNMgSApX3ThIh1y3ERk4vqD3yYQKM-0Y8YB-SoJu8i5RrdMeaBCbcuG7UrO1zKwNvRFbZO2JlPVyU20zpsrF8XAix8SstWJ06sRHrePz-94yYk.KZfp_Ca7FoWZ71fZG3valXYBt9AeU95PmhvznZtE9SU&dib_tag=se&qid=1716543489&s=software-intl-ship&sr=1-1",
	"https://www.amazon.com/-/en/dp/B09MKNL9M3/ref=sr_1_19?dib=eyJ2IjoiMSJ9.wKU47KNjMfIsPT6vKaGMHsRSyB-SUU10SErmfR-hiQwf31tSYA6UG7GLjWzYayHt7KKmlmNCNaxJ9V1qhdHCH-46xyZ8QkhQXKToGUbaKf3nvJnk8BVpTYJ35jBvwZoxpjIohNBUhQ1zHOxib6gR8rp92EV3S7ReWokBScz-VVY_maXpMQX9uxWfWJdwSxGKmGZwxvqFU5wFRStQv-ZqMlJaO21YBN29aSW_wc_ljIwEsyE1YRJEmAGIDNaEZfRmms2lbaUYE-Rh_3PbcSx2M965jFnFZ67V8NAtP1zCQQA.E3Q1Nz4cwRGZQwPNG_XeN0t6FSn9-QEKrw2pCVHj67Q&dib_tag=se&qid=1716544699&s=fashion-womens-intl-ship&sr=1-19&th=1",
	"https://www.amazon.com/-/en/dp/B07WN1FKS3/ref=sr_1_4?dib=eyJ2IjoiMSJ9.iMl2wKNffcNoehjuH_QlWAik428wpZp697-BzRqRti713zOuTmkQQEKvMk2cpGOQTdIGEz9skxbMYHyVW3ob8tzIc93bywcspo3Sx73K18nS-f39Tkkt7-GyRu2oKC4arh_Tc4hyIshVZGnmRbYPOFS0KmQ9Re7wpW27TUm9EV5PFtDh57-2LQMJG9caDt2gil_6MeV_JBLWrjNtVRjKCLeTEoqzJFr1nk-LZwfeRABQUTIbstrCckO4zyaHIoXEWLsmPfgT-EJ4GilkmwTgpenHSmk0dTYyVhUbuZXSbnE.mtkVzBWqHXWIFjy22rl6tZ_gWeWGQ77OB96sYlFoW5c&dib_tag=se&qid=1716553753&s=apparel&sr=1-4&th=1",
	"https://www.amazon.com/-/en/dp/B08152D6DL?ref_=Oct_DLandingS_D_948c90c7_0&th=1",
	"https://www.amazon.com/-/zh/dp/B072YVWBXH/ref=sr_1_6?dib=eyJ2IjoiMSJ9.ugc4h7Cjckige0OnPnfm0Q9c4L70fU-d3NXpOEc2bpHAsBlRuerqp0-crUeqlOcogxgdQAFeWH18BxkUxUVMqW6n8i4E-XnEraYiaGGUCBJOuUW69SQXKpN7bCjsyR0RK4bUTgQlZ6SWuf41CT15-rcezjYpzA8nSPX94oEfl97DQ_ycLDjvDUduWj-cFra2.Ia_a0Me33aSZ9ir5fj8gPnNNE-LSh-PqYcDEPZmYfPg&dib_tag=se&qid=1716557092&s=hpc-intl-ship&sr=1-6",
	"https://www.amazon.com/-/en/dp/B09P1DV7D8/ref=sr_1_7?_encoding=UTF8&content-id=amzn1.sym.36b3bd24-30cc-4394-a99c-32175deb1058&dib=eyJ2IjoiMSJ9.dhz8NIOr3fJI_Y_BZn2dio9iMwsxkxh92SLmNsJQ6BQbkcnabGX1H0kNpIfLWKVMx9z7vcLaUipPjPCx1kJipr-TKIdLX8qtObfswj2DMsoT4cxLB10Pjd1U1Aj-2UDJdRGn9CEYedsHlPWE98uvaaVf-0irxC0nLtQDHvlUmKYYV8FU9bkcgt8kH57BAR5n4YxaW1uNF6eCtdNf_8K1FCVNJY94kXMtObD5h0UmXNpdPJuY8oDqUdOl5oeGkKwoFI4xiZEaoEhyiknrDpF9_Sp4Zco3ecRuRClbpfs5m5A.Xt1scjg5sprI3h4pMUhmBm-Jv3bueGe2LDnTHB2kaEk&dib_tag=se&keywords=tech%2Bgifts%2Bfor%2Bmen&pd_rd_r=bea00079-f6b2-4e29-8ca3-6b322d8d4b85&pd_rd_w=z4Vk7&pd_rd_wg=LCOET&pf_rd_p=36b3bd24-30cc-4394-a99c-32175deb1058&pf_rd_r=F03EF7M4YHKFSFD6W28B&qid=1716628148&sr=8-7&th=1",
}

var searchAddrs = []string{
	"https://www.amazon.com/s?i=fashion-mens-intl-ship&bbn=16225019011&rh=n%3A7141123011%2Cn%3A16225019011%2Cn%3A7147441011%2Cn%3A1040658&dc&language=zh&ds=v1%3A0qMeFWgElNp7zKRHk7%2FFniqSLbOE3bo%2FlTElqN3vB3s&qid=1716639236&ref=sr_ex_n_1",
	"https://www.amazon.com/s?k=outdoor+sports&i=fashion-mens&__mk_es_US=%C3%85M%C3%85%C5%BD%C3%95%C3%91&crid=21PNQMXB0BHU0&sprefix=outdoor+sports+exercise+pedal%2Cfashion-mens%2C1163&ref=nb_sb_noss_1",
}

type CrawlingParams struct {
	Type       int    `validate:"required"` //  获取产品类型 1单一产品，2整个页面产品
	PageNum    int    `validate:"required"` //  获取多少页
	Limit      int    //  获取多少条
	CategoryId uint   `validate:"required,gt=0"` //  类目ID
	Url        string `validate:"required"`      //  产品URL
}

var params = []*CrawlingParams{
	{CategoryId: 2, Url: "https://www.amazon.com/-/es/s?k=womens+summer+tops&page=2&crid=29FS0PD8YQ5DT&qid=1716797458&sprefix=women%2Caps%2C593&ref=sr_pg_2", Type: 2, PageNum: 10, Limit: 0},
}

func main() {
	amazon := NewAmazon().SetObtainMassage(func(attr *ProductAttr) {
		logs.Logger.Info(logMsg, zap.Reflect("message", attr))
	}).Run(addrs[1])
	fmt.Println(amazon)
}

type Amazon struct {
	collector *colly.Collector
	url       *url.URL
	onMassage func(productAttr *ProductAttr)
}

// NewAmazon 新建实例
func NewAmazon() *Amazon {
	amazon := &Amazon{
		collector: colly.NewCollector(),
	}
	amazon.collector.Async = true
	amazon.collector.AllowURLRevisit = true

	amazon.collector.WithTransport(&http.Transport{
		DisableKeepAlives: true,
	})

	storage := &Storage{
		Address:  "127.0.0.1:6379",
		Password: "",
		DB:       0,
		Prefix:   "job01",
	}

	err := amazon.collector.SetStorage(storage)
	if err != nil {
		panic(err)
	}

	// 设置请求头
	amazon.collector.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("message", &ProductAttr{})
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Host", "www.amazon.com")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.9")
		r.Headers.Set("Origin", "https://www.amazon.com")
		r.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.120 Safari/537.36")
	})

	// 处理响应
	amazon.collector.OnResponse(func(resp *colly.Response) {

	})

	// 处理页面响应元素
	amazon.collector.OnHTML(amazon.obtainDetails("body"))

	// 响应错误
	amazon.collector.OnError(func(rsp *colly.Response, err error) {
		logs.Logger.Error(logMsg, zap.Error(err), zap.String("body", string(rsp.Body)))
	})

	// 爬取完成后执行
	amazon.collector.OnScraped(func(resp *colly.Response) {
		if p, ok := resp.Ctx.GetAny("message").(*ProductAttr); ok {
			amazon.onMassage(p)
		}
	})

	return amazon
}

// Run 执行程序
func (_Amazon *Amazon) Run(url string) error {
	// 监听html
	err := _Amazon.collector.Visit(url)
	_Amazon.collector.Wait()
	return err
}

// obtainDetails 获取产品信息
func (_Amazon *Amazon) obtainDetails(selector string) (string, func(e *colly.HTMLElement)) {
	return selector, func(e *colly.HTMLElement) {
		productInfo := NewProductAttr()

		// 获取标题
		productInfo.SetTitle(e.ChildText("#title"))

		// 获取产品金额
		e.ForEach("#corePriceDisplay_desktop_feature_div > div.aok-align-center ", func(i int, element *colly.HTMLElement) {
			priceText := ""
			if priceText = element.ChildText("span.aok-offscreen"); priceText == "" {
				priceText = element.ChildText("span.a-offscreen")
			}

			if priceText != "" {
				productInfo.SetPrice(priceText)
			}
		})

		e.ForEach("#corePrice_desktop > div > table > tbody > tr > td.a-span12 > span ", func(i int, element *colly.HTMLElement) {
			element.ForEach("span.a-offscreen", func(i int, element *colly.HTMLElement) {
				if priceText := element.Text; priceText != "" {
					productInfo.SetPrice(priceText)
				}
			})
		})

		// 获取产品图片
		e.ForEach("#altImages > ul > li", func(i int, element *colly.HTMLElement) {
			imageUrl := element.ChildAttr("img", "src")
			if !strings.Contains(imageUrl, ".gif") && !strings.Contains(imageUrl, "PKdp-play-icon-overlay") {
				if imageUrl != "" && len(productInfo.Images) < 5 {
					productInfo.SetImages(imageUrl)
				}
			}
		})

		// 获取产品规格  颜色，大小，样式等等
		e.ForEach("#native_dropdown_selected_size_name > option", func(i int, element *colly.HTMLElement) {
			label := element.ChildText("#native_size_name_-1")
			alt := element.ChildText("option.dropdownAvailable")
			if label != "" && alt != "" {
				stylLen := productInfo.GetStyleLen(label)
				if stylLen < 5 {
					productInfo.SetStyle(label, alt)
				}
			}
		})

		// 产品属性
		e.ForEach("#twister > div", func(i int, element *colly.HTMLElement) {
			label := element.ChildText("label.a-form-label")
			element.ForEach("select > option", func(i int, element *colly.HTMLElement) {
				alt := element.Text
				if label != "" && alt != "" {
					stylLen := productInfo.GetStyleLen(label)
					if stylLen < 5 {
						productInfo.SetStyle(label, alt)
					}
				}
			})

			element.ForEach("ul > li ", func(i int, element *colly.HTMLElement) {
				alt := ""
				if alt = element.ChildAttr("img", "alt"); alt == "" {
					alt = element.ChildText("button > div > div.twisterTextDiv.text")
				}
				if label != "" && alt != "" {
					stylLen := productInfo.GetStyleLen(label)
					if stylLen < 5 {
						productInfo.SetStyle(label, alt)
					}
				}
			})

			if element.DOM.Find("select").Length() == 0 && element.DOM.Find("ul").Length() == 0 {
				alt := element.ChildText("span.selection")
				productInfo.SetStyle(label, alt)
			}
		})
		e.ForEach("#poExpander > div.a-expander-content.a-expander-partial-collapse-content > div > table > tbody > tr.a-spacing-small.po-brand", func(i int, element *colly.HTMLElement) {
			label := element.ChildText("td > span.a-size-base.a-text-bold")
			alt := element.ChildText("td > span.a-size-base.po-break-word")
			productInfo.SetStyle(label, alt)
		})

		e.ForEach("#feature-bullets", func(i int, element *colly.HTMLElement) {
			doc := element.DOM
			html, err := doc.Find("ul").Html()
			if err != nil {
				return
			}
			logs.Logger.Info(logMsg, zap.String("feature-bullets", html))
		})

		// 获取产品描述
		if productInfo.Describe = e.DOM.Find("#productFactsDesktopExpander > div > ul").Text(); productInfo.Describe == "" {
			if productInfo.Describe = e.DOM.Find("#feature-bullets > ul").Text(); productInfo.Describe == "" {
				productInfo.Describe = e.DOM.Find("#productDescription").Text()
			}
		}

		e.Response.Ctx.Put("message", productInfo)
	}
}

// SetObtainMassage 处理消息方法
func (_Amazon *Amazon) SetObtainMassage(onMassage func(attr *ProductAttr)) *Amazon {
	_Amazon.onMassage = onMassage
	return _Amazon
}

// getHttpUrl 获取http服务器地址
func (_Amazon *Amazon) getHttpUrl() string {
	return _Amazon.url.Scheme + "://" + _Amazon.url.Host
}

// setHttpUrl 设置 http服务器地址
func (_Amazon *Amazon) setHttpUrl(rawUrl string) error {
	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		return err
	}
	_Amazon.url = parsedURL
	return nil
}

// isHttpHeader 是否包含http请求头
func (_Amazon *Amazon) isHttpHeader(url string) bool {
	httpSymbol := []byte("http")
	httpLen := len(httpSymbol)
	if len(url) < httpLen {
		return false
	}
	for i, v := range httpSymbol {
		if url[i] == v && i == httpLen-1 {
			return true
		}
		if url[i] != v {
			return false
		}
	}
	return false
}

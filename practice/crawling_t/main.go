package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"gotest/practice/okex/utils"
	"log"
	"strings"
)

const saveImagePath = "./practice/crawling_t/download/"

var addr = []string{"https://www.amazon.com/-/en/dp/B0BJZ8WJ86/ref=sr_1_1?dib=eyJ2IjoiMSJ9.h4wcHhGqJDYWDNbwqcxOaz0HfEV_he2cxauBO-GVLAupv6jgmKgPWbg2enOIDIbnnzBsAJrhqMu6chMxmP9cVv30njqzaRfCJwWO3Xk4jqdvKCVCSuAB4IqtpCZ-Gj5poN3P9maYDkf0GoOemTEln_d7aeoPYZV4dFLvSj94fIL-EleIHnDQUmGOrAdL0YyEZXPqxiYr-UIz2qQi8eFPOSrRFgbaprQP0tySywNmUa75W-G9E87mOeTxZA2uBcm3ykg8e48UKLc7OPtlsZGGu6n-92DrJxkBwzRtrUq3cjQ.q7Y70ZhT_g7iM-DL4xs4KSxtwigzIdKsqEmJmSUPhBI&dib_tag=se&qid=1716542042&s=fashion-mens-intl-ship&sr=1-1",
	"https://www.amazon.com/SAMSUNG-Smartphone-Unlocked-Android-Processor/dp/B0CMDPRN7M/ref=sr_1_7?crid=18UMAKQ6QSB0V&dib=eyJ2IjoiMSJ9.THp_WLIyFX2-n1VQorYio1deNbWqTeg2EwxjBOqjecRZ08NJmUJDoW-QxnEmw2g8J-xDVG-AW_FVkvH3S2RW8i70jOZGEdUoYX8Jf5BZ-oDlbxVE4mg3ceH4h-u1Dfx6HNvqasykovjQGHHCHxPtTvoWZv4sd23mbtqBO2pMHZ0djbWDlyzHnXFliBzG5eBg3A3H6877Pz-YnSNO4NSIbu0-Gzctv0KcWer0heJ-jm8.gLDA4ma_UGO5TrVK2e_OuDNlrH40VwVrA9nTj5sY9QQ&dib_tag=se&keywords=phone&qid=1716515618&sprefix=phon%2Caps%2C721&sr=8-7&th=1",
	"https://www.amazon.com/-/en/dp/B0CHN7H2SL/ref=sr_1_1?dib=eyJ2IjoiMSJ9.mr4o2X0z_1tgbQNezvkZPK5kcImaQnJIikB2oylCs2bsxGnFxs8kDOkIHIjnw0M3_CRMN3kJNKZnrz5_5vlvGVx7Fo1oMJrjN7EOHzla2Gg-93ucVwaOs-qAQZ6gHlNMgSApX3ThIh1y3ERk4vqD3yYQKM-0Y8YB-SoJu8i5RrdMeaBCbcuG7UrO1zKwNvRFbZO2JlPVyU20zpsrF8XAix8SstWJ06sRHrePz-94yYk.KZfp_Ca7FoWZ71fZG3valXYBt9AeU95PmhvznZtE9SU&dib_tag=se&qid=1716543489&s=software-intl-ship&sr=1-1",
	"https://www.amazon.com/-/en/dp/B09MKNL9M3/ref=sr_1_19?dib=eyJ2IjoiMSJ9.wKU47KNjMfIsPT6vKaGMHsRSyB-SUU10SErmfR-hiQwf31tSYA6UG7GLjWzYayHt7KKmlmNCNaxJ9V1qhdHCH-46xyZ8QkhQXKToGUbaKf3nvJnk8BVpTYJ35jBvwZoxpjIohNBUhQ1zHOxib6gR8rp92EV3S7ReWokBScz-VVY_maXpMQX9uxWfWJdwSxGKmGZwxvqFU5wFRStQv-ZqMlJaO21YBN29aSW_wc_ljIwEsyE1YRJEmAGIDNaEZfRmms2lbaUYE-Rh_3PbcSx2M965jFnFZ67V8NAtP1zCQQA.E3Q1Nz4cwRGZQwPNG_XeN0t6FSn9-QEKrw2pCVHj67Q&dib_tag=se&qid=1716544699&s=fashion-womens-intl-ship&sr=1-19&th=1",
	"https://www.amazon.com/-/en/dp/B07WN1FKS3/ref=sr_1_4?dib=eyJ2IjoiMSJ9.iMl2wKNffcNoehjuH_QlWAik428wpZp697-BzRqRti713zOuTmkQQEKvMk2cpGOQTdIGEz9skxbMYHyVW3ob8tzIc93bywcspo3Sx73K18nS-f39Tkkt7-GyRu2oKC4arh_Tc4hyIshVZGnmRbYPOFS0KmQ9Re7wpW27TUm9EV5PFtDh57-2LQMJG9caDt2gil_6MeV_JBLWrjNtVRjKCLeTEoqzJFr1nk-LZwfeRABQUTIbstrCckO4zyaHIoXEWLsmPfgT-EJ4GilkmwTgpenHSmk0dTYyVhUbuZXSbnE.mtkVzBWqHXWIFjy22rl6tZ_gWeWGQ77OB96sYlFoW5c&dib_tag=se&qid=1716553753&s=apparel&sr=1-4&th=1",
	"https://www.amazon.com/-/en/dp/B08152D6DL?ref_=Oct_DLandingS_D_948c90c7_0&th=1",
	"https://www.amazon.com/-/zh/dp/B072YVWBXH/ref=sr_1_6?dib=eyJ2IjoiMSJ9.ugc4h7Cjckige0OnPnfm0Q9c4L70fU-d3NXpOEc2bpHAsBlRuerqp0-crUeqlOcogxgdQAFeWH18BxkUxUVMqW6n8i4E-XnEraYiaGGUCBJOuUW69SQXKpN7bCjsyR0RK4bUTgQlZ6SWuf41CT15-rcezjYpzA8nSPX94oEfl97DQ_ycLDjvDUduWj-cFra2.Ia_a0Me33aSZ9ir5fj8gPnNNE-LSh-PqYcDEPZmYfPg&dib_tag=se&qid=1716557092&s=hpc-intl-ship&sr=1-6",
}

func main() {
	amazon := NewAmazon()
	c := colly.NewCollector()

	// 设置请求头
	c.OnRequest(amazon.ConfigHeaders)

	// 处理响应
	c.OnResponse(amazon.ObtainResponse)

	c.OnHTML("body", amazon.ObtainHTML)

	// 响应错误
	c.OnError(func(rsp *colly.Response, err error) {
		log.Println("OnError", err, "body", string(rsp.Body), "ERROR", err)
	})

	// 在所有OnHTML后执行
	c.OnScraped(func(r *colly.Response) {
		log.Println("OnScraped")
	})

	// 发送请求
	if err := c.Visit(addr[5]); err != nil {
		log.Println("Visit", err)
	}

	// 等待渲染完成
	c.Wait()
	amazon.Print()
}

type Amazon struct {
	Price    string
	Style    map[string][]string
	Images   []string
	Title    string
	Describe string
}

func NewAmazon() *Amazon {
	return &Amazon{
		Price:    "",
		Style:    make(map[string][]string),
		Images:   make([]string, 0),
		Title:    "",
		Describe: "",
	}
}

// ConfigHeaders 配置请求头
func (_Amazon *Amazon) ConfigHeaders(r *colly.Request) {
	r.Headers.Set("Accept", "*/*")
	r.Headers.Set("Host", "www.amazon.com")
	r.Headers.Set("Connection", "keep-alive")
	r.Headers.Set("Cookie", "session-id=134-7385922-0600243; session-id-time=2082787201l; i18n-prefs=USD; ubid-main=135-9468614-9654608; session-token=1nFruRr508lxlVmlE9ZwlHdQFmYVZDCStksKi8cGKybHPoqVLOsoVs1PBPMh4y7aM6L/loWH2rZRCc68mhu7grENU2TfRfjtXnpqW8thFpE8MUBS51gsj6Iy/9ReSaN0ha4Xh3+jvyqCRqJjXBcHnzTDn4wWSrB5i5YJwxgmfGM/dIacUrbWmQxLXfwCQ7vRFR2+LC3zf6nCjCaBr2jh5hDzPArkzzLeqNKmAjsK1IhnBho4YAxEGmHgnpfMozWC6Cz0xm4b7yFFYWHw5hpyPqFOO3GMNmNtLLjDZr2o7Xo04IKgo5JciWZv9z3Erz55cMQGAwDOWtxHedhmqUVwig9dkZvaWW+f; skin=noskin")
	r.Headers.Set("Accept-Language", "en-US,en;q=0.9")
	r.Headers.Set("Origin", "https://www.amazon.com")
	r.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.120 Safari/537.36")
}

// ObtainPictures 获取详情图片
func (_Amazon *Amazon) ObtainPictures(i int, element *colly.HTMLElement) {
	imageUrl := element.ChildAttr("img", "src")
	if !strings.Contains(imageUrl, ".gif") && !strings.Contains(imageUrl, "PKdp-play-icon-overlay") {
		index := []int{}
		for sum, j := 0, len(imageUrl)-1; j > 0; j-- {
			replaceByte := uint8('_')
			original := imageUrl[j]
			if original == replaceByte {
				index = append(index, j)
				sum++
			}
			if sum == 2 {
				break
			}
		}
		if len(index) > 0 {
			imageUrl = imageUrl[:index[1]+1] + "SL1500" + imageUrl[index[0]:]
		}
		_Amazon.Images = append(_Amazon.Images, imageUrl)
		//if err := element.Request.Visit(imageUrl); err != nil {
		//	log.Println("Visit", err)
		//}
	}
}

// ObtainHTML 处理页面元素
func (_Amazon *Amazon) ObtainHTML(e *colly.HTMLElement) {
	// 获取标题
	e.ForEach("#title", _Amazon.ObtainTitle)

	// 获取产品金额
	e.ForEach("#corePriceDisplay_desktop_feature_div > div ", func(i int, element *colly.HTMLElement) {
		_Amazon.Price = element.ChildText("span.aok-offscreen")
	})
	if _Amazon.Price == "" {
		e.ChildText("#corePrice_desktop > div > table > tbody > tr > td.a-span12 > span.a-price.a-text-price.a-size-medium.apexPriceToPay > span.a-offscreen")
	}

	// 获取产品图片
	e.ForEach("#altImages > ul > li", _Amazon.ObtainPictures)

	// 获取产品规格  颜色，大小，样式等等
	e.ForEach("#twister > div", _Amazon.ObtainStyle)

	// 获取产品描述
	if _Amazon.Describe = e.ChildText("#productFactsDesktopExpander > div > ul"); _Amazon.Describe == "" {
		_Amazon.Describe = e.ChildText("#feature-bullets > ul")
	}
}

// ObtainPrice 获取产品价格
func (_Amazon *Amazon) ObtainPrice(i int, element *colly.HTMLElement) {

}

// ObtainTitle 获取产品标题
func (_Amazon *Amazon) ObtainTitle(i int, element *colly.HTMLElement) {
	_Amazon.Title = element.Text
}

// ObtainResponse 处理响应
func (_Amazon *Amazon) ObtainResponse(response *colly.Response) {
	contentType := response.Headers.Get("Content-Type")
	if strings.Contains(contentType, "image") {
		if !utils.PathExists(saveImagePath) {
			utils.PathMkdirAll(saveImagePath)
		}
		// 保存图片
		imgPath := saveImagePath + response.FileName()
		if !utils.PathExists(imgPath) {
			err := response.Save(imgPath)
			if err != nil {
				panic(err)
			}
		}
	}
}

// ObtainStyle 获取产品属性
func (_Amazon *Amazon) ObtainStyle(i int, element *colly.HTMLElement) {
	label := element.ChildAttr("#native_dropdown_selected_size_name", "data-a-touch-header")
	if label != "" {
		text := element.ChildText("option.dropdownAvailable")
		_Amazon.Style[label] = append(_Amazon.Style[label], text)
	}

	label = element.ChildText("label.a-form-label")
	element.ForEach("ul > li ", func(i int, element *colly.HTMLElement) {
		alt := element.ChildAttr("img", "alt")
		if alt == "" {
			alt = element.ChildText("button")
		}
		_Amazon.Style[label] = append(_Amazon.Style[label], alt)
	})
}

// Print 打印
func (_Amazon *Amazon) Print() {
	fmt.Println(_Amazon)
}

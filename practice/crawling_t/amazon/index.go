package main

import (
	"errors"
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
	"gotest/common/module/logs"
	"gotest/practice/okex/utils"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	// 主站网址
	httpsUrl = "https://www.amazon.com"

	// 媒体数据url
	httpsUrlMedia = "https://m.media-amazon.com"

	// 保存图片地址
	saveImagePath = "/Users/taozi/Documents/Golang/shop/public/crawling/describe"

	// 暴露给前台展示的路径
	exposedImagePath = "/crawling/describe"

	// 日志信息标识
	logMsg = "crawling"

	// ObtainTypeOne 获取单一产品
	ObtainTypeOne = 1

	// ObtainTypeSearch 获取搜索产品
	ObtainTypeSearch = 2
)

var addrs = []string{
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
	Urls       string `validate:"required"`      //  产品URL
}

var params = []*CrawlingParams{
	{CategoryId: 2, Urls: "https://www.amazon.com/-/es/s?k=womens+summer+tops&page=2&crid=29FS0PD8YQ5DT&qid=1716797458&sprefix=women%2Caps%2C593&ref=sr_pg_2", Type: 2, PageNum: 10, Limit: 0},
}

func main() {
	w := sync.WaitGroup{}
	for _, info := range params {
		NewAmazon().SetNextPageLimit(info.Limit).ObtainMassage(func(attr *ProductAttr) {
			logs.Logger.Info(logMsg, zap.Reflect("data", attr))
		}).Run(info.Urls, info.Type, info.Limit)
	}
	w.Wait()
}

type ProductAttr struct {
	Price    []float64
	Images   []string
	Style    map[string][]string
	Title    string
	Name     string
	Describe string
}

func NewProductAttr() *ProductAttr {
	return &ProductAttr{
		Price:  make([]float64, 0),
		Images: make([]string, 0),
		Style:  make(map[string][]string),
	}
}

// GetOriginalPrice 获取原价
func (_ProductInfo *ProductAttr) GetOriginalPrice() float64 {

	priceLen := len(_ProductInfo.Price)
	if priceLen >= 1 {
		return _ProductInfo.Price[priceLen-1]
	}
	return 0
}

// GetCurrentPrice 获取现价
func (_ProductInfo *ProductAttr) GetCurrentPrice() float64 {
	if len(_ProductInfo.Price) >= 1 {
		return _ProductInfo.Price[0]
	}
	return 0
}

// SetStyle 获取现价
func (_ProductInfo *ProductAttr) SetStyle(key, value string) {
	key = strings.Replace(key, ":", "", 1)
	_ProductInfo.Style[key] = append(_ProductInfo.Style[key], value)
}

// GetStyleLen 获取现价
func (_ProductInfo *ProductAttr) GetStyleLen(key string) int {
	key = strings.Replace(key, ":", "", 1)
	return len(_ProductInfo.Style[key])
}

type Amazon struct {
	*colly.Collector
	sync.Mutex
	err           error
	onMassage     func(productAttr *ProductAttr)
	limit         int // 限制产品信条数
	nextPageLimit int // 爬取页数限制
	pageNumber    int // 爬取页数
}

// NewAmazon 新建实例
func NewAmazon() *Amazon {
	return &Amazon{
		Mutex:     sync.Mutex{},
		Collector: colly.NewCollector(),
	}
}

// Run 执行程序
func (_Amazon *Amazon) Run(url string, obtainType, limit int) *Amazon {
	// 设置请求头
	_Amazon.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Host", "www.amazon.com")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Cookie", "session-id=134-7385922-0600243; session-id-time=2082787201l; i18n-prefs=USD; ubid-main=135-9468614-9654608; session-token=dosOPPC26xPdOsT/xbDzeLv74Z3wrazNgpGdCYSxs4Xa2KI+eSXfvO+yec1bK6ORV3sbSWgfIFOKCijsBt8C7GTtRnkWBwG4UyqoB2wI2+qqnYKotpK5lZ6EVbgUbbGUFS+oTF0C2FP/rgfazU7dQtKoe5/hAX2TX96mnMwGvnDf4RigeDrh05I5/9NZPbaJ1JDzZrnVl8f13o1Ahx26nR0dvic/3RRfVQAoyHSfkDkt5lwbTMn/xQZe+97LsIoVPd5pX+1Ok/+Zi+hjngYYx0OVqTtruTOQmxHUGXkIWobfURKFyZYd0swZkEQLetuox0hAA46qzgAhKX5V5TRm+/keRSG7Idni; skin=noskin")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.9")
		r.Headers.Set("Origin", "https://www.amazon.com")
		r.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.120 Safari/537.36")
	})

	// 响应错误
	_Amazon.OnError(func(rsp *colly.Response, err error) {
		logs.Logger.Error(logMsg, zap.Error(err), zap.String("body", string(rsp.Body)))
		_Amazon.err = err
	})

	_Amazon.limit = limit

	// 判断是否查找整个分类
	switch obtainType {
	case ObtainTypeOne:
		_Amazon.ObtainDetails(url)
	case ObtainTypeSearch:
		_Amazon.searchClassLink(url)
	}

	return _Amazon
}

// ObtainDetails 获取产品信息
func (_Amazon *Amazon) ObtainDetails(url string) *Amazon {
	productInfo := NewProductAttr()

	// 处理响应
	_Amazon.OnResponse(func(response *colly.Response) {
		contentType := response.Headers.Get("Content-Type")
		if strings.Contains(contentType, "image") {
			savePath := saveImagePath + "/"
			if !utils.PathExists(savePath) {
				utils.PathMkdirAll(savePath)
			}

			// 保存图片
			imgPath := savePath + response.FileName()
			if !utils.PathExists(imgPath) {
				err := response.Save(imgPath)
				if err != nil {
					logs.Logger.Error(logMsg, zap.Error(err))
					_Amazon.err = err
					return
				}
			}

			productInfo.Images = append(productInfo.Images, exposedImagePath+"/"+response.FileName())
		}
	})

	// 处理页面数据
	_Amazon.OnHTML("#ppd", func(e *colly.HTMLElement) {
		// 获取标题
		e.ForEach("#title", func(i int, element *colly.HTMLElement) {
			productInfo.Title = strings.TrimSpace(element.Text)
		})

		// 获取产品金额
		e.ForEach("#apex_desktop", func(i int, element *colly.HTMLElement) {
			// 样式一
			element.ForEach("#corePriceDisplay_desktop_feature_div > div.aok-align-center ", func(i int, element *colly.HTMLElement) {
				priceText := ""
				if priceText = element.ChildText("span.a-size-small.aok-offscreen"); priceText == "" {
					priceText = element.ChildText("span.a-offscreen")
				}
				if priceText != "" {
					productInfo.Price = append(productInfo.Price, _Amazon.stringToNumber(priceText))
				}
			})

			// 样式二
			element.ForEach("#corePrice_desktop > div > table > tbody > tr > td.a-span12 > span ", func(i int, element *colly.HTMLElement) {
				if priceText := element.ChildText("span.a-offscreen"); priceText != "" {
					productInfo.Price = append(productInfo.Price, _Amazon.stringToNumber(priceText))
				}
			})
		})

		// 获取产品图片
		e.ForEach("#altImages > ul > li", func(i int, element *colly.HTMLElement) {
			imageUrl := element.ChildAttr("img", "src")
			if !strings.Contains(imageUrl, ".gif") && !strings.Contains(imageUrl, "PKdp-play-icon-overlay") {
				if imageUrl != "" && len(productInfo.Images) < 5 {
					img := _Amazon.convertImage(imageUrl)
					if err := element.Request.Visit(img); err != nil {
						// 如果重复图片则再入本地地址
						if errors.Is(err, colly.ErrAlreadyVisited) {
							productInfo.Images = append(productInfo.Images, exposedImagePath+"/"+_Amazon.copyConvertImage(img))
							return
						}
						_Amazon.err = err
					}
				}
			}
		})

		// 获取产品规格  颜色，大小，样式等等
		e.ForEach("#native_dropdown_selected_size_name > option", func(i int, element *colly.HTMLElement) {
			label := element.ChildText("#native_size_name_-1")
			alt := element.ChildText("option.dropdownAvailable")
			if label != "" && alt != "" {
				_Amazon.Lock()
				stylLen := productInfo.GetStyleLen(label)
				if stylLen < 5 {
					productInfo.SetStyle(label, alt)
				}
				_Amazon.Unlock()
			}
		})
		e.ForEach("#twister > div", func(i int, element *colly.HTMLElement) {
			label := element.ChildText("label.a-form-label")
			element.ForEach("select > option", func(i int, element *colly.HTMLElement) {
				alt := element.Text
				if label != "" && alt != "" {
					_Amazon.Lock()
					stylLen := productInfo.GetStyleLen(label)
					if stylLen < 5 {
						productInfo.SetStyle(label, alt)
					}
					_Amazon.Unlock()
				}
			})

			element.ForEach("ul > li ", func(i int, element *colly.HTMLElement) {
				alt := ""
				if alt = element.ChildAttr("img", "alt"); alt == "" {
					alt = element.ChildText("button > div > div.twisterTextDiv.text")
				}
				if label != "" && alt != "" {
					_Amazon.Lock()
					stylLen := productInfo.GetStyleLen(label)
					if stylLen < 5 {
						productInfo.SetStyle(label, alt)
					}
					_Amazon.Unlock()
				}
			})
		})

		// 获取产品描述
		if productInfo.Describe = e.ChildText("#productFactsDesktopExpander > div > ul"); productInfo.Describe == "" {
			productInfo.Describe = e.ChildText("#feature-bullets > ul")
		}

	})

	// 发送请求
	if err := _Amazon.Visit(url); err != nil {
		_Amazon.err = err
	}

	_Amazon.Wait()

	// 过滤数据
	if _Amazon.err == nil && productInfo.Title != "" && len(productInfo.Price) != 0 && len(productInfo.Images) != 0 && len(productInfo.Style) != 0 {
		_Amazon.onMassage(productInfo)
	}
	_Amazon.err = nil
	return _Amazon
}

type SearchPageStyle struct {
	nextPageUrl string
	detailsUrls []string
	isExist     bool
}

// GetSearchPageStyle 获取查询样式
func (_SearchPageStyle *SearchPageStyle) GetSearchPageStyle(parentElement *colly.HTMLElement) map[string]func(i int, element *colly.HTMLElement) {
	return map[string]func(i int, element *colly.HTMLElement){
		"#a-page > div.a-container.deals-page-outside-desktop > div > div.a-row": _SearchPageStyle.SearchPageStyleOne(parentElement),
		"#a-page > div.a-container.octopus-page-style > div.a-row.apb-browse-two-col-center-pad > div.a-column.a-span12.aok-float-right.apb-browse-col-pad-left.apb-browse-two-col-center-margin-right.a-span-last > div.apb-default-slot":                                                                      _SearchPageStyle.SearchPageStyleTwo(parentElement),
		"#CardInstancetEJUOMnqlASAERH_H76JbA > div.a-cardui._cDEzb_card_1L-Yx > div.p13n-desktop-grid > div.p13n-gridRow._cDEzb_grid-row_3Cywl":                                                                                                                                                                 _SearchPageStyle.SearchPageStyleThree(parentElement),
		"#search > div.s-desktop-width-max.s-desktop-content.s-opposite-dir.s-wide-grid-style.sg-row > div.sg-col-20-of-24.s-matching-dir.sg-col-16-of-20.sg-col.sg-col-8-of-12.sg-col-12-of-16 > div > span.rush-component.s-latency-cf-section > div.s-main-slot.s-result-list.s-search-results.sg-row > div": _SearchPageStyle.SearchPageStyleFour(parentElement),
	}
}

// SearchPageStyleOne 查询页面样式一
func (_SearchPageStyle *SearchPageStyle) SearchPageStyleOne(parentElement *colly.HTMLElement) func(i int, element *colly.HTMLElement) {
	if text := parentElement.ChildAttr("a.s-pagination-item.s-pagination-next.s-pagination-button.s-pagination-separator", "href"); text != "" {
		_SearchPageStyle.nextPageUrl = text
	}
	return func(i int, element *colly.HTMLElement) {
		if href := element.ChildAttr("a.a-link-normal", "href"); href != "" {
			_SearchPageStyle.detailsUrls = append(_SearchPageStyle.detailsUrls, href)
			if _SearchPageStyle.nextPageUrl != "" && len(_SearchPageStyle.detailsUrls) > 3 {
				_SearchPageStyle.isExist = true
			}
		}
	}
}

// SearchPageStyleTwo 查询页面样式二
func (_SearchPageStyle *SearchPageStyle) SearchPageStyleTwo(parentElement *colly.HTMLElement) func(i int, element *colly.HTMLElement) {
	if text := parentElement.ChildAttr("#CardInstancetEJUOMnqlASAERH_H76JbA > div.a-cardui._cDEzb_card_1L-Yx > div.a-text-center > ul > li.a-last > a", "href"); text != "" {
		_SearchPageStyle.nextPageUrl = text
	}
	return func(i int, element *colly.HTMLElement) {
		if href := element.ChildAttr("a.a-link-normal", "href"); href != "" {
			_SearchPageStyle.detailsUrls = append(_SearchPageStyle.detailsUrls, href)
			if _SearchPageStyle.nextPageUrl != "" && len(_SearchPageStyle.detailsUrls) > 3 {
				_SearchPageStyle.isExist = true
			}
		}
	}
}

// SearchPageStyleThree 查询页面样式三
func (_SearchPageStyle *SearchPageStyle) SearchPageStyleThree(parentElement *colly.HTMLElement) func(i int, element *colly.HTMLElement) {
	if text := parentElement.ChildAttr("#CardInstancetEJUOMnqlASAERH_H76JbA > div.a-cardui._cDEzb_card_1L-Yx > div.a-text-center > ul > li.a-last > a", "href"); text != "" {
		_SearchPageStyle.nextPageUrl = text
	}
	return func(i int, element *colly.HTMLElement) {
		if href := element.ChildAttr("div._cDEzb_iveVideoWrapper_JJ34T > div.zg-grid-general-faceout > div.p13n-sc-uncoverable-faceout > a.a-link-normal.aok-block", ""); href != "" {
			_SearchPageStyle.detailsUrls = append(_SearchPageStyle.detailsUrls, href)
			if _SearchPageStyle.nextPageUrl != "" && len(_SearchPageStyle.detailsUrls) > 3 {
				_SearchPageStyle.isExist = true
			}
		}
	}
}

// SearchPageStyleFour 查询页面样式三
func (_SearchPageStyle *SearchPageStyle) SearchPageStyleFour(parentElement *colly.HTMLElement) func(i int, element *colly.HTMLElement) {
	if text := parentElement.ChildAttr("#CardInstancetEJUOMnqlASAERH_H76JbA > div.a-cardui._cDEzb_card_1L-Yx > div.a-text-center > ul > li.a-last > a", "href"); text != "" {
		_SearchPageStyle.nextPageUrl = text
	}
	return func(i int, element *colly.HTMLElement) {
		if href := element.ChildAttr("a.a-link-normal.s-no-outline", "href"); href != "" {
			_SearchPageStyle.detailsUrls = append(_SearchPageStyle.detailsUrls, href)
			if _SearchPageStyle.nextPageUrl != "" && len(_SearchPageStyle.detailsUrls) > 3 {
				_SearchPageStyle.isExist = true
			}
		}
	}
}

// 获取产品url
func (_Amazon *Amazon) searchClassLink(url string) *Amazon {
	searchPageStyle := SearchPageStyle{}
	_Amazon.OnHTML("body", func(e *colly.HTMLElement) {
		// 获取产品链接
		for k, v := range searchPageStyle.GetSearchPageStyle(e) {
			e.ForEach(k, v)
			if searchPageStyle.isExist {
				logs.Logger.Error(logMsg, zap.Bool("break", searchPageStyle.isExist), zap.String("url", url), zap.String("nextPageUrl", searchPageStyle.nextPageUrl), zap.Strings("detailsUrls", searchPageStyle.detailsUrls))
				break
			}
		}
	})

	if err := _Amazon.Visit(url); err != nil {
		logs.Logger.Error(logMsg, zap.Error(err), zap.String("url", url))
		return _Amazon
	}
	_Amazon.Wait()

	// 获取各个产品详情的数据
	logs.Logger.Info(logMsg, zap.String("nextPageUrl", searchPageStyle.nextPageUrl), zap.Strings("detailsUrls", searchPageStyle.detailsUrls))
	for i, detailsUrl := range searchPageStyle.detailsUrls {
		time.Sleep(1 * time.Second)
		_Amazon.ObtainDetails(httpsUrl + detailsUrl)
		// 如果设置了限制则根据限制做，等于零则不进行限制
		if _Amazon.limit != 0 && _Amazon.limit < i {
			break
		}
	}

	if searchPageStyle.nextPageUrl != "" {
		if _Amazon.pageNumber < _Amazon.nextPageLimit {
			logs.Logger.Info(logMsg, zap.Int("pageNumber", _Amazon.pageNumber), zap.String("nextPageUrl", searchPageStyle.nextPageUrl))
			_Amazon.pageNumber++
			_Amazon.searchClassLink(httpsUrl + searchPageStyle.nextPageUrl)
		}
	}
	return _Amazon
}

// ObtainMassage 处理消息方法
func (_Amazon *Amazon) ObtainMassage(onMassage func(attr *ProductAttr)) *Amazon {
	_Amazon.onMassage = onMassage
	return _Amazon
}

// SetLimit 设置限制图片长度
func (_Amazon *Amazon) SetLimit(limit int) *Amazon {
	_Amazon.limit = limit
	return _Amazon
}

// SetNextPageLimit 设置爬取页数
func (_Amazon *Amazon) SetNextPageLimit(nextPageLimit int) *Amazon {
	_Amazon.nextPageLimit = nextPageLimit
	return _Amazon
}

// convertImage 替换最大规格的图片
func (_Amazon *Amazon) convertImage(image string) string {
	index := make([]int, 0)
	for sum, j := 0, len(image)-1; j > 0; j-- {
		if image[j] == uint8('.') {
			index = append(index, j)
			sum++
		}
		if sum == 2 {
			break
		}
	}
	if len(index) > 0 {
		image = image[:index[1]+1] + "_SL1500_" + image[index[0]:]
	}
	return image
}

// stringToNumber 字符串数字过滤
func (_Amazon *Amazon) stringToNumber(s string) float64 {
	// 是否小数点
	isPoint := false
	ss := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			ss += string(s[i])
			isPoint = true
		} else {
			n, err := strconv.Atoi(string(s[i]))
			if isPoint && err != nil {
				break
			}
			if err == nil {
				ss += strconv.Itoa(n)
			}
		}
	}

	nums, _ := strconv.ParseFloat(ss, 64)
	return nums
}

// copyConvertImage 复制图片路径转换
func (_Amazon *Amazon) copyConvertImage(img string) string {
	img = strings.ReplaceAll(img, httpsUrlMedia+"/", "")
	tmpName := []byte(img)
	isPoint := false
	for i := len(tmpName) - 1; i > 0; i-- {
		if tmpName[i] == '+' || tmpName[i] == '-' || tmpName[i] == '/' {
			tmpName[i] = '_'
		}

		if tmpName[i] == '.' && isPoint {
			tmpName = append(tmpName[:i], tmpName[i+1:]...)
		}
		if tmpName[i] == '.' {
			isPoint = true
		}
	}
	return string(tmpName)
}

package dh

import (
	"github.com/gocolly/colly/v2"
	"gotest/practice/okex/utils"
	"strconv"
	"strings"
	"sync"
)

const (
	saveImagePath = "./assets/uploads/crawling/product/" // 存储爬取图片
)

type ProductInfo struct {
	Attrs         map[string][]string //	商品属性
	Images        []map[string]string //	商品图片
	Name          string              //	商品标题
	OriginalMoney float64             //	商品原价
	Money         float64             //	商品价格
	Description   string              //	商品详情
}

type Dhgate struct {
	Debug          bool             //	是否打印日志
	sync           sync.Mutex       //	锁
	collyCollector *colly.Collector //	colly对象
	ProductInfo    *ProductInfo     //	商品信息
}

// NewDhgate 新建敦煌对象
func NewDhgate() *Dhgate {
	collyCollector := colly.NewCollector()
	return &Dhgate{
		ProductInfo: &ProductInfo{
			Attrs:  make(map[string][]string, 0),
			Images: make([]map[string]string, 0),
		},
		collyCollector: collyCollector,
	}
}

func (_Dhgate *Dhgate) Index(url string) (productDetailsUrl []string) {
	_Dhgate.collyCollector.OnHTML("#proGallery > div.content-middle.clearfix > div.gallery-box > div > div.gwrap", func(element *colly.HTMLElement) {
		element.ForEach("div.photo > a.pic", func(i int, element *colly.HTMLElement) {
			domClass := element.Attr("href")
			productDetailsUrl = append(productDetailsUrl, domClass)
		})
	})
	_Dhgate.collyCollector.Visit(url)
	return productDetailsUrl
}

// Details 获取产品详情
func (_Dhgate *Dhgate) Details(url string) *ProductInfo {
	var attrVal []string

	//	获取商品属性
	_Dhgate.collyCollector.OnHTML("#productdisplayForm > div > div.common-sku-item", func(element *colly.HTMLElement) {
		domClass := element.Attr("class")
		switch {
		case strings.Contains(domClass, "key-info-lt"):
			element.Text = strings.TrimRight(element.Text, ":")
			if len(attrVal) > 5 {
				attrVal = attrVal[:5]
			}
			_Dhgate.SetAttrs(element.Text, attrVal)
			attrVal = []string{}
		case strings.Contains(domClass, "key-info-line"):
			switch {
			case strings.Contains(domClass, "color-wrap"):
				element.ForEach("div.pop-note", func(i int, h *colly.HTMLElement) {
					attrVal = append(attrVal, h.Text)
				})
			case strings.Contains(domClass, "size-wrap"):
				element.ForEach("span.price", func(i int, h *colly.HTMLElement) {
					attrVal = append(attrVal, h.Text)
				})
			}
		}
	})

	//	获取商品标题
	_Dhgate.collyCollector.OnHTML("#productdisplayForm > div > div.key-info > div.hinfo.clearfix > h1", func(element *colly.HTMLElement) {
		_Dhgate.ProductInfo.Name = element.Text
	})

	//	获取商品图片 存储商品规格参数, 保存图片组
	_Dhgate.collyCollector.OnHTML("div.bimg > div.j-banner-max.bimg-banner.clearfix", func(element *colly.HTMLElement) {
		element.ForEach("div.bimg-inner", func(i int, element *colly.HTMLElement) {
			imgUrl := element.Attr("data-imgurl")
			if imgUrl != "" {
				element.Request.Visit(imgUrl)
			}
		})
	})
	_Dhgate.collyCollector.OnResponse(func(r *colly.Response) {
		if !utils.PathExists(saveImagePath) {
			utils.PathMkdirAll(saveImagePath)
		}

		if strings.Contains(r.Headers.Get("Content-Type"), "image") {
			// 保存图片
			imgPath := saveImagePath + r.FileName()
			if !utils.PathExists(imgPath) {
				err := r.Save(imgPath)
				if err != nil {
					panic(err)
				}
			}

			_Dhgate.ProductInfo.Images = append(_Dhgate.ProductInfo.Images, map[string]string{"label": "", "value": "/" + imgPath})
		}
	})

	//	获取商品价格
	_Dhgate.collyCollector.OnHTML("div.lineprice > div.wprice-list > ul > li.current", func(element *colly.HTMLElement) {
		element.ForEach("span", func(i int, element *colly.HTMLElement) {
			domClass := element.Attr("class")

			switch {
			case strings.Contains(domClass, "col1"):
				_Dhgate.ProductInfo.Money = _Dhgate.StringToNumber(element.Text)
			case strings.Contains(domClass, "col2"):
				_Dhgate.ProductInfo.OriginalMoney = _Dhgate.StringToNumber(element.Text)
			}
		})
	})

	// 获取商品详情
	_Dhgate.collyCollector.OnHTML("#productdisplayForm > div > div.product-feature.div-wrap", func(element *colly.HTMLElement) {
		_Dhgate.ProductInfo.Description = element.Text
	})
	_Dhgate.collyCollector.Visit(url)

	return _Dhgate.ProductInfo
}

// SetAttrs 设置属性
func (_Dhgate *Dhgate) SetAttrs(attrKey string, attrVal []string) {
	_Dhgate.sync.Lock()
	defer _Dhgate.sync.Unlock()

	if len(attrVal) > 0 {
		if _, ok := _Dhgate.ProductInfo.Attrs[attrKey]; !ok {
			_Dhgate.ProductInfo.Attrs[attrKey] = make([]string, 0)
		}

		_Dhgate.ProductInfo.Attrs[attrKey] = append(_Dhgate.ProductInfo.Attrs[attrKey], attrVal...)
	}
}

// StringToNumber 字符串数字过滤
func (_Dhgate *Dhgate) StringToNumber(s string) float64 {
	ss := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			ss += string(s[i])
		} else {
			n, err := strconv.Atoi(string(s[i]))
			if err == nil {
				ss += strconv.Itoa(n)
			}
		}
	}

	nums, _ := strconv.ParseFloat(ss, 64)
	return nums
}

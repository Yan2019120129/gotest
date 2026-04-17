package amazon

import (
	"gotest/practice/crawling_t/crawling"
	"strconv"
	"strings"
	"sync"
)

type ProductAttr struct {
	mutex sync.Mutex
	crawling.ProductAttr
}

func NewProductAttr() *ProductAttr {
	return &ProductAttr{
		mutex: sync.Mutex{},
		ProductAttr: crawling.ProductAttr{
			Images: make([]string, 0),
			Style:  make(map[string][]string),
		},
	}
}

// SetTitle 设置产品标题
func (_ProductInfo *ProductAttr) SetTitle(title string) *ProductAttr {
	_ProductInfo.Title = strings.TrimSpace(title)
	return _ProductInfo
}

// SetPrice 设置产品金额
func (_ProductInfo *ProductAttr) SetPrice(priceStr string) *ProductAttr {
	// 是否小数点
	isPoint := false
	ss := ""
	for i := 0; i < len(priceStr); i++ {
		if priceStr[i] == '.' {
			ss += string(priceStr[i])
			isPoint = true
		} else {
			n, err := strconv.Atoi(string(priceStr[i]))
			if isPoint && err != nil {
				break
			}
			if err == nil {
				ss += strconv.Itoa(n)
			}
		}
	}

	price, _ := strconv.ParseFloat(ss, 64)

	// 将最小的价格赋予现价，最高的价格赋予原价，避免出现现价比原价高的情况
	if _ProductInfo.CurrentPrice == 0 {
		_ProductInfo.CurrentPrice = price
	} else if _ProductInfo.CurrentPrice > price {
		_ProductInfo.OriginalPrice = _ProductInfo.CurrentPrice
		_ProductInfo.CurrentPrice = price
	} else {
		_ProductInfo.OriginalPrice = price
	}
	return _ProductInfo
}

// SetImages 设置产品图片
func (_ProductInfo *ProductAttr) SetImages(image string) *ProductAttr {
	_ProductInfo.Images = append(_ProductInfo.Images, image)
	return _ProductInfo
}

// SetStyle 获取现价
func (_ProductInfo *ProductAttr) SetStyle(key, value string) *ProductAttr {
	key = strings.Replace(key, ":", "", 1)
	if key == "*" {
		key = "Specification"
	}
	value = strings.TrimSpace(value)
	key = strings.TrimSpace(key)
	_ProductInfo.mutex.Lock()
	defer _ProductInfo.mutex.Unlock()
	if value != "Select" {
		_ProductInfo.Style[key] = append(_ProductInfo.Style[key], value)
	}
	return _ProductInfo
}

// GetStyleLen 获取现价
func (_ProductInfo *ProductAttr) GetStyleLen(key string) int {
	key = strings.Replace(key, ":", "", 1)
	_ProductInfo.mutex.Lock()
	defer _ProductInfo.mutex.Unlock()
	styleLen := len(_ProductInfo.Style[key])
	return styleLen
}

// SetDescribe 设置产品详情
func (_ProductInfo *ProductAttr) SetDescribe(Describe string) *ProductAttr {
	_ProductInfo.Describe = strings.TrimSpace(Describe)
	return _ProductInfo
}

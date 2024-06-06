package main

import (
	"strconv"
	"strings"
	"sync"
)

type ProductAttr struct {
	mutex        sync.Mutex
	Price        []float64           // 金额
	Images       []string            // 图片
	Style        map[string][]string // 样式
	Title        string              // 标题
	Name         string              // 产品名
	Describe     string              // 描述
	DescribeRich []string            // 详细描述
}

func NewProductAttr() *ProductAttr {
	return &ProductAttr{
		mutex:        sync.Mutex{},
		Images:       make([]string, 0),
		Style:        make(map[string][]string),
		DescribeRich: make([]string, 0),
	}
}

// SetDescribe 设置产品详情
func (_ProductInfo *ProductAttr) SetDescribe(describe string) *ProductAttr {
	_ProductInfo.Title = describe
	return _ProductInfo
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
	_ProductInfo.Price = append(_ProductInfo.Price, price)
	return _ProductInfo
}

// SetImages 设置产品图片
func (_ProductInfo *ProductAttr) SetImages(image string) *ProductAttr {
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
	_ProductInfo.Style[key] = append(_ProductInfo.Style[key], value)
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

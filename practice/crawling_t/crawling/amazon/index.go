package amazon

import (
	"gotest/practice/crawling_t/crawling"
	"gotest/practice/okex/utils"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"go.uber.org/zap"
)

const (
	SaveImagePath = "static/crawling/product"
	imagePath     = "/crawling/product"
)

func Headers() map[string]string {
	return map[string]string{"Origin": "https://www.amazon.com", "User-Agent": "Apifox/1.0.0 (https://apifox.com)"}
}

// SaveImage 保存图片
func SaveImage(imageUrl string, saveDir string) {
	// 发起 GET 请求获取图片数据
	response, err := http.Get(imageUrl)
	if err != nil {
		zap.L().Error(crawling.LogMsg, zap.Error(err))
		return
	}
	defer response.Body.Close()

	// 检查 HTTP 响应状态码
	if response.StatusCode != http.StatusOK {
		zap.L().Error(crawling.LogMsg, zap.Int("status", response.StatusCode))
		return
	}

	// 创建指定目录（如果不存在）
	if !utils.PathExists(saveDir) {
		utils.PathMkdirAll(saveDir)
	}

	// 生成本地文件名并创建文件
	localFileName := filepath.Join(saveDir, filepath.Base(ProcessPictures(imageUrl)))
	file, err := os.Create(localFileName)
	if err != nil {
		zap.L().Error(crawling.LogMsg, zap.Error(err))
		return
	}
	defer file.Close()

	// 将响应体中的图片数据复制到本地文件
	_, err = io.Copy(file, response.Body)
	if err != nil {
		zap.L().Error(crawling.LogMsg, zap.Error(err))
		return
	}
}

// ObtainDetails 获取产品信息
func ObtainDetails(settingAdminId uint, selector string) (string, func(e *colly.HTMLElement)) {
	return selector, func(e *colly.HTMLElement) {
		imagePathList := make([]string, 0)
		productInfo := NewProductAttr()
		e.ForEach("#ppd", func(i int, e1 *colly.HTMLElement) {
			// 获取标题
			productInfo.SetTitle(e1.ChildText("#title"))
			if productInfo.Title == "" {
				productInfo.SetTitle(e1.ChildText("#productTitle"))
			}

			// 获取产品金额
			e1.ForEach("#corePriceDisplay_desktop_feature_div > div.aok-align-center ", func(i int, e2 *colly.HTMLElement) {
				priceText := ""
				if priceText = e2.ChildText("span.aok-offscreen"); priceText == "" {
					priceText = e2.ChildText("span.a-offscreen")
				}

				if priceText != "" {
					productInfo.SetPrice(priceText)
				}
			})

			e1.ForEach("#corePrice_desktop > div > table > tbody > tr > td.a-span12 > span ", func(i int, e2 *colly.HTMLElement) {
				e2.ForEach("span.a-offscreen", func(i int, e3 *colly.HTMLElement) {
					if priceText := e3.Text; priceText != "" {
						productInfo.SetPrice(priceText)
					}
				})
			})

			// 获取产品图片
			e1.ForEach("#altImages > ul > li", func(i int, e2 *colly.HTMLElement) {
				imageUrl := e2.ChildAttr("img", "src")
				if !strings.Contains(imageUrl, ".gif") && !strings.Contains(imageUrl, "PKdp-play-icon-overlay") {
					if imageUrl != "" && len(productInfo.Images) < 5 {
						imageUrl = ReplacePictureSize(imageUrl)
						imagePathList = append(imagePathList, imageUrl)

						imgPath := imagePath + "/" + ProcessPictures(GetImageName(imageUrl))
						productInfo.SetImages(imgPath)
					}
				}
			})

			// 获取产品规格  颜色，大小，样式等等
			e1.ForEach("#native_dropdown_selected_size_name > option", func(i int, e2 *colly.HTMLElement) {
				label := e2.ChildText("#native_size_name_-1")
				alt := e2.ChildText("option.dropdownAvailable")
				if label != "" && alt != "" {
					stylLen := productInfo.GetStyleLen(label)
					if stylLen < 5 {
						productInfo.SetStyle(label, alt)
					}
				}
			})

			// 产品属性
			e1.ForEach("#twister > div", func(i int, e2 *colly.HTMLElement) {
				label := e2.ChildText("label.a-form-label")
				e2.ForEach("select > option", func(i int, e3 *colly.HTMLElement) {
					alt := e3.Text
					if label != "" && alt != "" {
						stylLen := productInfo.GetStyleLen(label)
						if stylLen < 5 {
							productInfo.SetStyle(label, alt)
						}
					}
				})

				// 产品属性
				e2.ForEach("ul > li ", func(i int, e3 *colly.HTMLElement) {
					alt := ""
					if alt = e3.ChildAttr("img", "alt"); alt == "" {
						alt = e3.ChildText("button > div > div.twisterTextDiv.text")
					}
					if label != "" && alt != "" {
						stylLen := productInfo.GetStyleLen(label)
						if stylLen < 5 {
							productInfo.SetStyle(label, alt)
						}
					}
				})

				if e2.DOM.Find("select").Length() == 0 && e2.DOM.Find("ul").Length() == 0 {
					alt := e2.ChildText("span.selection")
					productInfo.SetStyle(label, alt)
				}
			})

			// 当没有规格的时候使用品牌作为
			e1.ForEach("#productOverview_feature_div > div > table", func(i int, e2 *colly.HTMLElement) {
				e2.ForEach(" tbody > tr.a-spacing-small.po-brand", func(i int, e3 *colly.HTMLElement) {
					label1 := e3.ChildText("td.a-span3")
					alt := e3.ChildText("td.a-span9")
					productInfo.SetStyle(label1, alt)
				})
			})

			// 产品属性
			e1.ForEach("#poExpander > div.a-expander-content.a-expander-partial-collapse-content > div > table > tbody > tr.a-spacing-small.po-brand", func(i int, e2 *colly.HTMLElement) {
				label := e2.ChildText("td > span.a-size-base.a-text-bold")
				alt := e2.ChildText("td > span.a-size-base.po-break-word")
				productInfo.SetStyle(label, alt)
			})

			// 产品详情样式1
			e1.ForEach("#productFactsDesktop_feature_div", func(i int, e2 *colly.HTMLElement) {
				e2.DOM.Find("li span.a-list-item.a-size-base.a-color-base").Each(func(i int, s *goquery.Selection) {
					spanText := s.Text()
					productInfo.SetDescribe(spanText)
				})
			})

			// 产品详情样式2
			if productInfo.Describe == "" {
				e1.ForEach("#feature-bullets > ul", func(i int, e2 *colly.HTMLElement) {
					productInfo.SetDescribe(e2.Text)
				})
			}

			// 产品详情
			var spanText string
			e1.ForEach("#feature-bullets > ul", func(i int, e2 *colly.HTMLElement) {
				e2.DOM.Find("li span.a-list-item").Each(func(i int, s *goquery.Selection) {
					descText := strings.TrimSpace(s.Text())
					spanText += descText
				})
				if productInfo.Describe == "" {
					productInfo.Describe = spanText
				}
			})

			// 产品详情样式3
			if productInfo.Describe == "" {
				e1.ForEach("#detailBulletsWrapper_feature_div", func(i int, e2 *colly.HTMLElement) {
					productInfo.SetDescribe(e2.Text)
				})
			}
		})

		// 过滤不合格的数据
		if productInfo.Title == "" || productInfo.Describe == "" || len(productInfo.Images) == 0 || len(productInfo.Style) == 0 {
			return
		}

		// 如果该管理员存在对应的批发产品不进行返回
		//numb := model.NewModel().Where("admin_id = ?", settingAdminId).Where("name = ?", productInfo.Title).Where("money = ?", productInfo.CurrentPrice).Where("type = ?", models.ProductTypeWholesale).Find(&models.Product{}).RowsAffected
		//if numb > 0 {
		//	return
		//}

		e.Response.Ctx.Put("message", productInfo.ProductAttr)

		// 获取产品图片
		for _, httpUrl := range imagePathList {
			go SaveImage(httpUrl, SaveImagePath)
		}
	}
}

// SearchClassLink 获取产品url
func SearchClassLink(selector string) (string, func(e *colly.HTMLElement)) {
	return selector, func(element *colly.HTMLElement) {
		hrefs := element.ChildAttrs("a.a-link-normal.s-no-outline", "href")
		for _, href := range hrefs {
			err := element.Request.Visit(href)
			if err != nil {
				zap.L().Error(crawling.LogMsg, zap.Error(err))
				return
			}
		}
	}
}

func ReplacePictureSize(imageName string) string {
	index := make([]int, 0)
	for sum, j := 0, len(imageName)-1; j > 0; j-- {
		if imageName[j] == uint8('.') {
			index = append(index, j)
			sum++
		}
		if sum == 2 {
			break
		}
	}
	if len(index) > 0 {
		imageName = imageName[:index[1]+1] + "_SL1500_" + imageName[index[0]:]
	}
	return imageName
}

// GetImageName 获取文件名
func GetImageName(rawUrl string) string {
	parse, err := url.Parse(rawUrl)
	if err != nil {
		return ""
	}
	index := strings.LastIndex(parse.Path, "/")
	return parse.Path[index+1:]
}

// ProcessPictures 处理图片
func ProcessPictures(image string) string {
	return strings.ReplaceAll(image, "._SL1500_", "_SL1500_")
}

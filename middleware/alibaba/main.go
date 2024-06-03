// This file is auto-generated, don't edit it. Thanks.
package main

import (
	"encoding/json"
	"fmt"
	alimt20181012 "github.com/alibabacloud-go/alimt-20181012/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"go.uber.org/zap"
	"gotest/common/models"
	"gotest/common/module/gorm/database"
	"gotest/common/module/logs"
	"strings"
)

const (
	logMsg = "alibaba"
)

var (
	path            = "mt.cn-hangzhou.aliyuncs.com/"
	AccessKeyID     = "LTAI5tHg2m4pJP2EkoqZHSTX"
	AccessKeySecret = "Vj24DStm0oH5kRj0JjY1fsBft3sBnG"
	signVersion     = "1.0"
	signMethod      = "HMAC-SHA1"
)

type Massage struct {
	RequestId string `json:"RequestId"`
	Message   string `json:"Message"`
	Recommend string `json:"Recommend"`
	HostId    string `json:"HostId"`
	Code      string `json:"Code"`
}

type tmpProduct struct {
	Id       uint
	ParentId uint
	Title    string
	Desc     string
}

func main() {
	productList := make([]models.Product, 0)
	database.DB.Find(&productList)
	instance, err := NewLanguage(AccessKeyID, AccessKeySecret)
	if err != nil {
		return
	}

	value := "MEROKEETY Women's Boho Leopard Print Skirt Pleated A-Line Swing Midi Skirts"
	formatType := "text"
	scene := "title"
	sourceLanguage := "es"
	targetLanguage := "en"

	//for i, product := range productList {
	title := &alimt20181012.TranslateGeneralResponse{}
	val := &alimt20181012.TranslateGeneralRequest{
		FormatType:     &formatType,
		Scene:          &scene,
		SourceLanguage: &sourceLanguage,
		SourceText:     &value,
		TargetLanguage: &targetLanguage,
	}
	if err := instance.FindLanguage(val, title); err != nil {
		logs.Logger.Error(logMsg, zap.Error(err))
		panic(err)
	}
	//logs.Logger.Info(logMsg, zap.Reflect("val", val))
	//time.Sleep(500 * time.Millisecond)

	//desc := &alimt20181012.TranslateGeneralResponse{}
	//val.SourceText = &product.Desc
	//if err := instance.FindLanguage(val, desc); err != nil {
	//	logs.Logger.Error(logMsg, zap.Error(err))
	//	panic(err)
	//}

	logs.Logger.Info(logMsg, zap.Reflect("val", val))
	logs.Logger.Info(logMsg, zap.Reflect("title", title))
	//logs.Logger.Info(logMsg, zap.Reflect("desc", desc))

	//if desc != nil && title != nil {
	//	database.DB.Create(&tmpProduct{
	//		ParentId: product.ID,
	//		Title:    *title.Body.Data.Translated,
	//		Desc:     *desc.Body.Data.Translated,
	//	})
	//}
	//time.Sleep(500 * time.Millisecond)
	//if i == 1 {
	//	break
	//}
	//}
}

type Language struct {
	*alimt20181012.Client
}

func NewLanguage(accessKeyID, accessKeySecret string) (*Language, error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyID),
		AccessKeySecret: tea.String(accessKeySecret),
	}
	config.Endpoint = tea.String(path)
	language := &Language{}
	language.Client = &alimt20181012.Client{}
	var err error
	language.Client, err = alimt20181012.NewClient(config)
	return language, err
}

// FindLanguage 查找语言对应语种
func (l *Language) FindLanguage(val *alimt20181012.TranslateGeneralRequest, resp *alimt20181012.TranslateGeneralResponse) (_err error) {
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				logs.Logger.Error(logMsg, zap.Error(r))
				_e = r
			}
		}()

		// 复制代码运行请自行打印 API 的返回值
		resp, _err = l.TranslateGeneralWithOptions(val, runtime)
		if _err != nil {
			logs.Logger.Error(logMsg, zap.Error(_err))
			return _err
		}
		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}

		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		// 错误 message
		fmt.Println(tea.StringValue(error.Message))
		// 诊断地址
		var data interface{}
		d := json.NewDecoder(strings.NewReader(tea.StringValue(error.Data)))
		d.Decode(&data)
		if m, ok := data.(map[string]interface{}); ok {
			recommend, _ := m["Recommend"]
			fmt.Println(recommend)
		}
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			logs.Logger.Error(logMsg, zap.Error(_err))
			return _err
		}
	}
	return _err
}

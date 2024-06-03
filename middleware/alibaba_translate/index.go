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

func main() {
	productList := make([]models.Product, 0)
	database.DB.Find(&productList)
	instance, err := NewLanguage(AccessKeyID, AccessKeySecret)
	if err != nil {
		return
	}
	formatType := "text"
	scene := "title"
	sourceLanguage := "es"
	targetLanguage := ""
	for _, product := range productList {
		err = instance.FindLanguage(&alimt20181012.TranslateGeneralRequest{
			FormatType:     &formatType,
			Scene:          &scene,
			SourceLanguage: &sourceLanguage,
			SourceText:     &product.Name,
			TargetLanguage: &targetLanguage,
		})
		if err != nil {
			panic(err)
		}
	}
}

type Language struct {
	*alimt20181012.Client
	alimt20181012.TranslateGeneralRequest
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
func (l *Language) FindLanguage(val *alimt20181012.TranslateGeneralRequest) (_err error) {
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()

		// 复制代码运行请自行打印 API 的返回值
		value, _err := l.TranslateGeneralWithOptions(val, runtime)
		if _err != nil {
			return _err
		}
		logs.Logger.Info(logMsg, zap.Reflect("body", value))

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
			return _err
		}
	}
	return _err
}

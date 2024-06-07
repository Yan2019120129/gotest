package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"net/http"
	"strings"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	acceptLanguage := r.Header.Get("Accept-Language")
	arr := strings.Split(acceptLanguage, ",")
	locale := arr[0]
	val := mux.Vars(r)
	//locale := val["locale"]
	name := val["name"]
	fmt.Println(language.Korean.String(), language.English.String(), language.Chinese.String())

	var lang language.Tag

	// 使用 language.MustParse()为区域设置分配适当的语言标签。
	switch locale {
	case "ko-KR":
		lang = language.MustParse(language.Korean.String())
	case "en-US":
		lang = language.MustParse(language.English.String())
	case "zh-CN":
		lang = language.MustParse(language.Chinese.String())
	default:
		http.NotFound(w, r)
		return
	}

	// 使用对应语言初始化一个message.Printer实例
	p := message.NewPrinter(lang)

	p.Fprintf(w, "Welcome!\n")

	// 将欢迎信息翻译成目标语言。
	if _, err := p.Fprintf(w, "Hello %s!\n", name); err != nil {
		fmt.Println(err)
		return
	}
}

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	val := mux.Vars(r)
	locale := val["locale"]

	var lang language.Tag

	// 使用 language.MustParse()为区域设置分配适当的语言标签。
	switch locale {
	case "en-gb":
		lang = language.MustParse("en-GB")
	case "de-de":
		lang = language.MustParse("de-DE")
	case "zh-cn":
		lang = language.MustParse("zh-CN")
	default:
		http.NotFound(w, r)
		return
	}

	// 使用对应语言初始化一个message.Printer实例
	p := message.NewPrinter(lang)
	// 将欢迎信息翻译成目标语言。
	if _, err := p.Fprintf(w, "Welcome!\n"); err != nil {
		fmt.Println(err)
		return
	}
}

package language_t

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"testing"
)

// TestLanguage 测试基础语言包的作用
func TestLanguage(t *testing.T) {
	err := message.SetString(language.Chinese, "%s is %d years old", "%s 今年 %d 岁了")
	if err != nil {
		return
	}
	msg := message.Render("%s is %d years old", "张三")
	fmt.Println(msg)
}

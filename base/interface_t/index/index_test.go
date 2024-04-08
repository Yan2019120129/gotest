package index

import (
	"bytes"
	"fmt"
	"gotest/common/models"
	"reflect"
	"testing"
	"unicode"
)

// TestInterfaceToStruct 测试接口转换为结构体
func TestInterfaceToStruct(t *testing.T) {
	InterfaceToStruct()
}

// TestInterfaceToStruct 测试接口转换为结构体
func TestIsInterfaceOrBaseType(t *testing.T) {

	data := []interface{}{"yan", &models.User{}, models.AdminUser{}}
	for _, v := range data {
		var valueType = reflect.TypeOf(v)
		//if valueType.Kind() == reflect.Ptr {
		//	valueType = valueType.Elem()
		//}
		switch valueType.Kind() {
		case reflect.String:
			fmt.Println("字符串", valueType.Name())
		case reflect.Struct:
			name := CamelToSnake(valueType.Name())
			fmt.Println("结构体", name)
		case reflect.Ptr:
			name := CamelToSnake(valueType.Elem().Name())
			fmt.Println("结构体取地址", name)
		default:
			fmt.Println("请传入模型结构体或表名")
		}
	}
}

// CamelToSnake 将驼峰命名转换为下划线命名
func CamelToSnake(s string) string {
	var buf bytes.Buffer
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				buf.WriteRune('_')
			}
			buf.WriteRune(unicode.ToLower(r))
		} else {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

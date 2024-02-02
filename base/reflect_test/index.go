package reflect_test

import (
	"fmt"
	"gotest/common/models"
	"reflect"
)

// ReflectModel 使用反射创建结构体
func ReflectModel() {
	var userReflect interface{}
	userReflect = &models.User{}
	fmt.Printf("%T\n", userReflect)
	structType := reflect.TypeOf(userReflect)
	structInstance := reflect.TypeOf(structType)
	structInstance.Elem()
}

package reflect_t

import (
	"fmt"
	"gotest/common/models"
	"gotest/common/utils"
	"reflect"
	"testing"
)

// TestReflectModel 使用反射创建结构体
func TestReflectModel(t *testing.T) {
	model := []*models.User{models.GetDefaultUser(), models.GetDefaultUser()}
	modelsType := reflect.TypeOf(model)
	if modelsType.Kind() == reflect.Ptr {
		modelsType = modelsType.Elem()
	}
	fmt.Println("modelsType", modelsType.Name())
	for i := modelsType.NumField() - 1; 0 < i; i-- {
		fmt.Println("models field", modelsType.Field(i).Name)
	}
}

// TestReflectFieldValue
func TestReflectFieldValue(t *testing.T) {
	userInfo := models.GetDefaultUser()
	// 使用反射获取结构体字段的值
	value := reflect.ValueOf(*userInfo)
	field := value.FieldByName("UserName")
	fmt.Println("Name:", field)
	if field.IsValid() {
		fmt.Println("Name:", field.Interface())
	} else {
		fmt.Println("Field not found")
	}
}

// TestUtilsReflect 测试反射工具
func TestUtilsReflect(t *testing.T) {
	//var data interface{}
	//data = []*models.User{models.GetDefaultUser(), models.GetDefaultUser()}
	//data = []*models.AdminUser{models.GetDefaultAdminUser(), models.GetDefaultAdminUser()}
	userInfo := models.GetDefaultUser()
	instance := utils.NewReflectModel(userInfo)
	userName := instance.GetValue("UserName")
	money := instance.GetValue("Money")
	fmt.Println(userName)
	fmt.Println(money.(float64))
}

// TestUtilsReflects 测试反射工具slice切片处理
func TestUtilsReflects(t *testing.T) {
	data := []*models.ProductCategory{models.GetProductCategoryDefault(), models.GetProductCategoryDefault()}
	instance := utils.NewReflectModel(data)
	for _, v := range instance.GetClient() {
		fmt.Println("modelValue:", v.GetModelValue())
		fmt.Println("name:", v.GetName())
		fmt.Println("fields:", v.GetFields())
		fmt.Println("ParentId:", v.GetValue("ParentId"))
	}
}

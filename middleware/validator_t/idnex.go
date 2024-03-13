package validator_t

import (
	"fmt"
	"gotest/common/module/validator"
)

// Validator 用于测试validator 验证器
type Validator struct {
	Name string `validate:"required"`
	Age  int    `validate:"required"`
	Sex  int    `validate:"required"`
}

// ValidatorParam 验证参数
func ValidatorParam() {
	param := &Validator{
		Name: "Yan",
		Age:  18,
	}
	if err := validator.Validator(param); err != nil {
		fmt.Println("err:", err)
	}
}

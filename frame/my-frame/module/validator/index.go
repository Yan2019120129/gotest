package validator

import (
	"github.com/go-playground/validator/v10"
)

var instant = validator.New()

func Validator(data interface{}) error {
	if errs := instant.Struct(data); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			return err
		}
	}
	return nil
}

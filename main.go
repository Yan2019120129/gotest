package main

import (
	"fmt"
	"gotest/common/models"
	"gotest/common/utils"
)

func main() {
	userInfo := models.User{}
	reflectInstance := utils.NewReflectModel(userInfo)
	modelsName := reflectInstance.GetName()
	modelsName = utils.CamelToSnake(modelsName)
	fields := reflectInstance.GetFields()
	for i, field := range fields {
		fields[i] = utils.CamelToSnake(field)
	}
	tag := reflectInstance.GetFieldsDesc("status", "gorm", "comment:")
	fieldInfo := reflectInstance.GetFieldInfo("STATUS")
	fieldValue := reflectInstance.GetValue("Status")
	fmt.Println(modelsName, fields, tag, fieldInfo, fieldValue)
}

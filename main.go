package main

import (
	"fmt"
	"gotest/common/models"
	"gotest/common/module/gorm/database"
	"os"
	"strings"
)

const (
	sqlPath      = "/Users/taozi/Documents/Golang/gotest/common/file/product"
	imgPath      = "/public/product"
	removePath   = "/Users/taozi/Documents/Golang/gotest/common/file/img"
	originalPath = "/Users/taozi/Documents/Golang/gotest/common/file"
)

func main() {
	InitProductInfo()
}

// InitProductInfo 插入产品信息
func InitProductInfo() {
	dir, err := os.ReadDir(sqlPath)
	if err != nil {
		fmt.Println("read file path err", sqlPath)
	}

	tx := database.DB.Begin()
	for _, file := range dir {
		if !file.IsDir() {
			fileName := file.Name()
			if strings.LastIndex(fileName, ".sql") > -1 {
				allPath := sqlPath + "/" + fileName
				value, err := os.ReadFile(allPath)
				if err != nil {
					return
				}
				valueArr := strings.Split(string(value), ";\n")
				for _, v := range valueArr {
					if v != "" {
						if err = tx.Exec(v).Error; err != nil {
							tx.Rollback()
							return
						}
					}
				}
			}
		}
	}
	tx.Commit()
}

// RemoveFile 移动不存在的文件
func RemoveFile() {
	dir, err := os.ReadDir(originalPath)
	if err != nil {
		return
	}

	productList := make([]*models.Product, 0)
	database.DB.Find(&productList)
	dbImgMap := make(map[string]string)
	for _, product := range productList {
		for _, image := range product.Images {
			dbImgMap[image] = image
		}
	}

	for _, file := range dir {
		if !file.IsDir() {
			fileName := file.Name()
			if strings.LastIndex(fileName, ".json") > -1 {
				allPath := imgPath + "/" + fileName
				if _, ok := dbImgMap[allPath]; !ok {
					err = os.Rename(allPath, removePath+"/"+fileName)
					if err != nil {
						return
					}
				}
			}
		}
	}
}

package mysql

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gotest/my_frame/config"
	"log"
	"os"
)

type Mysql struct{}

func (m *Mysql) Connect() (db *gorm.DB) {
	db, err := gorm.Open(mysql.Open(m.GetDsn()), &gorm.Config{})
	if err != nil {
		log.Println("连接Mysql出错！！！")
		panic(err)
	}
	return db
}

func (m *Mysql) GetDsn() string {
	fileData, err := os.ReadFile(config.FilePath)
	if err != nil {
		log.Println("读取config.yaml文件错误！！！")
		panic(err)
	}

	dsn := new(config.Dsn)
	err = yaml.Unmarshal(fileData, &dsn)
	if err != nil {
		log.Println("解析config.yaml问价出错！！！")
		panic(err)
	}

	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", dsn.User, dsn.Pass, dsn.Host, dsn.Port, dsn.DbnName)
}

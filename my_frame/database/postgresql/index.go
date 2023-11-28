package postgresql

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gotest/my_frame/config"
	"log"
	"os"
)

type Postgresql struct{}

func (m *Postgresql) Connect() (db *gorm.DB) {
	db, err := gorm.Open(mysql.Open(m.GetDsn()), &gorm.Config{})
	if err != nil {
		log.Println("连接Mysql出错！！！")
		panic(err)
	}
	return db
}

func (m *Postgresql) GetDsn() string {
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
	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai", dsn.Host, dsn.User, dsn.Pass, dsn.DbnName, dsn.Port)
}

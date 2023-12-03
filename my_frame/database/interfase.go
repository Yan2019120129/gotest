package database

import (
	"gorm.io/gorm"
	"gotest/my_frame/config"
)

type Database interface {
	GetDsn() string
	Connect() *gorm.DB
}

type ObstructFactory interface {
	NewMysql() Database
	NewPostgresql() Database
}

type New struct {
}

func (c *New) NewMysql() Database {
	return &config.Mysql{}
}

func (c *New) NewPostgresql() Database {
	return &config.Postgresql{}
}

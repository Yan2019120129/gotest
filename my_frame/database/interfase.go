package database

import (
	"gorm.io/gorm"
	"gotest/my_frame/init_config"
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
	return &init_config.Mysql{}
}

func (c *New) NewPostgresql() Database {
	return &init_config.Postgresql{}
}

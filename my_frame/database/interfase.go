package database

import (
	"gorm.io/gorm"
	"gotest/my_frame/config/mysql"
	"gotest/my_frame/config/postgresql"
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
	return &redis.Mysql{}
}

func (c *New) NewPostgresql() Database {
	return &postgresql.Postgresql{}
}

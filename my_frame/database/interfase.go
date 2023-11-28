package database

import (
	"gorm.io/gorm"
	"gotest/my_frame/database/mysql"
	"gotest/my_frame/database/postgresql"
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
	return &mysql.Mysql{}
}

func (c *New) NewPostgresql() Database {
	return &postgresql.Postgresql{}
}

package init_config

import "gorm.io/gorm"

var Db *gorm.DB

func InitDatabase() {
	InitConfig()
	switch Cfg.Database.UseDatabase {
	// 初始化Postgresql数据库
	case DatabaseTypePostgresql:
		postgresql := new(Postgresql)
		Db = postgresql.Connect()
	// 初始化mysql数据库
	case DatabaseTypeMysql:
		mysql := new(Mysql)
		Db = mysql.Connect()
	}
	InitRedis()
}

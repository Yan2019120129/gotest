package database

type AbstractFactory interface {
	CreateDatabase() Database
}

// MysqlFactory 具体工厂类型
type MysqlFactory struct{}

func (f *MysqlFactory) CreateDatabase() Database {
	return &Mysql{}
}

// OracleFactory 具体工厂类型
type OracleFactory struct{}

func (f *OracleFactory) CreateDatabase() Database {
	return &Oracle{}
}

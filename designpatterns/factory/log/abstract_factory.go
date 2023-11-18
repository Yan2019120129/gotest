package log

type AbstractFactory interface {
	CreateLog4j() Log
	CreateLogger() Log
}

package config

const (
	FilePath = "./my_frame/config/config.yaml"
)

type Database struct {
	MysqlDsn      *Dsn `json:"mysql"`
	PostgresqlDsn *Dsn `json:"postgresql"`
}

type Dsn struct {
	Host    int64  `json:"host"`
	DbnName string `json:"dbn_name"`
	User    string `json:"user"`
	Pass    string `json:"pass"`
	Port    int64  `json:"port"`
}

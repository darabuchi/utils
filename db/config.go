package db

type Database string

const (
	MySql      Database = "mysql"
	Postgres   Database = "postgres"
	Sqlite     Database = "sqlite"
	Sqlserver  Database = "sqlserver"
	Clickhouse Database = "clickhouse"
)

type Config struct {
	Dsn      string   `json:"dsn,omitempty"`
	Database Database `json:"database,omitempty"`
	Debug    bool     `json:"debug,omitempty"`

	DriverName string `json:"driver,omitempty"`
}

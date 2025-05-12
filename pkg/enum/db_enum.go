package enum

type Database string

const (
	PostgresDB Database = "postgresql"
	Redis      Database = "redis"
)

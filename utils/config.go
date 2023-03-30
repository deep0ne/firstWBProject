package utils

type Config struct {
	DBDriver      string
	DBSource      string
	ServerAddress string
}

func NewConfig() Config {
	return Config{
		DBDriver:      "pgx",
		DBSource:      "postgresql://root:wbpass@localhost:5432/wborders?sslmode=disable",
		ServerAddress: "0.0.0.0:8081",
	}
}

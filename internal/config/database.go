package config

import "github.com/freecsdon30/url-shortener/internal/utils"

type DBConfig struct {
	username string
	password string
	host     string
	port     string
}

func NewDBConfig() DBConfig {
	return DBConfig{
		username: utils.GetEnv("DB_USERNAME", "root"),
		password: utils.GetEnv("DB_PASSWORD", "root"),
		host:     utils.GetEnv("DB_HOST", "localhost"),
		port:     utils.GetEnv("DB_PORT", "5432"),
	}
}

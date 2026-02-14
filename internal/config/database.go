package config

import "github.com/freecsdon30/url-shortener/internal/utils"

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func NewDBConfig() DBConfig {
	return DBConfig{
		Username: utils.GetEnv("DB_USERNAME", "root"),
		Password: utils.GetEnv("DB_PASSWORD", "root"),
		Host:     utils.GetEnv("DB_HOST", "localhost"),
		Port:     utils.GetEnv("DB_PORT", "5432"),
		Name:     utils.GetEnv("DB_NAME", "shortener"),
	}
}

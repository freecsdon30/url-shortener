package config

var DB DBConfig

func LoadAllConfig() {
	DB = NewDBConfig()
}

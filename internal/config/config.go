package config

var db DBConfig

func LoadAllConfig() {
	db = NewDBConfig()
}

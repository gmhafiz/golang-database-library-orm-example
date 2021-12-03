package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	DB Database
}

type Database struct {
	Host     string `default:"localhost"`
	Port     int    `default:"5432"`
	User     string `default:"user"`
	Name     string `default:"postgres"`
	Password string `default:"password"`
	SSLMode  string `default:"disable"`
}

func New() *Configuration {
	cfg := &Configuration{
		DB: DB(),
	}

	return cfg
}

func DB() Database {
	var db Database
	envconfig.MustProcess("DB", &db)

	return db
}

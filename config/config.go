package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	DB Database
}

type Database struct {
	Type     string `default:"postgres"`
	Host     string `default:"localhost"`
	Port     int    `default:"5432"`
	Name     string `default:"db_test"`
	User     string `default:"user"`
	Password string `default:"password"`
	SSLMode  string `default:"disable"`
}

func New() *Configuration {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

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

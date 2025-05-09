package db

import "github.com/erlendromo/forsete-atr/src/util"

type DBConfig struct {
	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PASS string
	DB_NAME string
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		DB_HOST: util.MustGetEnv("DB_HOST"),
		DB_PORT: util.MustGetEnv("DB_PORT"),
		DB_USER: util.MustGetEnv("DB_USER"),
		DB_PASS: util.MustGetEnv("DB_PASS"),
		DB_NAME: util.MustGetEnv("DB_NAME"),
	}
}

package config

import (
	"github.com/erlendromo/forsete-atr/src/config/api"
	"github.com/erlendromo/forsete-atr/src/config/db"
)

var config *Config

type Config struct {
	apiConfig *api.APIConfig
	dbConfig  *db.DBConfig
}

func GetConfig() *Config {
	if config == nil {
		config = &Config{
			apiConfig: api.NewAPIConfig(),
			dbConfig:  db.NewDBConfig(),
		}
	}

	return config
}

func (c *Config) APIConfig() *api.APIConfig {
	return c.apiConfig
}

func (c *Config) DBConfig() *db.DBConfig {
	return c.dbConfig
}

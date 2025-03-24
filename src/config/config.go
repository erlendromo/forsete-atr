package config

import (
	"fmt"
	"os"
)

type Config struct {
	API_PORT string
	DEVICE   string
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		config = &Config{
			API_PORT: mustGetEnv("API_PORT"),
			DEVICE:   mustGetEnv("DEVICE"),
		}
	}

	return config
}

func mustGetEnv(key string) string {
	v, found := os.LookupEnv(key)
	if !found {
		panic(fmt.Errorf("Environment variable '%s' not set...", key))
	}

	return v
}

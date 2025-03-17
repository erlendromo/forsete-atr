package config

import (
	"fmt"
	"os"
)

type Config struct {
	API_PORT string
}

func NewConfig() *Config {
	return &Config{
		API_PORT: mustGetEnv("API_PORT"),
	}
}

func mustGetEnv(key string) string {
	v, found := os.LookupEnv(key)
	if !found {
		panic(fmt.Errorf("Environment variable '%s' not set...", key))
	}

	return v
}

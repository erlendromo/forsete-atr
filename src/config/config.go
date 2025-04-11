package config

import (
	"fmt"
	"os"
	"time"

	"github.com/erlendromo/forsete-atr/src/util"
)

type Config struct {
	API_PORT string
	DEVICE   string
	TIMEOUT  time.Duration
}

var config *Config

func GetConfig() *Config {
	timeout, err := time.ParseDuration(mustGetEnv(util.TIMEOUT))
	if err != nil {
		panic(err)
	}

	if config == nil {
		config = &Config{
			API_PORT: mustGetEnv(util.API_PORT),
			DEVICE:   mustGetEnv(util.DEVICE),
			TIMEOUT:  timeout,
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

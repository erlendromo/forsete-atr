package api

import (
	"time"

	"github.com/erlendromo/forsete-atr/src/util"
)

type APIConfig struct {
	API_PORT string
	DEVICE   string
	TIMEOUT  time.Duration
}

func NewAPIConfig() *APIConfig {
	timeout, err := time.ParseDuration(util.MustGetEnv("TIMEOUT"))
	if err != nil {
		panic(err)
	}

	return &APIConfig{
		API_PORT: util.MustGetEnv("API_PORT"),
		DEVICE:   util.MustGetEnv("DEVICE"),
		TIMEOUT:  timeout,
	}
}

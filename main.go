package main

import (
	"os"

	"github.com/erlendromo/forsete-atr/src/cmd/rest"
	"github.com/erlendromo/forsete-atr/src/util"
)

func init() {
	if _, found := os.LookupEnv("API_PORT"); !found {
		os.Setenv("API_PORT", util.DEFAULT_API_PORT)
	}
}

func main() {
	rest.StartRestService()
}

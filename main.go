package main

import (
	"os"

	"github.com/erlendromo/forsete-atr/src/cmd/rest"
	"github.com/erlendromo/forsete-atr/src/domain/model"
	"github.com/erlendromo/forsete-atr/src/util"
)

func init() {
	if _, found := os.LookupEnv(util.API_PORT); !found {
		os.Setenv(util.API_PORT, util.DEFAULT_API_PORT)
	}

	if err := model.InitModels(); err != nil {
		panic(err)
	}
}

func main() {
	rest.StartRestService()
}

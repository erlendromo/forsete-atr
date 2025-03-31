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

	if _, found := os.LookupEnv(util.DEVICE); !found {
		os.Setenv(util.DEVICE, util.DEFAULT_DEVICE)
	}

	if _, found := os.LookupEnv(util.TIMEOUT); !found {
		os.Setenv(util.TIMEOUT, util.DEFAULT_TIMEOUT)
	}

	if err := model.InitModels(); err != nil {
		panic(err)
	}
}

func main() {
	rest.StartRestService()
}

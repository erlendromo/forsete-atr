package rest

import (
	"github.com/erlendromo/forsete-atr/src/api/router/httprouter"
	"github.com/erlendromo/forsete-atr/src/config"
)

func StartRestService() {
	config := config.NewConfig()
	router := httprouter.NewHTTPRouter(config.API_PORT)

	if err := router.Serve(); err != nil {
		panic(err)
	}
}

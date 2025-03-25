package rest

import (
	"github.com/erlendromo/forsete-atr/src/api/router/httprouter"
	"github.com/erlendromo/forsete-atr/src/config"
	"github.com/erlendromo/forsete-atr/src/util"
)

func StartRestService() {
	util.StartTimer()
	config := config.GetConfig()
	router := httprouter.NewHTTPRouter(config.API_PORT)

	if err := router.Serve(); err != nil {
		panic(err)
	}
}

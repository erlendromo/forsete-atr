package cmd

import (
	"github.com/erlendromo/forsete-atr/src/api/router/httprouter"
	"github.com/erlendromo/forsete-atr/src/config"
	"github.com/erlendromo/forsete-atr/src/util"
)

func StartService() {
	util.StartTimer()
	apiConfig := config.GetConfig().APIConfig()
	router := httprouter.NewHTTPRouter(apiConfig.API_PORT)

	if err := router.Serve(); err != nil {
		panic(err)
	}
}

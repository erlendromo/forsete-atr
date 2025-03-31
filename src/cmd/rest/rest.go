package rest

import (
	"fmt"

	"github.com/erlendromo/forsete-atr/src/api/router/httprouter"
	"github.com/erlendromo/forsete-atr/src/config"
	"github.com/erlendromo/forsete-atr/src/util"
)

func StartRestService() {
	util.StartTimer()
	config := config.GetConfig()
	router := httprouter.NewHTTPRouter(config.API_PORT)

	fmt.Printf(
		"\n%sEnvironment%s\nAPI_PORT: %s\nDEVICE: %s\nTIMEOUT: %s\n\n",
		util.PURPLE,
		util.RESET,
		config.API_PORT,
		config.DEVICE,
		config.TIMEOUT,
	)

	if err := router.Serve(); err != nil {
		panic(err)
	}
}

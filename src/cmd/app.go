package cmd

import (
	"context"
	"fmt"

	appcontext "github.com/erlendromo/forsete-atr/src/api/app_context"
	"github.com/erlendromo/forsete-atr/src/api/router/httprouter"
	"github.com/erlendromo/forsete-atr/src/config"
	"github.com/erlendromo/forsete-atr/src/database/postgresql"
	"github.com/erlendromo/forsete-atr/src/util"
)

func StartService() {
	util.StartUTCTimer()
	apiConfig := config.GetConfig().APIConfig()
	router := httprouter.NewHTTPRouter(apiConfig.API_PORT)

	db := postgresql.NewPostgreSQLDatabase()
	appcontext.InitAppContext(db.Database())

	// Setup pipelines on launch
	if _, err := appcontext.GetAppContext().ATRService.CreatePipelines(context.Background()); err != nil {
		panic(fmt.Sprintf("unable to initialize pipelines: %s", err.Error()))
	}

	// Can migrate down on database when service stops,
	// but this will remove all entries in the db,
	// so it will be commented out for now and removed later.

	// defer db.MigrateDown()

	if err := router.Serve(); err != nil {
		panic(err)
	}
}

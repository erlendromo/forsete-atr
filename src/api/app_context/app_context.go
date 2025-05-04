package appcontext

import (
	"context"
	"fmt"

	atrservice "github.com/erlendromo/forsete-atr/src/business/usecase/service/atr_service"
	authservice "github.com/erlendromo/forsete-atr/src/business/usecase/service/auth_service"
	fileservice "github.com/erlendromo/forsete-atr/src/business/usecase/service/file_service"
	"github.com/jmoiron/sqlx"
)

var appCtx *AppContext

type AppContext struct {
	AuthService *authservice.AuthService
	FileService *fileservice.FileService
	ATRService  *atrservice.ATRService

	db *sqlx.DB
	//cache *cache.Cache
}

func InitAppContext(db *sqlx.DB) {
	if appCtx == nil {
		appCtx = &AppContext{
			AuthService: authservice.NewAuthService(db),
			FileService: fileservice.NewFileService(db),
			ATRService:  atrservice.NewATRService(db),
			db:          db,
		}
	}

	if _, err := appCtx.ATRService.CreatePipelines(context.Background()); err != nil {
		panic(fmt.Sprintf("unable to initialize pipelines: %s", err.Error()))
	}
}

func GetAppContext() *AppContext {
	return appCtx
}

func (a *AppContext) DB() *sqlx.DB {
	return a.db
}

/*
func (a *AppContext) Cache() *cache.Cache {
	return a.cache
}
*/

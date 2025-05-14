package appcontext

import (
	atrservice "github.com/erlendromo/forsete-atr/src/business/usecase/service/atr_service"
	authservice "github.com/erlendromo/forsete-atr/src/business/usecase/service/auth_service"
	"github.com/erlendromo/forsete-atr/src/database"
)

var appCtx *AppContext

type AppContext struct {
	AuthService *authservice.AuthService
	ATRService  *atrservice.ATRService

	DB database.Database
	//cache *cache.Cache
}

func InitAppContext(db database.Database) {
	if appCtx == nil {
		appCtx = &AppContext{
			AuthService: authservice.NewAuthService(db),
			ATRService:  atrservice.NewATRService(db),
			DB:          db,
		}
	}
}

func GetAppContext() *AppContext {
	return appCtx
}

/*
func (a *AppContext) Cache() *cache.Cache {
	return a.cache
}
*/

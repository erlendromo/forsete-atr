package appcontext

import (
	modelrepository "github.com/erlendromo/forsete-atr/src/business/usecase/repository/model_repository"
	authservice "github.com/erlendromo/forsete-atr/src/business/usecase/service/auth_service"
	fileservice "github.com/erlendromo/forsete-atr/src/business/usecase/service/file_service"
	"github.com/jmoiron/sqlx"
)

var appCtx *AppContext

type AppContext struct {
	AuthService     *authservice.AuthService
	FileService     *fileservice.FileService
	ModelRepository *modelrepository.ModelRepository

	db *sqlx.DB
	//cache *cache.Cache
}

func InitAppContext(db *sqlx.DB) {
	if appCtx == nil {
		appCtx = &AppContext{
			AuthService:     authservice.NewAuthService(db),
			FileService:     fileservice.NewFileService(db),
			ModelRepository: modelrepository.NewModelRepository(db),
			db:              db,
		}
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

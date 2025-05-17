package appcontext

import (
	imagequerier "github.com/erlendromo/forsete-atr/src/business/usecase/querier/image"
	modelquerier "github.com/erlendromo/forsete-atr/src/business/usecase/querier/model"
	outputquerier "github.com/erlendromo/forsete-atr/src/business/usecase/querier/output"
	pipelinequerier "github.com/erlendromo/forsete-atr/src/business/usecase/querier/pipeline"
	sessionquerier "github.com/erlendromo/forsete-atr/src/business/usecase/querier/session"
	userquerier "github.com/erlendromo/forsete-atr/src/business/usecase/querier/user"
	imagerepo "github.com/erlendromo/forsete-atr/src/business/usecase/repository/image"
	modelrepo "github.com/erlendromo/forsete-atr/src/business/usecase/repository/model"
	outputrepo "github.com/erlendromo/forsete-atr/src/business/usecase/repository/output"
	pipelinerepo "github.com/erlendromo/forsete-atr/src/business/usecase/repository/pipeline"
	sessionrepo "github.com/erlendromo/forsete-atr/src/business/usecase/repository/session"
	userrepo "github.com/erlendromo/forsete-atr/src/business/usecase/repository/user"
	atrservice "github.com/erlendromo/forsete-atr/src/business/usecase/service/atr"
	authservice "github.com/erlendromo/forsete-atr/src/business/usecase/service/auth"
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
			AuthService: authservice.NewAuthService(
				userrepo.NewUserRepository(
					userquerier.NewSQLUserQuerier(
						db.DB(),
					),
				),

				sessionrepo.NewSessionRepository(
					sessionquerier.NewSQLSessionQuerier(
						db.DB(),
					),
				),
			),

			ATRService: atrservice.NewATRService(
				modelrepo.NewModelRepository(
					modelquerier.NewSQLModelQuerier(
						db.DB(),
					),
				),

				pipelinerepo.NewPipelineRepository(
					pipelinequerier.NewSQLPipelineQuerier(
						db.DB(),
					),
				),

				imagerepo.NewImageRepository(
					imagequerier.NewSQLImageQuerier(
						db.DB(),
					),
				),

				outputrepo.NewOutputRepository(
					outputquerier.NewSQLOutputQuerier(
						db.DB(),
					),
				),
			),

			DB: db,
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

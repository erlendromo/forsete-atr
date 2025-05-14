package router

import (
	"fmt"
	"net/http"

	_ "github.com/erlendromo/forsete-atr/docs"
	appcontext "github.com/erlendromo/forsete-atr/src/api/app_context"
	"github.com/erlendromo/forsete-atr/src/api/handler/v2/atr"
	"github.com/erlendromo/forsete-atr/src/api/handler/v2/auth"
	"github.com/erlendromo/forsete-atr/src/api/handler/v2/image"
	"github.com/erlendromo/forsete-atr/src/api/handler/v2/model"
	"github.com/erlendromo/forsete-atr/src/api/handler/v2/output"
	"github.com/erlendromo/forsete-atr/src/api/handler/v2/status"
	"github.com/erlendromo/forsete-atr/src/api/middleware"
	"github.com/erlendromo/forsete-atr/src/util"
	swaggo "github.com/swaggo/http-swagger/v2"
)

type Router interface {
	Serve() error
}

// @title			Forsete-ATR
// @version		v2
// @description	RESTful JSON-API for Automatic Text Recognition (ATR) developed as part of Bachelor Thesis "FORSETE" at NTNU Gjøvik.
func WithV2Endpoints(mux *http.ServeMux, appCtx *appcontext.AppContext) *http.ServeMux {
	// App Context
	authService := appCtx.AuthService
	atrService := appCtx.ATRService
	db := appCtx.DB

	// Swaggo
	mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodGet, util.SWAGGO_ENDPOINT),
		swaggo.Handler(swaggo.URL(util.SWAGGO_DOCS_ENDPOINT)),
	)

	// Auth
	mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodPost, util.REGISTER_ENDPOINT),
		auth.Register(authService),
	)
	mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodPost, util.LOGIN_ENDPOINT),
		auth.Login(authService),
	)
	mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodPost, util.LOGOUT_ENDPOINT),
		middleware.AuthMiddleware(authService, auth.Logout(authService)),
	)
	mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodPost, util.REFRESH_ENDPOINT),
		middleware.AuthMiddleware(authService, auth.RefreshSession(authService)),
	)
	mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodDelete, util.DELETE_USER_ENDPOINT),
		middleware.AuthMiddleware(authService, auth.DeleteUser(authService, atrService)),
	)

	// Images
	mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodGet, util.IMAGES_ENDPOINT),
		middleware.AuthMiddleware(authService, image.GetImages(atrService)),
	)
	mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodPost, util.UPLOAD_IMAGES_ENDPOINT),
		middleware.AuthMiddleware(authService, image.UploadImages(atrService)),
	)
	mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodGet, util.IMAGE_BY_ID_ENDPOINT),
		middleware.AuthMiddleware(authService, image.GetImageByID(atrService)),
	)
	mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodDelete, util.IMAGE_BY_ID_ENDPOINT),
		middleware.AuthMiddleware(authService, image.DeleteImageByID(atrService)),
	)
	mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodGet, util.IMAGE_DATA_ENDPOINT),
		middleware.AuthMiddleware(authService, image.GetImageData(atrService)),
	)

	// Outputs
	mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodGet, util.OUTPUTS_ENDPOINT),
		middleware.AuthMiddleware(authService, output.GetOutputsByImageID(atrService)),
	)
	mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodGet, util.OUTPUT_BY_ID_ENDPOINT),
		middleware.AuthMiddleware(authService, output.GetOutputByID(atrService)),
	)
	mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodPut, util.OUTPUT_BY_ID_ENDPOINT),
		middleware.AuthMiddleware(authService, output.UpdateOutputByID(atrService)),
	)
	mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodDelete, util.OUTPUT_BY_ID_ENDPOINT),
		middleware.AuthMiddleware(authService, output.DeleteOutputByID(atrService)),
	)
	mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodGet, util.OUTPUT_DATA_ENDPOINT),
		middleware.AuthMiddleware(authService, output.GetOutputData(atrService)),
	)

	// Models
	mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodGet, util.MODELS_ENDPOINT),
		model.GetModels(atrService.ModelRepo),
	)
	mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodGet, util.REGION_SEGMENTATION_ENDPOINT),
		model.GetModelsByType(atrService.ModelRepo, util.REGION_SEGMENTATION),
	)
	mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodGet, util.LINE_SEGMENTATION_ENDPOINT),
		model.GetModelsByType(atrService.ModelRepo, util.LINE_SEGMENTATION),
	)
	mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodGet, util.TEXT_RECOGNITION_ENDPOINT),
		model.GetModelsByType(atrService.ModelRepo, util.TEXT_RECOGNITION),
	)

	// ATR
	mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodPost, util.ATR_ENDPOINT),
		middleware.AuthMiddleware(authService, atr.Run(atrService)),
	)

	// Status
	mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodHead, util.STATUS_ENDPOINT),
		status.HeadStatus(db),
	)
	mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodGet, util.STATUS_ENDPOINT),
		status.GetStatus(db),
	)

	return mux
}

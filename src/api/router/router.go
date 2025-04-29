package router

import (
	"fmt"
	"net/http"

	_ "github.com/erlendromo/forsete-atr/docs"
	appcontext "github.com/erlendromo/forsete-atr/src/api/app_context"
	atrv1 "github.com/erlendromo/forsete-atr/src/api/handler/v1/atr"
	modelv1 "github.com/erlendromo/forsete-atr/src/api/handler/v1/model"
	statusv1 "github.com/erlendromo/forsete-atr/src/api/handler/v1/status"
	"github.com/erlendromo/forsete-atr/src/api/handler/v2/auth"
	"github.com/erlendromo/forsete-atr/src/api/handler/v2/image"
	"github.com/erlendromo/forsete-atr/src/api/handler/v2/model"
	"github.com/erlendromo/forsete-atr/src/api/handler/v2/status"
	"github.com/erlendromo/forsete-atr/src/api/middleware"
	"github.com/erlendromo/forsete-atr/src/util"
	swaggo "github.com/swaggo/http-swagger/v2"
)

type Router interface {
	Serve() error
}

// @title			Forsete-ATR
// @version		v1
// @description	RESTful JSON-API for Automatic Text Recognition (ATR) developed as part of Bachelor Thesis "FORSETE" at NTNU Gjøvik.
func WithV1Endpoints(mux *http.ServeMux) *http.ServeMux {
	// Swaggo
	mux.HandleFunc(
		fmt.Sprintf("GET %s", util.SWAGGO_ENDPOINT),
		swaggo.Handler(swaggo.URL(util.SWAGGO_DOCS_ENDPOINT)),
	)

	// Status
	mux.HandleFunc(fmt.Sprintf("HEAD %s", util.STATUS_ENDPOINT), statusv1.HeadStatus)
	mux.HandleFunc(fmt.Sprintf("GET %s", util.STATUS_ENDPOINT), statusv1.GetStatus)

	// Models
	mux.HandleFunc(fmt.Sprintf("GET %s", util.MODELS_ENDPOINT), modelv1.GetModels)
	mux.HandleFunc(fmt.Sprintf("GET %s", util.REGION_SEGMENTATION_ENDPOINT), modelv1.GetRegionSegmentationModels)
	mux.HandleFunc(fmt.Sprintf("POST %s", util.REGION_SEGMENTATION_ENDPOINT), modelv1.PostRegionSegmentationModel)
	mux.HandleFunc(fmt.Sprintf("GET %s", util.LINE_SEGMENTATION_ENDPOINT), modelv1.GetLineSegmentationModels)
	mux.HandleFunc(fmt.Sprintf("POST %s", util.LINE_SEGMENTATION_ENDPOINT), modelv1.PostLineSegmentationModel)
	mux.HandleFunc(fmt.Sprintf("GET %s", util.TEXT_RECOGNITION_ENDPOINT), modelv1.GetTextRecognitionModels)
	mux.HandleFunc(fmt.Sprintf("POST %s", util.TEXT_RECOGNITION_ENDPOINT), modelv1.PostTextRecognitionModel)

	// ATR
	mux.HandleFunc(fmt.Sprintf("POST %s", util.BASIC_DOCUMENTS_ENDPOINT), atrv1.PostBasicDocument)
	mux.HandleFunc(fmt.Sprintf("POST %s", util.TIPNOTE_DOCUMENTS_ENDPOINT), atrv1.PostTipnoteDocument)

	return mux
}

func WithV2Endpoints(mux *http.ServeMux) *http.ServeMux {
	// App Context
	appCtx := appcontext.GetAppContext()
	authService := appCtx.AuthService
	imageRepo := appCtx.ImageRepository
	modelRepo := appCtx.ModelRepository

	// Swaggo
	//mux.HandleFunc("GET /forsete-atr/v2/swaggo/", swaggo.Handler(swaggo.URL("/forsete-atr/v2/swaggo/doc.json")))

	// Auth
	mux.HandleFunc("POST /forsete-atr/v2/auth/register/", auth.Register(authService))
	mux.HandleFunc("POST /forsete-atr/v2/auth/login/", auth.Login(authService))
	mux.HandleFunc("POST /forsete-atr/v2/auth/logout/", middleware.AuthMiddleware(authService, auth.Logout(authService)))

	// Images
	//mux.HandleFunc("GET /forsete-atr/v2/images/", middleware.AuthMiddleware(authService, image.GetImages(imageRepo))
	mux.HandleFunc("POST /forsete-atr/v2/images/upload/", middleware.AuthMiddleware(authService, image.UploadImages(imageRepo)))

	// Outputs
	//mux.HandleFunc("GET /forsete-atr/v2/outputs/{imageID}", middleware.AuthMiddleware(authService, output.GetOutputByImageID(outputRepo))
	//mux.HandleFunc("GET /forsete-atr/v2/outputs/", middleware.AuthMiddleware(authService, output.GetOutputs(outputRepo))

	// Models
	mux.HandleFunc("GET /forsete-atr/v2/models/", model.GetModels(modelRepo))
	mux.HandleFunc("GET /forsete-atr/v2/models/region-segmentation-models/", model.GetModelsByType(modelRepo, "regionsegmentation"))
	mux.HandleFunc("GET /forsete-atr/v2/models/line-segmentation-models/", model.GetModelsByType(modelRepo, "linesegmentation"))
	mux.HandleFunc("GET /forsete-atr/v2/models/text-recognition-models/", model.GetModelsByType(modelRepo, "textrecognition"))

	// ATR
	//mux.HandleFunc("POST /forsete-atr/v2/atr/", middleware.AuthMiddleware(authService, atr.HTRflow(...)))

	// Status
	mux.HandleFunc("HEAD /forsete-atr/v2/status/", status.HeadStatus())
	mux.HandleFunc("GET /forsete-atr/v2/status/", status.GetStatus())

	return mux
}

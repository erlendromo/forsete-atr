package router

import (
	"fmt"
	"net/http"

	_ "github.com/erlendromo/forsete-atr/docs"
	"github.com/erlendromo/forsete-atr/src/api/handler/atr"
	"github.com/erlendromo/forsete-atr/src/api/handler/model"
	"github.com/erlendromo/forsete-atr/src/api/handler/status"
	"github.com/erlendromo/forsete-atr/src/util"
	swaggo "github.com/swaggo/http-swagger/v2"
)

type Router interface {
	Serve() error
}

// @title			Forsete-ATR
// @version		v1
// @description	RESTful JSON-API for Automatic Text Recognition (ATR) developed as part of Bachelor Thesis "FORSETE" at NTNU Gjøvik.
func WithEndpoints(mux *http.ServeMux) *http.ServeMux {
	// Swaggo
	mux.HandleFunc(
		fmt.Sprintf("GET %s", util.SWAGGO_ENDPOINT),
		swaggo.Handler(swaggo.URL(util.SWAGGO_DOCS_ENDPOINT)),
	)

	// Status
	mux.HandleFunc(fmt.Sprintf("HEAD %s", util.STATUS_ENDPOINT), status.HeadStatus)
	mux.HandleFunc(fmt.Sprintf("GET %s", util.STATUS_ENDPOINT), status.GetStatus)

	// Models
	mux.HandleFunc(fmt.Sprintf("GET %s", util.MODELS_ENDPOINT), model.GetModels)
	mux.HandleFunc(fmt.Sprintf("GET %s", util.REGION_SEGMENTATION_ENDPOINT), model.GetRegionSegmentationModels)
	mux.HandleFunc(fmt.Sprintf("POST %s", util.REGION_SEGMENTATION_ENDPOINT), model.PostRegionSegmentationModel)
	mux.HandleFunc(fmt.Sprintf("GET %s", util.LINE_SEGMENTATION_ENDPOINT), model.GetLineSegmentationModels)
	mux.HandleFunc(fmt.Sprintf("POST %s", util.LINE_SEGMENTATION_ENDPOINT), model.PostLineSegmentationModel)
	mux.HandleFunc(fmt.Sprintf("GET %s", util.TEXT_RECOGNITION_ENDPOINT), model.GetTextRecognitionModels)
	mux.HandleFunc(fmt.Sprintf("POST %s", util.TEXT_RECOGNITION_ENDPOINT), model.PostTextRecognitionModel)

	// ATR
	mux.HandleFunc(fmt.Sprintf("POST %s", util.BASIC_DOCUMENTS_ENDPOINT), atr.PostBasicDocument)
	mux.HandleFunc(fmt.Sprintf("POST %s", util.TIPNOTE_DOCUMENTS_ENDPOINT), atr.PostTipnoteDocument)

	return mux
}

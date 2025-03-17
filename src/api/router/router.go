package router

import (
	"fmt"
	"net/http"

	_ "github.com/erlendromo/forsete-atr/docs"
	"github.com/erlendromo/forsete-atr/src/api/handler"
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

	// Models
	mux.HandleFunc(fmt.Sprintf("GET %s", util.MODELS_ENDPOINT), handler.GetModels)
	mux.HandleFunc(fmt.Sprintf("GET %s", util.REGION_SEGMENTATION_ENDPOINT), handler.GetRegionSegmentationModels)
	mux.HandleFunc(fmt.Sprintf("POST %s", util.REGION_SEGMENTATION_ENDPOINT), handler.PostRegionSegmentationModel)
	mux.HandleFunc(fmt.Sprintf("GET %s", util.LINE_SEGMENTATION_ENDPOINT), handler.GetLineSegmentationModels)
	mux.HandleFunc(fmt.Sprintf("POST %s", util.LINE_SEGMENTATION_ENDPOINT), handler.PostLineSegmentationModel)
	mux.HandleFunc(fmt.Sprintf("GET %s", util.TEXT_RECOGNITION_ENDPOINT), handler.GetTextRecognitionModels)
	//mux.HandleFunc(fmt.Sprintf("POST %s", util.TEXT_RECOGNITION_ENDPOINT), handler.PostTextRecognitionModel)

	// ATR
	mux.HandleFunc(fmt.Sprintf("POST %s", util.BASIC_DOCUMENTS_ENDPOINT), handler.PostBasicDocument)
	//mux.HandleFunc(fmt.Sprintf("POST %s", util.TIPNOTE_ENDPOINT), handler.PostTipnoteDocument)

	return mux
}

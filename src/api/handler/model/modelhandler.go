package model

import (
	"net/http"

	"github.com/erlendromo/forsete-atr/src/domain/model"
	"github.com/erlendromo/forsete-atr/src/util"
)

// ModelsResponse
//
//	@Summary		ModelsResponse
//	@Description	Json-Response for models
//
//	@Tags			Models
type ModelsResponse struct {
	RegionSegmentationModels []model.Model `json:"region_segmentation_models,omitempty"`
	LineSegmentationModels   []model.Model `json:"line_segmentation_models,omitempty"`
	TextRecognitionModels    []model.Model `json:"text_recognition_models,omitempty"`
}

// GetModels
//
//	@Summary		Models
//	@Description	Retrieve all active models
//	@Tags			Models
//	@Produce		json
//	@Success		200	{object}	ModelsResponse
//	@Router			/forsete-atr/v1/models/ [get]
func GetModels(w http.ResponseWriter, r *http.Request) {
	util.JSON(w, http.StatusOK, &ModelsResponse{
		RegionSegmentationModels: model.ModelsByType(util.REGION_SEGMENTATION),
		LineSegmentationModels:   model.ModelsByType(util.LINE_SEGMENTATION),
		TextRecognitionModels:    model.ModelsByType(util.TEXT_RECOGNITION),
	})
}

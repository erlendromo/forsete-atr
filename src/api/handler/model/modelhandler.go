package model

import (
	"net/http"

	"github.com/erlendromo/forsete-atr/src/domain/modelstore"
	"github.com/erlendromo/forsete-atr/src/domain/modelstore/model"
	"github.com/erlendromo/forsete-atr/src/util"
)

// ModelsResponse
//
//	@Summary		ModelsResponse
//	@Description	Json-Response for models
//
//	@Tags			Models
type ModelsResponse struct {
	RegionSegmentationModels []*model.Model `json:"region_segmentation_models,omitempty"`
	LineSegmentationModels   []*model.Model `json:"line_segmentation_models,omitempty"`
	TextRecognitionModels    []*model.Model `json:"text_recognition_models,omitempty"`
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
		RegionSegmentationModels: modelstore.GetModelstore().ModelsByType(util.REGION_SEGMENTATION),
		LineSegmentationModels:   modelstore.GetModelstore().ModelsByType(util.LINE_SEGMENTATION),
		TextRecognitionModels:    modelstore.GetModelstore().ModelsByType(util.TEXT_RECOGNITION),
	})
}

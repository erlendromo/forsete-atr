package model

import (
	"fmt"
	"net/http"

	modelrepository "github.com/erlendromo/forsete-atr/src/business/usecase/model_repository"
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
	util.EncodeJSON(w, http.StatusOK, &ModelsResponse{
		RegionSegmentationModels: modelstore.GetModelstore().ModelsByType(util.REGION_SEGMENTATION),
		LineSegmentationModels:   modelstore.GetModelstore().ModelsByType(util.LINE_SEGMENTATION),
		TextRecognitionModels:    modelstore.GetModelstore().ModelsByType(util.TEXT_RECOGNITION),
	})
}

func GetModelsV2(m *modelrepository.ModelRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		models, err := m.AllModels(r.Context())
		if err != nil {
			util.NewInternalErrorLog("MODELS", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		util.EncodeJSON(w, http.StatusOK, models)
	}
}

func GetModelsByTypeV2(m *modelrepository.ModelRepository, modelType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		models, err := m.ModelsByType(r.Context(), modelType)
		if err != nil {
			util.NewInternalErrorLog("MODELS BY TYPE", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		util.EncodeJSON(w, http.StatusOK, models)
	}
}

package model

import (
	"fmt"
	"net/http"

	_ "github.com/erlendromo/forsete-atr/src/business/domain/model"
	modelrepository "github.com/erlendromo/forsete-atr/src/business/usecase/repository/model_repository"
	"github.com/erlendromo/forsete-atr/src/util"
)

// GetModels
//
//	@Summary		Get models
//	@Description	Get all models.
//	@Tags			Models
//	@Produce		json
//	@Success		200	{object}	[]model.Model
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v2/models/ [get]
func GetModels(m *modelrepository.ModelRepository) http.HandlerFunc {
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

// GetModelsByType
//
//	@Summary		Get models by type
//	@Description	Get models by type.
//	@Tags			Models
//	@Produce		json
//	@Success		200	{object}	[]model.Model
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v2/models/region-segmentation-models/ [get]
//	@Router			/forsete-atr/v2/models/line-segmentation-models/ [get]
//	@Router			/forsete-atr/v2/models/text-recognition-models/ [get]
func GetModelsByType(m *modelrepository.ModelRepository, modelType string) http.HandlerFunc {
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

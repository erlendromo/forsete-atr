package model

import (
	"fmt"
	"net/http"

	modelrepository "github.com/erlendromo/forsete-atr/src/business/usecase/repository/model_repository"
	"github.com/erlendromo/forsete-atr/src/util"
)

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

package handler

import (
	"net/http"

	"github.com/erlendromo/forsete-atr/src/domain/model"
	"github.com/erlendromo/forsete-atr/src/util"
)

type ModelsResponse struct {
	RegionSegmentationModels []model.Model `json:"region_segmentation_models,omitempty"`
	LineSegmentationModels   []model.Model `json:"line_segmentation_models,omitempty"`
	TextRecognitionModels    []model.Model `json:"text_recognition_models,omitempty"`
}

func GetModels(w http.ResponseWriter, r *http.Request) {
	util.JSON(w, http.StatusOK, &ModelsResponse{
		RegionSegmentationModels: model.ModelsByType("regionsegmentation"),
		LineSegmentationModels:   model.ModelsByType("linesegmentation"),
		TextRecognitionModels:    model.ModelsByType("textrecognition"),
	})
}

func GetRegionSegmentationModels(w http.ResponseWriter, r *http.Request) {
	util.JSON(w, http.StatusOK, &ModelsResponse{
		RegionSegmentationModels: model.ModelsByType("regionsegmentation"),
	})
}

func PostRegionSegmentationModel(w http.ResponseWriter, r *http.Request) {
	modelFile, modelHeader, err := r.FormFile("model")
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, _ = modelFile, modelHeader

	util.JSON(w, http.StatusNoContent, nil)
}

func GetLineSegmentationModels(w http.ResponseWriter, r *http.Request) {
	util.JSON(w, http.StatusOK, &ModelsResponse{
		LineSegmentationModels: model.ModelsByType("linesegmentation"),
	})
}

func PostLineSegmentationModel(w http.ResponseWriter, r *http.Request) {
	modelFile, modelHeader, err := r.FormFile("model")
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, _ = modelFile, modelHeader

	util.JSON(w, http.StatusNoContent, nil)
}

func GetTextRecognitionModels(w http.ResponseWriter, r *http.Request) {
	util.JSON(w, http.StatusOK, &ModelsResponse{
		TextRecognitionModels: model.ModelsByType("textrecognition"),
	})
}

func PostTextRecognitionModel(w http.ResponseWriter, r *http.Request) {
	modelFile, modelHeader, err := r.FormFile("model")
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, _ = modelFile, modelHeader

	util.JSON(w, http.StatusNoContent, nil)
}

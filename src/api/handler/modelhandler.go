package handler

import (
	"errors"
	"net/http"
	"os"
	"strings"

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
		RegionSegmentationModels: model.ModelsByType("regionsegmentation"),
		LineSegmentationModels:   model.ModelsByType("linesegmentation"),
		TextRecognitionModels:    model.ModelsByType("textrecognition"),
	})
}

// GetRegionSegmentationModels
//
//	@Summary		RegionSegmentationModels
//	@Description	Retrieve all active region segmentation models
//	@Tags			RegionSegmentationModels
//	@Produce		json
//	@Success		200	{object}	ModelsResponse
//	@Router			/forsete-atr/v1/models/region-segmentation-models/ [get]
func GetRegionSegmentationModels(w http.ResponseWriter, r *http.Request) {
	util.JSON(w, http.StatusOK, &ModelsResponse{
		RegionSegmentationModels: model.ModelsByType("regionsegmentation"),
	})
}

// PostRegionSegmentationModel
//
//	@Summary		RegionSegmentationModels
//	@Description	Add a region segmentation model
//	@Tags			RegionSegmentationModels
//	@Accept			mpfd
//	@Param			model_name	formData	string	required	"Name of the model"
//	@Param			model		formData	file	required	"model.pt"
//	@Produce		json
//	@Success		204
//	@Failure		404	{object}	util.ErrorResponse
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v1/models/region-segmentation-models/ [post]
func PostRegionSegmentationModel(w http.ResponseWriter, r *http.Request) {
	modelName := strings.ToLower(r.FormValue("model_name"))
	if len(modelName) < 3 {
		util.ERROR(w, http.StatusBadRequest, errors.New("improper model_name, must be atleast 3 characters long"))
		return
	}

	if _, err := os.Open("models/regionsegmentation/model.pt"); err == nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("model already exists"))
		return
	}

	modelFile, _, err := r.FormFile("model")
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if err := model.AddYoloModel(modelName, "regionsegmentation", modelFile); err != nil {
		util.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	util.JSON(w, http.StatusNoContent, nil)
}

// GetLineSegmentationModels
//
//	@Summary		LineSegmentationModels
//	@Description	Retrieve all active line segmentation models
//	@Tags			LineSegmentationModels
//	@Produce		json
//	@Success		200	{object}	ModelsResponse
//	@Router			/forsete-atr/v1/models/line-segmentation-models/ [get]
func GetLineSegmentationModels(w http.ResponseWriter, r *http.Request) {
	util.JSON(w, http.StatusOK, &ModelsResponse{
		LineSegmentationModels: model.ModelsByType("linesegmentation"),
	})
}

// PostLineSegmentationModel
//
//	@Summary		LineSegmentationModels
//	@Description	Add a line segmentation model
//	@Tags			LineSegmentationModels
//	@Accept			mpfd
//	@Param			model_name	formData	string	required	"Name of the model"
//	@Param			model		formData	file	required	"model.pt"
//	@Produce		json
//	@Success		204
//	@Failure		404	{object}	util.ErrorResponse
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v1/models/line-segmentation-models/ [post]
func PostLineSegmentationModel(w http.ResponseWriter, r *http.Request) {
	modelName := strings.ToLower(r.FormValue("model_name"))
	if len(modelName) < 3 {
		util.ERROR(w, http.StatusBadRequest, errors.New("improper model_name, must be atleast 3 characters long"))
		return
	}

	if _, err := os.Open("models/linesegmentation/model.pt"); err == nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("model already exists"))
		return
	}

	modelFile, _, err := r.FormFile("model")
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, err)
		return
	}

	defer modelFile.Close()

	if err := model.AddYoloModel(modelName, "linesegmentation", modelFile); err != nil {
		util.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	util.JSON(w, http.StatusNoContent, nil)
}

// GetTextRecognitionModels
//
//	@Summary		TextRecognitionModels
//	@Description	Retrieve all active text recognition models
//	@Tags			TextRecognitionModels
//	@Produce		json
//	@Success		200	{object}	ModelsResponse
//	@Router			/forsete-atr/v1/models/text-recognition-models/ [get]
func GetTextRecognitionModels(w http.ResponseWriter, r *http.Request) {
	util.JSON(w, http.StatusOK, &ModelsResponse{
		TextRecognitionModels: model.ModelsByType("textrecognition"),
	})
}

// PostTextRecognitionModel
//
//	@Summary		TextRecognitionModels
//	@Description	Add a text recognition model
//	@Tags			TextRecognitionModels
//	@Accept			mpfd
//	@Param			model_name			formData	string	required	"Name of the model"
//	@Param			model				formData	file	required	"model.safetensors"
//	@Param			config				formData	file	required	"config.json"
//	@Param			generation_config	formData	file	required	"generation_config.json"
//	@Param			merges				formData	file	required	"merges.txt"
//	@Param			preprocessor_config	formData	file	required	"preprocessor_config-json"
//	@Param			special_tokens_map	formData	file	required	"special_tokens_map.json"
//	@Param			tokenizer			formData	file	required	"tokenizer.json"
//	@Param			tokenizer_config	formData	file	required	"tokenizer_config.json"
//	@Param			vocab				formData	file	required	"vocab.json"
//	@Produce		json
//	@Success		204
//	@Failure		404	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v1/models/text-recognition-models/ [post]
func PostTextRecognitionModel(w http.ResponseWriter, r *http.Request) {
	modelName := strings.ToLower(r.FormValue("model_name"))
	if len(modelName) < 3 {
		util.ERROR(w, http.StatusBadRequest, errors.New("Improper model_name"))
		return
	}

	modelFile, modelHeader, err := r.FormFile("model")
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, _ = modelFile, modelHeader

	util.JSON(w, http.StatusNoContent, nil)
}

package model

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/erlendromo/forsete-atr/src/domain/model"
	"github.com/erlendromo/forsete-atr/src/util"
)

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
		TextRecognitionModels: model.ModelsByType(util.TEXT_RECOGNITION),
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
//	@Param			preprocessor_config	formData	file	required	"preprocessor_config.json"
//	@Param			special_tokens_map	formData	file	required	"special_tokens_map.json"
//	@Param			tokenizer			formData	file	required	"tokenizer.json"
//	@Param			tokenizer_config	formData	file	required	"tokenizer_config.json"
//	@Param			vocab				formData	file	required	"vocab.json"
//	@Produce		json
//	@Success		204
//	@Failure		404	{object}	util.ErrorResponse
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v1/models/text-recognition-models/ [post]
func PostTextRecognitionModel(w http.ResponseWriter, r *http.Request) {

	// Only accept alphanumeric and dashes

	modelName := string(
		regexp.MustCompile(
			"[^a-zA-Z0-9 -]+",
		).ReplaceAll(
			[]byte(strings.ToLower(r.FormValue("model_name"))),
			[]byte(""),
		),
	)

	// Initial checks

	if len(modelName) < 3 && len(modelName) > 50 {
		util.ERROR(w, http.StatusBadRequest, errors.New("invalid model_name, must be between 3 and 50 characters long"))
		return
	}

	if _, err := os.ReadDir(fmt.Sprintf("%s/%s/%s", util.MODELS, util.TEXT_RECOGNITION, modelName)); err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("model already exists"))
		return
	}

	// Handling request-data

	modelFile, modelHeader, err := r.FormFile("model")
	if !strings.Contains(modelHeader.Filename, "safetensors") || err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("invalid model, must be a .safetensors file"))
		return
	}

	defer modelFile.Close()

	configFile, configHeader, err := r.FormFile("config")
	if !strings.Contains(configHeader.Filename, "json") || err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("invalid config, must be a .json file"))
		return
	}

	defer configFile.Close()

	generationConfigFile, generationConfigHeader, err := r.FormFile("generation_config")
	if !strings.Contains(generationConfigHeader.Filename, "json") || err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("invalid generation_config, must be a .json file"))
		return
	}

	defer generationConfigFile.Close()

	mergesFile, mergesHeader, err := r.FormFile("merges")
	if !strings.Contains(mergesHeader.Filename, "txt") || err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("invalid merges, must be a .txt file"))
		return
	}

	defer mergesFile.Close()

	preprocessorConfigFile, preprocessorConfigHeader, err := r.FormFile("preprocessor_config")
	if !strings.Contains(preprocessorConfigHeader.Filename, "json") || err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("invalid preprocessor_config, must be a .json file"))
		return
	}

	defer preprocessorConfigFile.Close()

	specialTokensMapFile, specialTokensMapHeader, err := r.FormFile("special_tokens_map")
	if !strings.Contains(specialTokensMapHeader.Filename, "json") || err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("invalid special_tokens_map, must be a .json file"))
		return
	}

	defer specialTokensMapFile.Close()

	tokenizerFile, tokenizerHeader, err := r.FormFile("tokenizer")
	if !strings.Contains(tokenizerHeader.Filename, "json") || err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("invalid tokenizer, must be a .json file"))
		return
	}

	defer tokenizerFile.Close()

	tokenizerConfigFile, tokenizerConfigHeader, err := r.FormFile("tokenizer_config")
	if !strings.Contains(tokenizerConfigHeader.Filename, "json") || err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("invalid tokenizer_config, must be a .json file"))
		return
	}

	defer tokenizerConfigFile.Close()

	vocabFile, vocabHeader, err := r.FormFile("vocab")
	if !strings.Contains(vocabHeader.Filename, "json") || err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("invalid vocab, must be a .json file"))
		return
	}

	defer vocabFile.Close()

	files := map[string]multipart.File{
		"model.safetensors":        modelFile,
		"config.json":              configFile,
		"generation_config.json":   generationConfigFile,
		"merges.txt":               mergesFile,
		"preprocessor_config.json": preprocessorConfigFile,
		"special_tokens_map.json":  specialTokensMapFile,
		"tokenizer.json":           tokenizerFile,
		"tokenizer_config.json":    tokenizerConfigFile,
		"vocab.json":               vocabFile,
	}

	if err := model.AddTrOCRModel(modelName, util.TEXT_RECOGNITION, files); err != nil {
		fmt.Printf("\n%sADD TROCRMODEL ERROR%s\n%s\n", util.RED, util.RESET, err.Error())
		util.ERROR(w, http.StatusInternalServerError, errors.New(util.INTERNAL_SERVER_ERROR))
		return
	}

	util.JSON(w, http.StatusNoContent, nil)
}

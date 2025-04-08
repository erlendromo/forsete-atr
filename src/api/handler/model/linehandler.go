package model

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/erlendromo/forsete-atr/src/domain/modelstore"
	"github.com/erlendromo/forsete-atr/src/util"
)

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
		LineSegmentationModels: modelstore.GetModelstore().ModelsByType(util.LINE_SEGMENTATION),
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
	// Only accept alphanumeric and dashes
	modelName := string(
		regexp.MustCompile(
			"[^a-zA-Z0-9 -]+",
		).ReplaceAll(
			[]byte(strings.ToLower(r.FormValue("model_name"))),
			[]byte(""),
		),
	)

	if len(modelName) < 3 && len(modelName) > 50 {
		util.ERROR(w, http.StatusBadRequest, errors.New("invalid model_name, must be between 3 and 50 characters long"))
		return
	}

	if _, err := os.ReadDir(fmt.Sprintf("%s/%s/%s", util.MODELS, util.LINE_SEGMENTATION, modelName)); err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("model already exists"))
		return
	}

	modelFile, modelHeader, err := r.FormFile("model")
	if !strings.Contains(modelHeader.Filename, ".pt") || err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("invalid model, must be a .pt file"))
		return
	}

	defer modelFile.Close()

	files := map[string]multipart.File{
		"model.pt": modelFile,
	}

	if err := modelstore.GetModelstore().AddModel(modelName, util.LINE_SEGMENTATION, files); err != nil {
		util.NewInternalErrorLog("ADD YOLOMODEL ERROR", err).PrintLog("SERVER ERROR")
		util.ERROR(w, http.StatusInternalServerError, errors.New(util.INTERNAL_SERVER_ERROR))
		return
	}

	util.JSON(w, http.StatusNoContent, nil)
}

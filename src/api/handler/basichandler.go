package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/erlendromo/forsete-atr/src/domain/image"
	"github.com/erlendromo/forsete-atr/src/domain/model"
	"github.com/erlendromo/forsete-atr/src/domain/pipeline"
	"github.com/erlendromo/forsete-atr/src/util"
)

// PostBasicDocument
//
//	@Summary		Run ATR on image-file
//	@Description	ATR
//	@Accept			mpfd
//	@Param			image					formData	file	required	"imagefile (png)"
//	@Param			line_segmentation_model	formData	string	required	"chosen line segmentation model"
//	@Param			text_recognition_model	formData	string	required	"chosen text recognition model"
//	@Produce		json
//	@Success		200
//	@Failure		400	{object}	util.ErrorResponse
//	@Failure		422	{object}	util.ErrorResponse
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v1/atr/basic-documents/ [post]
func PostBasicDocument(w http.ResponseWriter, r *http.Request) {

	// Handling request-data (image, segmentationmodel, textrecognitionmodel)

	imageFile, imageHeader, err := r.FormFile("image")
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, err)
		return
	}

	lineModel, found := model.Path(r.FormValue("line_segmentation_model"))
	if !found {
		util.ERROR(w, http.StatusBadRequest, errors.New("line_segmentation_model invalid"))
		return
	}

	textModel, found := model.Path(r.FormValue("text_recognition_model"))
	if !found {
		util.ERROR(w, http.StatusBadRequest, errors.New("text_recognition_model invalid"))
		return
	}

	// Process image and create yaml

	imagePath, err := image.ProcessImage(imageFile, imageHeader)
	if err != nil {
		util.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	yamlPath, err := pipeline.NewBasicPipeline(lineModel, textModel).Encode("tmp/yaml", "basic.yaml")
	if err != nil {
		util.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Execute htrflow

	if err := exec.Command("/bin/bash", "scripts/htrflow.sh", yamlPath, imagePath).Run(); err != nil {
		util.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// Read and decode json-output

	jsonOutput, err := os.Open(fmt.Sprintf("tmp/outputs/images/%s.json", strings.Split(strings.Split(imagePath, "/")[2], ".")[0]))
	if err != nil {
		util.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	var out any
	if err := json.NewDecoder(jsonOutput).Decode(&out); err != nil {
		util.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// Respond to client

	util.JSON(w, http.StatusOK, out)
}

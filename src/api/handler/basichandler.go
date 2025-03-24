package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/erlendromo/forsete-atr/src/config"
	"github.com/erlendromo/forsete-atr/src/domain/image"
	"github.com/erlendromo/forsete-atr/src/domain/model"
	"github.com/erlendromo/forsete-atr/src/domain/pipeline"
	"github.com/erlendromo/forsete-atr/src/util"
)

// ATRResponse
//
//	@Summary		ATRResponse
//	@Description	Json-Response for ATR
type ATRResponse struct {
	Filename  string `json:"file_name"`
	Imagename string `json:"image_name"`
	Label     string `json:"label"`
	Contains  []struct {
		Segment struct {
			BBox struct {
				XMin int `json:"xmin"`
				YMin int `json:"ymin"`
				XMax int `json:"xmax"`
				YMax int `json:"ymax"`
			} `json:"bbox"`
			Polygon struct {
				Points []struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"points"`
			} `json:"polygon"`
			Score      float64  `json:"score"`
			ClassLabel string   `json:"class_label"`
			OrigShape  []int    `json:"orig_shape"`
			Data       struct{} `json:"data,omitempty"`
		} `json:"segment"`
		TextResult struct {
			Texts  []string  `json:"texts"`
			Scores []float64 `json:"scores"`
			Label  string    `json:"label"`
		} `json:"text_result"`
	} `json:"contains"`
}

// PostBasicDocument
//
//	@Summary		ATR
//	@Description	Run ATR on image-file
//	@Tags			ATR
//	@Accept			mpfd
//	@Param			image					formData	file	required	"png"
//	@Param			line_segmentation_model	formData	string	required	"name of line segmentation model"
//	@Param			text_recognition_model	formData	string	required	"name of text recognition model"
//	@Produce		json
//	@Success		200	{object}	ATRResponse
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

	yamlPath, err := pipeline.NewBasicPipeline(lineModel, textModel, config.GetConfig().DEVICE).Encode("tmp/yaml", "basic.yaml")
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

	var atrResponse ATRResponse
	if err := json.NewDecoder(jsonOutput).Decode(&atrResponse); err != nil {
		util.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// Respond to client

	util.JSON(w, http.StatusOK, atrResponse)
}

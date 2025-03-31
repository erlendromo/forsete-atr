package atr

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
//	@Param			image					formData	file	required	"png, jpg, jpeg"
//	@Param			line_segmentation_model	formData	string	required	"name of line segmentation model"
//	@Param			text_recognition_model	formData	string	required	"name of text recognition model"
//	@Produce		json
//	@Success		200	{object}	ATRResponse
//	@Failure		400	{object}	util.ErrorResponse
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v1/atr/basic-documents/ [post]
func PostBasicDocument(w http.ResponseWriter, r *http.Request) {

	// Handling request-data (image, linesegmentationmodel, textrecognitionmodel)

	imageFile, imageHeader, err := r.FormFile("image")
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("invalid image"))
		return
	}

	if !strings.Contains(imageHeader.Filename, "png") && !strings.Contains(imageHeader.Filename, "jpg") && !strings.Contains(imageHeader.Filename, "jpeg") {
		util.ERROR(w, http.StatusBadRequest, errors.New("invalid image format, should be png, jpg or jpeg"))
		return
	}

	lineModel, found := model.Path(r.FormValue("line_segmentation_model"))
	if !found {
		util.ERROR(w, http.StatusBadRequest, errors.New("invalid line_segmentation_model, see /forsete-atr/v1/models/line-segmentation-models/ for valid models"))
		return
	}

	textModel, found := model.Path(r.FormValue("text_recognition_model"))
	if !found {
		util.ERROR(w, http.StatusBadRequest, errors.New("invalid text_recognition_model, see /forsete-atr/v1/models/text-recognition-models/ for valid models"))
		return
	}

	// Process image and yaml

	imagePath, err := image.ProcessImage(imageFile, imageHeader)
	if err != nil {
		fmt.Printf("\n%sIMAGE ERROR%s\n%s\n", util.RED, util.RESET, err.Error())
		util.ERROR(w, http.StatusInternalServerError, errors.New(util.INTERNAL_SERVER_ERROR))
		return
	}

	yamlPath, err := pipeline.NewBasicPipeline(lineModel, textModel, config.GetConfig().DEVICE).Encode("tmp/yaml", "basic.yaml")
	if err != nil {
		fmt.Printf("\n%sPIPELINE ERROR%s\n%s\n", util.RED, util.RESET, err.Error())
		util.ERROR(w, http.StatusInternalServerError, errors.New(util.INTERNAL_SERVER_ERROR))
		return
	}

	// Execute htrflow

	cmd := exec.Command("/bin/bash", "scripts/htrflow.sh", yamlPath, imagePath)
	if output, err := cmd.CombinedOutput(); err != nil {
		fmt.Printf("\n%sHTRFLOW ERROR%s\n%s\n", util.RED, util.RESET, output)
		util.ERROR(w, http.StatusInternalServerError, errors.New(util.INTERNAL_SERVER_ERROR))
		return
	}

	// Read and decode json-output

	jsonOutput, err := os.Open(fmt.Sprintf("tmp/outputs/images/%s.json", strings.Split(strings.Split(imagePath, "/")[2], ".")[0]))
	if err != nil {
		fmt.Printf("\n%sREAD RESPONSE ERROR%s\n%s\n", util.RED, util.RESET, err.Error())
		util.ERROR(w, http.StatusInternalServerError, errors.New(util.INTERNAL_SERVER_ERROR))
		return
	}

	var atrResponse ATRResponse
	if err := json.NewDecoder(jsonOutput).Decode(&atrResponse); err != nil {
		fmt.Printf("\n%sENCODE RESPONSE ERROR%s\n%s\n", util.RED, util.RESET, err.Error())
		util.ERROR(w, http.StatusInternalServerError, errors.New(util.INTERNAL_SERVER_ERROR))
		return
	}

	// Respond to client

	util.JSON(w, http.StatusOK, atrResponse)
}

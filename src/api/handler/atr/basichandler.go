package atr

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/erlendromo/forsete-atr/src/config"
	"github.com/erlendromo/forsete-atr/src/domain/htrflow"
	"github.com/erlendromo/forsete-atr/src/domain/image"
	"github.com/erlendromo/forsete-atr/src/domain/modelstore"
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

	lineModel, found := modelstore.GetModelstore().PathToModel(r.FormValue("line_segmentation_model"))
	if !found {
		util.ERROR(w, http.StatusBadRequest, errors.New("invalid line_segmentation_model, see /forsete-atr/v1/models/line-segmentation-models/ for valid models"))
		return
	}

	textModel, found := modelstore.GetModelstore().PathToModel(r.FormValue("text_recognition_model"))
	if !found {
		util.ERROR(w, http.StatusBadRequest, errors.New("invalid text_recognition_model, see /forsete-atr/v1/models/text-recognition-models/ for valid models"))
		return
	}

	// Process image

	image, err := image.NewImage(imageHeader.Filename, imageFile)
	if err != nil {
		util.NewInternalErrorLog("NEW IMAGE ERROR", err).PrintLog("SERVER ERROR")
		util.ERROR(w, http.StatusInternalServerError, errors.New(util.INTERNAL_SERVER_ERROR))
		return
	}

	imagePath, err := image.CreateLocalImage()
	if err != nil {
		util.NewInternalErrorLog("CREATE LOCAL IMAGE ERROR", err).PrintLog("SERVER ERROR")
		util.ERROR(w, http.StatusInternalServerError, errors.New(util.INTERNAL_SERVER_ERROR))
		return
	}

	// Process pipeline

	pipeline, err := pipeline.NewPipeline(
		config.GetConfig().DEVICE,
		fmt.Sprintf(
			"%s_%s",
			r.FormValue("line_segmentation_model"),
			r.FormValue("text_recognition_model"),
		),
	)
	if err != nil {
		util.NewInternalErrorLog("PIPELINE ERROR", err).PrintLog("SERVER ERROR")
		util.ERROR(w, http.StatusInternalServerError, errors.New(util.INTERNAL_SERVER_ERROR))
		return
	}

	pipeline.AppendYoloStep(
		lineModel,
	).AppendTrOCRStep(
		textModel,
	).AppendOrderStep(
		"OrderLines",
	).AppendExportStep(
		"json",
	)

	yamlPath, err := pipeline.CreateLocalYaml()
	if err != nil {
		util.NewInternalErrorLog("PIPELINE ERROR", err).PrintLog("SERVER ERROR")
		util.ERROR(w, http.StatusInternalServerError, errors.New(util.INTERNAL_SERVER_ERROR))
		return
	}

	// Execute htrflow

	htrflow := htrflow.NewHTRflow(
		yamlPath,
		imagePath,
		fmt.Sprintf("tmp/outputs/images/%s.json", strings.Split(strings.Split(imagePath, "/")[2], ".")[0]),
	)

	outputFile, err := htrflow.Run()
	if err != nil {
		util.NewInternalErrorLog("HTRFLOW ERROR", err).PrintLog("SERVER ERROR")
		util.ERROR(w, http.StatusInternalServerError, errors.New(util.INTERNAL_SERVER_ERROR))
		return
	}

	// Decode and write response

	var atrResponse ATRResponse
	if err := json.NewDecoder(outputFile).Decode(&atrResponse); err != nil {
		util.NewInternalErrorLog("DECODE RESPONSE ERROR", err).PrintLog("SERVER ERROR")
		util.ERROR(w, http.StatusInternalServerError, errors.New(util.INTERNAL_SERVER_ERROR))
		return
	}

	util.JSON(w, http.StatusOK, atrResponse)
}

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

// PostTipnoteDocument
//
//	@Summary		ATR
//	@Description	Run ATR on image-file
//	@Tags			ATR
//	@Accept			mpfd
//	@Param			image						formData	file	required	"png, jpg, jpeg"
//	@Param			region_segmentation_model	formData	string	required	"name of the region segmentation model"
//	@Param			line_segmentation_model		formData	string	required	"name of line segmentation model"
//	@Param			text_recognition_model		formData	string	required	"name of text recognition model"
//	@Produce		json
//	@Success		200	{object}	ATRResponse
//	@Failure		400	{object}	util.ErrorResponse
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v1/atr/tipnote-documents/ [post]
func PostTipnoteDocument(w http.ResponseWriter, r *http.Request) {

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

	regionModel, found := model.Path(r.FormValue("region_segmentation_model"))
	if !found {
		util.ERROR(w, http.StatusBadRequest, errors.New("invalid region_segmentation_model, see /forsete-atr/v1/models/region-segmentation-models/ for valid models"))
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

	yamlPath, err := pipeline.NewTipNotePipeline(regionModel, lineModel, textModel, config.GetConfig().DEVICE).Encode("tmp/yaml", "tipnote.yaml")
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

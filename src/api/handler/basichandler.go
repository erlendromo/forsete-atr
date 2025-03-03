package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/erlendromo/forsete-atr/src/domain/pipeline"
	"github.com/erlendromo/forsete-atr/src/util"
)

type Basic struct {
	LineSegmentationModel string `json:"line_segmentation_model"`
	TextRecognitionModel  string `json:"text_recognition_model"`
}

func (b *Basic) ToPipeline() (pipeline.Pipeline, error) {
	if b.LineSegmentationModel == "" || b.TextRecognitionModel == "" {
		return nil, errors.New("Invalid request, line_segmentation_model or text_recognition_model is not present.")
	}

	return pipeline.NewBasicPipeline(b.LineSegmentationModel, b.TextRecognitionModel), nil
}

func GetBasic(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var basic Basic
	if err := json.NewDecoder(r.Body).Decode(&basic); err != nil {
		util.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	pipeline, err := basic.ToPipeline()
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, err)
		return
	}

	yamlPath, err := pipeline.Encode("/tmp/yaml", "basic.yaml")
	if err != nil {
		util.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	imagePath := "test.png"

	log.Println("Executing htrflow...")
	cmd, err := exec.Command("/bin/bash", fmt.Sprintf("scripts/htrflow.sh %s %s", yamlPath, imagePath)).Output()
	if err != nil {
		log.Println(string(cmd))
		util.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	output, err := os.Open(fmt.Sprintf("/outputs/%s.json", strings.Split(imagePath, ".")[0]))
	if err != nil {
		util.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	util.JSON(w, http.StatusOK, output)
}

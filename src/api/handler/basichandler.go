package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/erlendromo/forsete-atr/src/domain/model"
	"github.com/erlendromo/forsete-atr/src/domain/pipeline"
	"github.com/erlendromo/forsete-atr/src/util"
)

func GetBasic(w http.ResponseWriter, r *http.Request) {
	imageFile, imageHeader, err := r.FormFile("image")
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, err)
		return
	}

	imagePath, err := processImage(imageFile, imageHeader)
	if err != nil {
		util.ERROR(w, http.StatusInternalServerError, err)
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

	yamlPath, err := pipeline.NewBasicPipeline(lineModel, textModel).Encode("tmp/yaml", "basic.yaml")
	if err != nil {
		util.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	cmd := exec.Command("/bin/bash", "scripts/htrflow.sh", yamlPath, imagePath)
	output, err := cmd.Output()
	if err != nil {
		util.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	log.Println(string(output))

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

	util.JSON(w, http.StatusOK, out)
}

// TODO Move and improve this
func processImage(imageFile multipart.File, imageHeader *multipart.FileHeader) (string, error) {
	localImage, err := os.Create("tmp/images/" + imageHeader.Filename)
	if err != nil {
		return "", err
	}

	if _, err := io.Copy(localImage, imageFile); err != nil && err != io.EOF {
		return "", err
	}

	return localImage.Name(), nil
}

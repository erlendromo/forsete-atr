package handler

import (
	"encoding/json"
	"net/http"

	"github.com/erlendromo/forsete-atr/src/util"
)

type Tipnote struct {
	RegionSegmentationModel string `json:"region_segmentation_model"`
	LineSegmentationModel   string `json:"line_segmentation_model"`
	TextRecognitionModel    string `json:"text_recognition_model"`
}

func PostTipnoteDocument(w http.ResponseWriter, r *http.Request) {
	var tipnote Tipnote
	if err := json.NewDecoder(r.Body).Decode(&tipnote); err != nil {
		util.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
}

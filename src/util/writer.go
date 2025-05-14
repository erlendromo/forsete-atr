package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func setHeaders(w http.ResponseWriter, statuscode int, contentType string) {
	w.Header().Set(CONTENT_TYPE, contentType)
	w.WriteHeader(statuscode)
}

func EncodeJSON(w http.ResponseWriter, statuscode int, value any) {
	setHeaders(w, statuscode, APPLICATION_JSON)
	if statuscode == http.StatusNoContent || value == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(value); err != nil {
		ERROR(w, http.StatusInternalServerError, fmt.Errorf(INTERNAL_SERVER_ERROR))
	}
}

func EncodeImage(w http.ResponseWriter, statuscode int, image *os.File) {
	setHeaders(w, statuscode, IMAGE_PNG)
	if _, err := io.Copy(w, image); err != nil {
		ERROR(w, http.StatusInternalServerError, fmt.Errorf(INTERNAL_SERVER_ERROR))
	}
}

package util

import (
	"encoding/json"
	"net/http"
)

func setHeaders(w http.ResponseWriter, statuscode int) {
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.WriteHeader(statuscode)
}

func JSON(w http.ResponseWriter, statuscode int, value any) {
	setHeaders(w, statuscode)
	if statuscode == http.StatusNoContent || value == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(value); err != nil {
		ERROR(w, http.StatusInternalServerError, err)
	}
}

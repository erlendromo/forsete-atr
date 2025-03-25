package handler

import (
	"net/http"

	"github.com/erlendromo/forsete-atr/src/util"
)

// Status
//
// @Summary Status
// @Description Json-response for Status
type Status struct {
	Version string `json:"version"`
	Uptime  string `json:"uptime"`
}

// HeadStatus
//
// @Summary HeadStatus
// @Description Retrieve status of service
// @Tags Status
// @Success 204
// @Router /forsete-atr/v1/status/ [head]
func HeadStatus(w http.ResponseWriter, r *http.Request) {
	util.JSON(w, http.StatusNoContent, nil)
}

// GetStatus
//
// @Summary GetStatus
// @Description Retrieve status of service
// @Tags Status
// @Produce json
// @Success 200 {object} Status
// @Router /forsete-atr/v1/status/ [get]
func GetStatus(w http.ResponseWriter, r *http.Request) {
	util.JSON(w, http.StatusOK, &Status{
		Version: util.VERSION,
		Uptime:  util.UpTimeInHHMMSS(),
	})
}

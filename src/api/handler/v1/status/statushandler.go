package status

import (
	"errors"
	"net/http"
	"os/exec"

	"github.com/erlendromo/forsete-atr/src/util"
)

// Status
//
//	@Summary		Status
//	@Description	Json-response for Status
type Status struct {
	ATR     string `json:"atr"`
	Version string `json:"version"`
	Uptime  string `json:"uptime"`
}

// HeadStatus
//
//	@Summary		HeadStatus
//	@Description	Retrieve status of service
//	@Tags			Status
//	@Success		204
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v1/status/ [head]
func HeadStatus(w http.ResponseWriter, r *http.Request) {
	if err := exec.Command("/bin/bash", "-c", "source /htrflow/venv/bin/activate && htrflow pipeline --help").Run(); err != nil {
		util.NewInternalErrorLog("HTRFLOW STATUS ERROR", err).PrintLog("SERVER ERROR")
		util.ERROR(w, http.StatusInternalServerError, errors.New(util.INTERNAL_SERVER_ERROR))
	}

	util.EncodeJSON(w, http.StatusNoContent, nil)
}

// GetStatus
//
//	@Summary		GetStatus
//	@Description	Retrieve status of service
//	@Tags			Status
//	@Produce		json
//	@Success		200	{object}	Status
//	@Router			/forsete-atr/v1/status/ [get]
func GetStatus(w http.ResponseWriter, r *http.Request) {
	atr := "ready"
	if err := exec.Command("/bin/bash", "-c", "source /htrflow/venv/bin/activate && htrflow pipeline --help").Run(); err != nil {
		util.NewInternalErrorLog("HTRFLOW STATUS ERROR", err).PrintLog("SERVER ERROR")
		atr = "unavailable"
	}

	util.EncodeJSON(w, http.StatusOK, &Status{
		ATR:     atr,
		Version: util.VERSION,
		Uptime:  util.UpTimeInHHMMSS(),
	})
}

package status

import (
	"errors"
	"fmt"
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
		fmt.Printf("\n%sHTRFLOW ERROR%s\n%s\n", util.RED, util.RESET, err.Error())
		util.ERROR(w, http.StatusInternalServerError, errors.New(util.INTERNAL_SERVER_ERROR))
	}

	util.JSON(w, http.StatusNoContent, nil)
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
		fmt.Printf("\n%sHTRFLOW ERROR%s\n%s\n", util.RED, util.RESET, err.Error())
		atr = "unavailable"
	}

	util.JSON(w, http.StatusOK, &Status{
		ATR:     atr,
		Version: util.VERSION,
		Uptime:  util.UpTimeInHHMMSS(),
	})
}

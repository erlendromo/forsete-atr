package status

import (
	"errors"
	"net/http"
	"os/exec"

	"github.com/erlendromo/forsete-atr/src/util"
)

type Status struct {
	ATR     string `json:"atr"`
	Version string `json:"version"`
	Uptime  string `json:"uptime"`
}

func HeadStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := exec.Command("/bin/bash", "-c", "source /htrflow/venv/bin/activate && htrflow pipeline --help").Run(); err != nil {
			util.NewInternalErrorLog("HTRFLOW STATUS ERROR", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, errors.New(util.INTERNAL_SERVER_ERROR))
		}

		util.EncodeJSON(w, http.StatusNoContent, nil)
	}
}

func GetStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}

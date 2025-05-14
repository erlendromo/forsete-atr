package status

import (
	"errors"
	"net/http"
	"os/exec"

	"github.com/erlendromo/forsete-atr/src/util"
	"github.com/jmoiron/sqlx"
)

// Status
//
//	@Summary		Status
//	@Description	Status containing ATR-readiness, database-readiness, version and uptime.
type Status struct {
	ATR      string `json:"atr"`
	Database string `json:"database"`
	Version  string `json:"version"`
	Uptime   string `json:"uptime"`
}

// HeadStatus
//
//	@Summary		Head status
//	@Description	Head status.
//	@Tags			Status
//	@Success		204
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v2/status/ [head]
func HeadStatus(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := exec.Command("/bin/bash", "-c", "source /htrflow/venv/bin/activate && htrflow pipeline --help").Run(); err != nil {
			util.NewInternalErrorLog("HTRFLOW STATUS ERROR", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, errors.New(util.INTERNAL_SERVER_ERROR))
		}

		if err := db.Ping(); err != nil {
			util.NewInternalErrorLog("DATABASE STATUS ERROR", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, errors.New(util.INTERNAL_SERVER_ERROR))
		}

		util.EncodeJSON(w, http.StatusNoContent, nil)
	}
}

// GetStatus
//
//	@Summary		Get status
//	@Description	Get status.
//	@Tags			Status
//	@Produce		json
//	@Success		200	{object}	Status
//	@Router			/forsete-atr/v2/status/ [get]
func GetStatus(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		atr := "ready"
		if err := exec.Command("/bin/bash", "-c", "source /htrflow/venv/bin/activate && htrflow pipeline --help").Run(); err != nil {
			util.NewInternalErrorLog("HTRFLOW STATUS ERROR", err).PrintLog("SERVER ERROR")
			atr = "unavailable, service restart needed"
		}

		dbStatus := "ready"
		if err := db.Ping(); err != nil {
			dbStatus = "unavailable, service restart needed"
		}

		util.EncodeJSON(w, http.StatusOK, &Status{
			ATR:      atr,
			Database: dbStatus,
			Version:  util.VERSION,
			Uptime:   util.UpTimeInHHMMSS(),
		})
	}
}

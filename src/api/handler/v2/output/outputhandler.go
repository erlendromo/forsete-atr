package output

import (
	"fmt"
	"net/http"

	"github.com/erlendromo/forsete-atr/src/api/middleware"
	"github.com/erlendromo/forsete-atr/src/business/domain/output"
	atrservice "github.com/erlendromo/forsete-atr/src/business/usecase/service/atr_service"
	"github.com/erlendromo/forsete-atr/src/util"
	"github.com/google/uuid"
)

// GetOutputsByImageID
//
//	@Summary		Get outputs by image id
//	@Description	Get outputs by image id.
//	@Tags			Outputs
//	@Param			imageID			query	string	true	"uuid of image"
//	@Param			Authorization	header	string	true	"'Bearer token' must be set for valid response"
//	@Produce		json
//	@Success		200	{object}	[]output.Output
//	@Failure		401	{object}	util.ErrorResponse
//	@Failure		422	{object}	util.ErrorResponse
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v2/images/{imageID}/outputs/ [get]
func GetOutputsByImageID(atrService *atrservice.ATRService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxValues, ok := r.Context().Value(middleware.ContextValuesKey).(*middleware.ContextValues)
		if !ok {
			err := fmt.Errorf("missing 'context_values' in request-context")
			util.NewInternalErrorLog("OUTPUTS BY IMAGE ID (CONTEXT-VALUES)", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		imageID, err := uuid.Parse(r.PathValue("imageID"))
		if err != nil {
			util.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("unable to parse imageID"))
			return
		}

		outputs, err := atrService.OutputRepo.OutputsByImageID(r.Context(), imageID, ctxValues.User.ID)
		if err != nil {
			util.NewInternalErrorLog("OUTPUTS BY IMAGE ID", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		util.EncodeJSON(w, http.StatusOK, outputs)
	}
}

// GetOutputByID
//
//	@Summary		Get output by id
//	@Description	Get output by id.
//	@Tags			Outputs
//	@Param			imageID			query	string	true	"uuid of image"
//	@Param			outputID		query	string	true	"uuid of output"
//	@Param			Authorization	header	string	true	"'Bearer token' must be set for valid response"
//	@Produce		json
//	@Success		200	{object}	output.Output
//	@Failure		401	{object}	util.ErrorResponse
//	@Failure		422	{object}	util.ErrorResponse
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v2/images/{imageID}/outputs/{outputID}/ [get]
func GetOutputByID(atrService *atrservice.ATRService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxValues, ok := r.Context().Value(middleware.ContextValuesKey).(*middleware.ContextValues)
		if !ok {
			err := fmt.Errorf("missing 'context_values' in request-context")
			util.NewInternalErrorLog("OUTPUT BY ID (CONTEXT-VALUES)", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		imageID, err := uuid.Parse(r.PathValue("imageID"))
		if err != nil {
			util.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("unable to parse imageID"))
			return
		}

		outputID, err := uuid.Parse(r.PathValue("outputID"))
		if err != nil {
			util.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("unable to parse outputID"))
			return
		}

		output, err := atrService.OutputRepo.OutputByID(r.Context(), outputID, imageID, ctxValues.User.ID)
		if err != nil {
			util.NewInternalErrorLog("OUTPUT BY ID", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		util.EncodeJSON(w, http.StatusOK, output)
	}
}

// UpdateOutputForm
//
//	@Summary		UpdateOutputForm
//	@Description	Form containing confirmed and data associated with the update request.
type UpdateOutputForm struct {
	Confirmed bool               `json:"confirmed"`
	Data      output.ATRResponse `json:"data"`
}

// UpdateOutputByID
//
//	@Summary		Update output by id
//	@Description	Update output by id.
//	@Tags			Outputs
//	@Param			imageID			query	string				true	"uuid of image"
//	@Param			outputID		query	string				true	"uuid of output"
//	@Param			Authorization	header	string				true	"'Bearer token' must be set for valid response"
//	@Param			request			body	UpdateOutputForm	true	"Body containing confirmed and data to update"
//	@Produce		json
//	@Success		200	{object}	output.Output
//	@Failure		401	{object}	util.ErrorResponse
//	@Failure		422	{object}	util.ErrorResponse
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v2/images/{imageID}/outputs/{outputID}/ [put]
func UpdateOutputByID(atrService *atrservice.ATRService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxValues, ok := r.Context().Value(middleware.ContextValuesKey).(*middleware.ContextValues)
		if !ok {
			err := fmt.Errorf("missing 'context_values' in request-context")
			util.NewInternalErrorLog("UPDATE OUTPUT BY ID (CONTEXT-VALUES)", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		imageID, err := uuid.Parse(r.PathValue("imageID"))
		if err != nil {
			util.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("unable to parse imageID"))
			return
		}

		outputID, err := uuid.Parse(r.PathValue("outputID"))
		if err != nil {
			util.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("unable to parse outputID"))
			return
		}

		form, err := util.DecodeJSON[UpdateOutputForm](r.Body)
		if err != nil {
			util.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("unable to parse request body"))
			return
		}

		output, err := atrService.OutputRepo.UpdateOutputByID(r.Context(), outputID, imageID, ctxValues.User.ID, form.Confirmed)
		if err != nil {
			util.NewInternalErrorLog("UPDATE OUTPUT BY ID", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		if err := output.CreateLocal(&form.Data); err != nil {
			util.NewInternalErrorLog("UPDATE OUTPUT BY ID", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		util.EncodeJSON(w, http.StatusOK, output)
	}
}

// GetOutputData
//
//	@Summary		Get output data
//	@Description	Get output data.
//	@Tags			Outputs
//	@Param			imageID			query	string	true	"uuid of image"
//	@Param			outputID		query	string	true	"uuid of output"
//	@Param			Authorization	header	string	true	"'Bearer token' must be set for valid response"
//	@Produce		json
//	@Success		200	{object}	output.ATRResponse
//	@Failure		401	{object}	util.ErrorResponse
//	@Failure		422	{object}	util.ErrorResponse
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v2/images/{imageID}/outputs/{outputID}/data/ [get]
func GetOutputData(atrService *atrservice.ATRService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxValues, ok := r.Context().Value(middleware.ContextValuesKey).(*middleware.ContextValues)
		if !ok {
			err := fmt.Errorf("missing 'context_values' in request-context")
			util.NewInternalErrorLog("OUTPUT DATA (CONTEXT-VALUES)", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		imageID, err := uuid.Parse(r.PathValue("imageID"))
		if err != nil {
			util.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("unable to parse imageID"))
			return
		}

		outputID, err := uuid.Parse(r.PathValue("outputID"))
		if err != nil {
			util.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("unable to parse outputID"))
			return
		}

		output, err := atrService.OutputRepo.OutputByID(r.Context(), outputID, imageID, ctxValues.User.ID)
		if err != nil {
			util.NewInternalErrorLog("OUTPUT DATA", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		atrResponse, err := output.ReadJson(fmt.Sprintf("%s/%s.%s", output.Path, output.ID, output.Format))
		if err != nil {
			util.NewInternalErrorLog("OUTPUT DATA", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		util.EncodeJSON(w, http.StatusOK, atrResponse)
	}
}

// DeleteOutputByID
//
//	@Summary		Delete output by id
//	@Description	Delete output by id.
//	@Tags			Outputs
//	@Param			imageID			query	string	true	"uuid of image"
//	@Param			outputID		query	string	true	"uuid of output"
//	@Param			Authorization	header	string	true	"'Bearer token' must be set for valid response"
//	@Produce		json
//	@Success		204
//	@Failure		401	{object}	util.ErrorResponse
//	@Failure		422	{object}	util.ErrorResponse
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v2/images/{imageID}/outputs/{outputID}/ [delete]
func DeleteOutputByID(atrService *atrservice.ATRService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxValues, ok := r.Context().Value(middleware.ContextValuesKey).(*middleware.ContextValues)
		if !ok {
			err := fmt.Errorf("missing 'context_values' in request-context")
			util.NewInternalErrorLog("DELETE OUTPUT BY ID (CONTEXT-VALUES)", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		imageID, err := uuid.Parse(r.PathValue("imageID"))
		if err != nil {
			util.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("unable to parse imageID"))
			return
		}

		outputID, err := uuid.Parse(r.PathValue("outputID"))
		if err != nil {
			util.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("unable to parse outputID"))
			return
		}

		if err := atrService.OutputRepo.DeleteOutputByID(r.Context(), outputID, imageID, ctxValues.User.ID); err != nil {
			util.NewInternalErrorLog("DELETE OUTPUT BY ID (DELETE)", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		util.EncodeJSON(w, http.StatusNoContent, nil)
	}
}

package atr

import (
	"fmt"
	"net/http"

	"github.com/erlendromo/forsete-atr/src/api/middleware"
	"github.com/erlendromo/forsete-atr/src/business/domain/output"
	"github.com/erlendromo/forsete-atr/src/business/domain/pipeline"
	atrservice "github.com/erlendromo/forsete-atr/src/business/usecase/service/atr_service"
	fileservice "github.com/erlendromo/forsete-atr/src/business/usecase/service/file_service"
	"github.com/erlendromo/forsete-atr/src/util"
	"github.com/google/uuid"
)

// ATRRequest
//
//	@Summary		ATRRequest
//	@Description	Body containing line_segmentation_model, text_recognition_model and image_ids.
type ATRRequest struct {
	LineModelName string   `json:"line_segmentation_model"`
	TextModelName string   `json:"text_recognition_model"`
	ImageIDs      []string `json:"image_ids"`
}

// Run
//
//	@Summary		Run ATR
//	@Description	Run ATR on images
//	@Tags			ATR
//	@Accept			json
//	@Param			Authorization	header	string		true	"'Bearer <token>' must be set for valid response"
//	@Param			request			body	ATRRequest	true	"Body containing which models to use, alongside the image_ids"
//	@Body			ATRRequest 																																																							{object} 	json 	true	"request-form"
//	@Produce		json
//	@Success		200	{object}	[]output.Output
//	@Failure		400	{object}	util.ErrorResponse
//	@Failure		401	{object}	util.ErrorResponse
//	@Failure		422	{object}	util.ErrorResponse
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v2/atr/ [post]
func Run(fileService *fileservice.FileService, atrService *atrservice.ATRService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contextValues, ok := r.Context().Value(middleware.ContextValuesKey).(*middleware.ContextValues)
		if !ok {
			err := fmt.Errorf("missing 'context_values' in request-context")
			util.NewInternalErrorLog("RUN ATR (CONTEXT-VALUES)", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		atrRequest, err := util.DecodeJSON[ATRRequest](r.Body)
		if err != nil {
			util.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		var pipeline *pipeline.Pipeline
		if atrRequest.LineModelName == "" {
			p, err := atrService.PipelineRepo.PipelineByModel(r.Context(), atrRequest.TextModelName)
			if err != nil {
				util.NewInternalErrorLog("RUN ATR (PIPELINE)", err).PrintLog("SERVER ERROR")
				util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
				return
			}

			pipeline = p
		} else {
			p, err := atrService.PipelineRepo.PipelineByModels(r.Context(), atrRequest.LineModelName, atrRequest.TextModelName)
			if err != nil {
				util.NewInternalErrorLog("RUN ATR (PIPELINE)", err).PrintLog("SERVER ERROR")
				util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
				return
			}

			pipeline = p
		}

		outputs := make([]*output.Output, 0)
		for _, imageID := range atrRequest.ImageIDs {
			parsedImageID, err := uuid.Parse(imageID)
			if err != nil {
				util.ERROR(w, http.StatusBadRequest, err)
				return
			}

			image, err := fileService.ImageRepo.ImageByID(r.Context(), parsedImageID, contextValues.User.ID)
			if err != nil {
				util.NewInternalErrorLog("RUN ATR (IMAGE)", err).PrintLog("SERVER ERROR")
				util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
				return
			}

			output, err := atrService.RunATROnImage(r.Context(), image, pipeline, contextValues.User.ID)
			if err != nil {
				util.NewInternalErrorLog("RUN ATR (OUTPUT)", err).PrintLog("SERVER ERROR")
				util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
				return
			}

			outputs = append(outputs, output)
		}

		util.EncodeJSON(w, http.StatusOK, outputs)
	}
}

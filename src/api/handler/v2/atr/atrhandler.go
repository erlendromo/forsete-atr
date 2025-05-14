package atr

import (
	"fmt"
	"net/http"

	"github.com/erlendromo/forsete-atr/src/api/middleware"
	atrservice "github.com/erlendromo/forsete-atr/src/business/usecase/service/atr_service"
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

// Parse image-ids to uuid.UUID type
func (ar *ATRRequest) parseImageIDs() ([]uuid.UUID, error) {
	parsedImageIDs := make([]uuid.UUID, 0)
	for _, imageID := range ar.ImageIDs {
		parsedImageID, err := uuid.Parse(imageID)
		if err != nil {
			return nil, err
		}

		parsedImageIDs = append(parsedImageIDs, parsedImageID)
	}

	return parsedImageIDs, nil
}

// Run
//
//	@Summary		Run ATR
//	@Description	Run ATR on images
//	@Tags			ATR
//	@Accept			json
//	@Param			Authorization	header	string		true	"'Bearer token' must be set for valid response"
//	@Param			request			body	ATRRequest	true	"Body containing which models to use, alongside the image_ids"
//	@Body			ATRRequest 																																																																																					{object} 	json 	true	"request-form"
//	@Produce		json
//	@Success		200	{object}	[]output.Output
//	@Failure		400	{object}	util.ErrorResponse
//	@Failure		401	{object}	util.ErrorResponse
//	@Failure		422	{object}	util.ErrorResponse
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v2/atr/ [post]
func Run(atrService *atrservice.ATRService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxValues, ok := r.Context().Value(middleware.ContextValuesKey).(*middleware.ContextValues)
		if !ok {
			err := fmt.Errorf("missing 'context_values' in request-context")
			util.NewInternalErrorLog("RUN ATR (CONTEXT-VALUES)", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		atrRequest, err := util.DecodeJSON[ATRRequest](r.Body)
		if err != nil {
			util.NewInternalErrorLog("RUN ATR (DECODE BODY)", err).PrintLog("CLIENT ERROR")
			util.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("unable to parse response-body"))
			return
		}

		parsedImageIDs, err := atrRequest.parseImageIDs()
		if err != nil {
			util.NewInternalErrorLog("RUN ATR (PARSE IMAGE_IDS)", err).PrintLog("CLIENT ERROR")
			util.ERROR(w, http.StatusBadRequest, fmt.Errorf("unable to parse imageIDs"))
			return
		}

		pipeline, err := atrService.PipelineRepo.PipelineByModels(r.Context(), atrRequest.LineModelName, atrRequest.TextModelName)
		if err != nil {
			util.NewInternalErrorLog("RUN ATR (PIPELINE)", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		outputs, err := atrService.RunATROnImages(r.Context(), pipeline.ID, ctxValues.User.ID, parsedImageIDs)
		if err != nil {
			util.NewInternalErrorLog("RUN ATR (OUTPUTS)", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		util.EncodeJSON(w, http.StatusOK, outputs)
	}
}

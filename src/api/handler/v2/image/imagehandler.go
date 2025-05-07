package image

import (
	"fmt"
	"net/http"
	"os"

	"github.com/erlendromo/forsete-atr/src/api/middleware"
	_ "github.com/erlendromo/forsete-atr/src/business/domain/image"
	atrservice "github.com/erlendromo/forsete-atr/src/business/usecase/service/atr_service"
	"github.com/erlendromo/forsete-atr/src/util"
	"github.com/google/uuid"
)

// UploadImages
//
//	@Summary		Upload images
//	@Description	Upload up to 32MB worth of images.
//	@Tags			Images
//	@Param			Authorization	header		string	true		"'Bearer token' must be set for valid response"
//	@Param			images			formData	file	required	"images to upload"
//	@Produce		json
//	@Success		200	{object}	image.Image
//	@Failure		400	{object}	util.ErrorResponse
//	@Failure		401	{object}	util.ErrorResponse
//	@Failure		422	{object}	util.ErrorResponse
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v2/images/upload/ [post]
func UploadImages(atrService *atrservice.ATRService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxValues, ok := r.Context().Value(middleware.ContextValuesKey).(*middleware.ContextValues)
		if !ok {
			err := fmt.Errorf("missing 'context_values' in request-context")
			util.NewInternalErrorLog("UPLOAD IMAGES", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		// 32 MB
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			util.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("failed to parse multipart form: %w", err))
			return
		}

		fileHeaders, ok := r.MultipartForm.File["images"]
		if !ok || len(fileHeaders) == 0 {
			util.ERROR(w, http.StatusBadRequest, fmt.Errorf("missing images in multipart form"))
			return
		}

		images, err := atrService.UploadImages(r.Context(), ctxValues.User.ID, fileHeaders)
		if err != nil {
			util.NewInternalErrorLog("UPLOAD IMAGES", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		util.EncodeJSON(w, http.StatusOK, images)
	}
}

// GetImages
//
//	@Summary		Get images
//	@Description	Get all images the user has uploaded.
//	@Tags			Images
//	@Param			Authorization	header	string	true	"'Bearer token' must be set for valid response"
//	@Produce		json
//	@Success		200	{object}	[]image.Image
//	@Failure		401	{object}	util.ErrorResponse
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v2/images/ [get]
func GetImages(atrService *atrservice.ATRService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxValues, ok := r.Context().Value(middleware.ContextValuesKey).(*middleware.ContextValues)
		if !ok {
			err := fmt.Errorf("missing 'context_values' in request-context")
			util.NewInternalErrorLog("GET IMAGES", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		images, err := atrService.ImageRepo.ImagesByUserID(r.Context(), ctxValues.User.ID)
		if err != nil {
			util.NewInternalErrorLog("GET IMAGES", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		util.EncodeJSON(w, http.StatusOK, images)
	}
}

// GetImageByID
//
//	@Summary		Get image by id
//	@Description	Get image by id.
//	@Tags			Images
//	@Param			imageID			query	string	true	"uuid of image"
//	@Param			Authorization	header	string	true	"'Bearer token' must be set for valid response"
//	@Produce		json
//	@Success		200	{object}	image.Image
//	@Failure		401	{object}	util.ErrorResponse
//	@Failure		422	{object}	util.ErrorResponse
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v2/images/{imageID}/ [get]
func GetImageByID(atrService *atrservice.ATRService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxValues, ok := r.Context().Value(middleware.ContextValuesKey).(*middleware.ContextValues)
		if !ok {
			err := fmt.Errorf("missing 'context_values' in request-context")
			util.NewInternalErrorLog("GET IMAGE BY ID", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		id, err := uuid.Parse(r.PathValue("imageID"))
		if err != nil {
			util.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("invalid imageID"))
			return
		}

		image, err := atrService.ImageRepo.ImageByID(r.Context(), id, ctxValues.User.ID)
		if err != nil {
			util.NewInternalErrorLog("GET IMAGE BY ID", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		util.EncodeJSON(w, http.StatusOK, image)
	}
}

// GetImageByID
//
//	@Summary		Get image data
//	@Description	Get image data.
//	@Tags			Images
//	@Param			imageID			query	string	true	"uuid of image"
//	@Param			Authorization	header	string	true	"'Bearer token' must be set for valid response"
//	@Produce		json
//	@Success		200	body		file	"image file"
//	@Failure		401	{object}	util.ErrorResponse
//	@Failure		422	{object}	util.ErrorResponse
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v2/images/{imageID}/data/ [get]
func GetImageData(atrService *atrservice.ATRService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxValues, ok := r.Context().Value(middleware.ContextValuesKey).(*middleware.ContextValues)
		if !ok {
			err := fmt.Errorf("missing 'context_values' in request-context")
			util.NewInternalErrorLog("GET IMAGE DATA", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		imageID, err := uuid.Parse(r.PathValue("imageID"))
		if err != nil {
			util.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("invalid imageID"))
			return
		}

		image, err := atrService.ImageRepo.ImageByID(r.Context(), imageID, ctxValues.User.ID)
		if err != nil {
			util.NewInternalErrorLog("GET IMAGE DATA", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		imagePath := fmt.Sprintf("%s/%s.%s", image.Path, image.ID.String(), image.Format)

		file, err := os.Open(imagePath)
		if err != nil {
			util.NewInternalErrorLog("GET IMAGE DATA", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		defer file.Close()

		util.EncodeImage(w, http.StatusOK, file)
	}
}

// DeleteImageByID
//
//	@Summary		Delete image by id
//	@Description	Delete image (and corresponding output data) by imageID.
//	@Tags			Images
//	@Param			imageID			query	string	true	"uuid of image"
//	@Param			Authorization	header	string	true	"'Bearer token' must be set for valid response"
//	@Produce		json
//	@Success		204
//	@Failure		401	{object}	util.ErrorResponse
//	@Failure		422	{object}	util.ErrorResponse
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v2/images/{imageID}/ [delete]
func DeleteImageByID(atrService *atrservice.ATRService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxValues, ok := r.Context().Value(middleware.ContextValuesKey).(*middleware.ContextValues)
		if !ok {
			err := fmt.Errorf("missing 'context_values' in request-context")
			util.NewInternalErrorLog("DELETE IMAGE BY ID", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		imageID, err := uuid.Parse(r.PathValue("imageID"))
		if err != nil {
			util.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("invalid imageID"))
			return
		}

		if err := atrService.DeleteImageAndOutputs(r.Context(), imageID, ctxValues.User.ID); err != nil {
			err := fmt.Errorf("missing 'context_values' in request-context")
			util.NewInternalErrorLog("DELETE IMAGE BY ID", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		util.EncodeJSON(w, http.StatusNoContent, nil)
	}
}

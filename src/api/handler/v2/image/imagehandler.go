package image

import (
	"fmt"
	"net/http"

	"github.com/erlendromo/forsete-atr/src/api/middleware"
	imagerepository "github.com/erlendromo/forsete-atr/src/business/usecase/repository/image_repository"
	"github.com/erlendromo/forsete-atr/src/util"
)

func UploadImages(imageRepo *imagerepository.ImageRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxValues, ok := r.Context().Value(middleware.ContextValuesKey).(*middleware.ContextValues)
		if !ok {
			util.ERROR(w, http.StatusUnauthorized, fmt.Errorf("invalid user, login to use this endpoint"))
			return
		}

		// 32 MB
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			util.ERROR(w, http.StatusBadRequest, fmt.Errorf("failed to parse multipart form: %w", err))
			return
		}

		fileHeaders, ok := r.MultipartForm.File["images"]
		if !ok || len(fileHeaders) == 0 {
			util.ERROR(w, http.StatusBadRequest, fmt.Errorf("missing images in multipart form"))
			return
		}

		images, err := imageRepo.UploadImages(r.Context(), fileHeaders, ctxValues.User.ID)
		if err != nil {
			util.NewInternalErrorLog("UPLOAD IMAGES", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		util.EncodeJSON(w, http.StatusOK, images)
	}
}

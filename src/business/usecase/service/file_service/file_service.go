package fileservice

import (
	"context"
	"fmt"
	"mime/multipart"
	"path"
	"path/filepath"
	"strings"

	"github.com/erlendromo/forsete-atr/src/business/domain/image"
	imagerepository "github.com/erlendromo/forsete-atr/src/business/usecase/repository/image_repository"
	outputrepository "github.com/erlendromo/forsete-atr/src/business/usecase/repository/output_repository"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type FileService struct {
	ImageRepo  *imagerepository.ImageRepository
	OutputRepo *outputrepository.OutputRepository
}

func NewFileService(db *sqlx.DB) *FileService {
	return &FileService{
		ImageRepo:  imagerepository.NewImageRepository(db),
		OutputRepo: outputrepository.NewOutputRepository(db),
	}
}

func (f *FileService) UploadImages(ctx context.Context, userID uuid.UUID, fileHeaders []*multipart.FileHeader) ([]*image.Image, error) {
	images := make([]*image.Image, 0)
	errs := make([]error, 0)

	for _, fileHeader := range fileHeaders {
		img, err := f.uploadImage(ctx, userID, fileHeader)
		if err != nil {
			errs = append(errs, err)
		}

		images = append(images, img)
	}

	if len(errs) > 0 {
		return nil, fmt.Errorf("unable to upload images: %+v", errs)
	}

	return images, nil
}

func (f *FileService) uploadImage(ctx context.Context, userID uuid.UUID, fileHeader *multipart.FileHeader) (*image.Image, error) {
	ext := filepath.Ext(fileHeader.Filename)
	originalName := strings.TrimSuffix(fileHeader.Filename, ext)

	name := strings.ToLower(originalName)
	format := strings.ToLower(strings.TrimPrefix(ext, "."))
	path := path.Join("assets", "users", userID.String(), "images")

	img, err := f.ImageRepo.RegisterImage(ctx, name, format, path, userID)
	if err != nil {
		return nil, err
	}

	if err := img.CreateLocal(fileHeader); err != nil {
		if _, err := f.ImageRepo.DeleteImageByID(ctx, img.ID); err != nil {
			return nil, err
		}

		return nil, err
	}

	return img, nil
}

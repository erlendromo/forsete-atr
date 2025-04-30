package imagerepository

import (
	"context"
	"fmt"
	"mime/multipart"
	"path"
	"path/filepath"
	"strings"

	"github.com/erlendromo/forsete-atr/src/business/domain/image"
	"github.com/erlendromo/forsete-atr/src/database"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ImageRepository struct {
	db *sqlx.DB
}

func NewImageRepository(db *sqlx.DB) *ImageRepository {
	return &ImageRepository{
		db: db,
	}
}

func (i *ImageRepository) ImageByID(ctx context.Context, id, userID uuid.UUID) (*image.Image, error) {
	query := `
		SELECT
			id,
			name,
			format,
			path,
			uploaded_at,
			deleted_at,
			user_id
		FROM
			"image"
		WHERE
			id = $1
		AND
			user_id = $2
		AND
			deleted_at IS NULL
	`

	return database.QueryRowx[image.Image](ctx, i.db, query, id, userID)
}

func (i *ImageRepository) ImagesByUserID(ctx context.Context, userID uuid.UUID) ([]*image.Image, error) {
	query := `
		SELECT
			id,
			name,
			format,
			path,
			uploaded_at,
			deleted_at,
			user_id
		FROM
			"image"
		WHERE
			user_id = $1
		AND
			deleted_at IS NULL
		ORDER BY
			uploaded_at
		ASC
	`

	return database.Queryx[image.Image](ctx, i.db, query, userID)
}

func (i *ImageRepository) UploadImage(ctx context.Context, fileHeader *multipart.FileHeader, userID uuid.UUID) (*image.Image, error) {
	dst := path.Join("assets", "users", userID.String(), "images")
	img := &image.Image{}

	path, err := img.CreateLocal(dst, fileHeader)
	if err != nil {
		return nil, err
	}

	format := strings.ToLower(strings.TrimPrefix(filepath.Ext(fileHeader.Filename), "."))

	query := `
		INSERT INTO
			"image" (name, format, path, user_id)
		VALUES
			($1, $2, $3, $4)
		RETURNING
			id,
			name,
			format
	`

	return database.QueryRowx[image.Image](ctx, i.db, query, fileHeader.Filename, format, path, userID)
}

func (i *ImageRepository) UploadImages(ctx context.Context, fileHeaders []*multipart.FileHeader, userID uuid.UUID) ([]*image.Image, error) {
	images := make([]*image.Image, 0)
	errs := make([]error, 0)
	for _, fileHeader := range fileHeaders {
		img, err := i.UploadImage(ctx, fileHeader, userID)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		images = append(images, img)
	}

	if len(errs) > 0 {
		return nil, fmt.Errorf("error uploading images: %+v", errs)
	}

	return images, nil
}

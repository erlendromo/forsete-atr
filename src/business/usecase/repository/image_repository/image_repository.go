package imagerepository

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/image"
	"github.com/erlendromo/forsete-atr/src/database"
	"github.com/erlendromo/forsete-atr/src/querier"
	"github.com/erlendromo/forsete-atr/src/querier/sqlx"
	"github.com/google/uuid"
)

type ImageRepository struct {
	querier querier.Querier[image.Image]
}

func NewImageRepository(db database.Database) *ImageRepository {
	return &ImageRepository{
		querier: sqlx.NewSqlxQuerier[image.Image](db),
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

	return i.querier.QueryRowx(ctx, query, id, userID)
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

	return i.querier.Queryx(ctx, query, userID)
}

func (i *ImageRepository) RegisterImage(ctx context.Context, name, format, path string, userID uuid.UUID) (*image.Image, error) {
	query := `
		INSERT INTO
			"image" (name, format, path, user_id)
		VALUES
			($1, $2, $3, $4)
		RETURNING
			id,
			name,
			format,
			path,
			uploaded_at,
			deleted_at,
			user_id
	`

	return i.querier.QueryRowx(ctx, query, name, format, path, userID)
}

func (i *ImageRepository) DeleteImageByID(ctx context.Context, id, userID uuid.UUID) error {
	query := `
		UPDATE
			"image"
		SET
			deleted_at = now()
		WHERE
			id = $1
		AND
			user_id = $2
		AND
			deleted_at IS NULL
	`

	return i.querier.Executex(ctx, query, id, userID)
}

// Unused for now
func (i *ImageRepository) UpdateNameByID(ctx context.Context, name string, imageID, userID uuid.UUID) error {
	query := `
		UPDATE
			"image"
		SET
			name = $1
		WHERE
			id = $2
		AND
			user_id = $3
		AND
			deleted_at IS NULL
	`

	return i.querier.Executex(ctx, query, name, imageID, userID)
}

func (i *ImageRepository) DeleteUserImages(ctx context.Context, userID uuid.UUID) error {
	query := `
		UPDATE
			"image"
		SET
			deleted_at = now()
		WHERE
			user_id = $1
		AND
			deleted_at IS NULL
	`

	return i.querier.Executex(ctx, query, userID)
}

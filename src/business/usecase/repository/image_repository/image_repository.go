package imagerepository

import (
	"context"

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

	return database.QueryRowx[image.Image](ctx, i.db, query, name, format, path, userID)
}

func (i *ImageRepository) DeleteImageByID(ctx context.Context, id, userID uuid.UUID) (int, error) {
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

	return database.ExecuteContext(ctx, i.db, query, id, userID)
}

// Unused for now
func (i *ImageRepository) UpdateNameByID(ctx context.Context, imageID, userID uuid.UUID, name string) (int, error) {
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

	return database.ExecuteContext(ctx, i.db, query, name, imageID, userID)
}

func (i *ImageRepository) DeleteUserImages(ctx context.Context, userID uuid.UUID) (int, error) {
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

	return database.ExecuteContext(ctx, i.db, query, userID)
}

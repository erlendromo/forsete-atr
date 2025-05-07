package outputrepository

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/output"
	"github.com/erlendromo/forsete-atr/src/database"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type OutputRepository struct {
	db *sqlx.DB
}

func NewOutputRepository(db *sqlx.DB) *OutputRepository {
	return &OutputRepository{
		db: db,
	}
}

func (o *OutputRepository) OutputByID(ctx context.Context, outputID, imageID, userID uuid.UUID) (*output.Output, error) {
	query := `
		SELECT
			o.id,
			o.name,
			o.format,
			o.path,
			o.created_at,
			o.updated_at,
			o.deleted_at,
			o.confirmed,
			o.image_id
		FROM
			"output" o
		JOIN
			"image" i
		ON
			o.image_id = i.id
		WHERE
			o.id = $1
		AND
			o.image_id = $2
		AND
			i.deleted_at IS NULL
		AND
			i.user_id = $3
		AND
			o.deleted_at IS NULL
	`

	return database.QueryRowx[output.Output](ctx, o.db, query, outputID, imageID, userID)
}

func (o *OutputRepository) OutputsByImageID(ctx context.Context, imageID, userID uuid.UUID) ([]*output.Output, error) {
	query := `
		SELECT
			o.id,
			o.name,
			o.format,
			o.path,
			o.created_at,
			o.updated_at,
			o.deleted_at,
			o.confirmed,
			o.image_id
		FROM
			"output" o
		JOIN
			"image" i
		ON
			o.image_id = i.id
		WHERE
			o.image_id = $1
		AND
			i.deleted_at IS NULL
		AND
			i.user_id = $2
		AND
			o.deleted_at IS NULL
	`

	return database.Queryx[output.Output](ctx, o.db, query, imageID, userID)
}

func (o *OutputRepository) RegisterOutput(ctx context.Context, name, format, path string, imageID, userID uuid.UUID) (*output.Output, error) {
	query := `
		INSERT INTO
			"output" (name, format, path, image_id)
		SELECT
			$1,
			$2,
			$3,
			i.id
		FROM
			"image" i
		WHERE
			i.id = $4
		AND
			i.user_id = $5
		AND
			i.deleted_at IS NULL
		RETURNING
			id,
			name,
			format,
			path,
			created_at,
			updated_at,
			deleted_at,
			confirmed,
			image_id
	`

	return database.QueryRowx[output.Output](ctx, o.db, query, name, format, path, imageID, userID)
}

func (o *OutputRepository) UpdateOutputByID(ctx context.Context, outputID, imageID, userID uuid.UUID, confirmed bool) (*output.Output, error) {
	query := `
		UPDATE
			"output" o
		SET
			confirmed = $1,
			updated_at = now()
		FROM
			"image" i
		WHERE
			o.id = $2
		AND
			o.image_id = $3
		AND
			o.deleted_at IS NULL
		AND
			i.id = o.image_id
		AND
			i.user_id = $4
		AND
			i.deleted_at IS NULL
		RETURNING
			o.id,
			o.name,
			o.format,
			o.path,
			o.created_at,
			o.updated_at,
			o.deleted_at,
			o.confirmed,
			o.image_id
	`

	return database.QueryRowx[output.Output](ctx, o.db, query, confirmed, outputID, imageID, userID)
}

func (o *OutputRepository) DeleteOutputByID(ctx context.Context, outputID, imageID, userID uuid.UUID) (int, error) {
	query := `
		UPDATE
			"output" o
		SET
			deleted_at = now()
		FROM
			"image" i
		WHERE
			o.id = $1
		AND
			o.deleted_at IS NULL
		AND
			i.id = o.image_id
		AND
			i.id = $2
		AND
			i.user_id = $3
		AND
			i.deleted_at IS NULL
	`

	return database.ExecuteContext(ctx, o.db, query, outputID, imageID, userID)
}

func (o *OutputRepository) DeleteOutputsByImageID(ctx context.Context, imageID, userID uuid.UUID) (int, error) {
	query := `
		UPDATE
			"output" o
		SET
			deleted_at = now()
		FROM
			"image" i
		WHERE
			o.image_id = $1
		AND
			o.deleted_at IS NULL
		AND
			i.id = o.image_id
		AND
			i.user_id = $2
		AND
			i.deleted_at IS NULL
    `

	return database.ExecuteContext(ctx, o.db, query, imageID, userID)
}

func (o *OutputRepository) DeleteUserOutputs(ctx context.Context, userID uuid.UUID) (int, error) {
	query := `
		UPDATE
			"output"
		SET
			deleted_at = now()
		WHERE
			image_id IN (
				SELECT
					id
				FROM
					"image"
				WHERE
					user_id = $1
				AND
					deleted_at IS NULL
			)
		AND
			deleted_at IS NULL
	`

	return database.ExecuteContext(ctx, o.db, query, userID)
}

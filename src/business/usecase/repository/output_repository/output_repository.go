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

func (o *OutputRepository) OutputByID(ctx context.Context, id, imageID uuid.UUID) (*output.Output, error) {
	query := `
		SELECT
			id,
			name,
			format,
			path,
			created_at,
			updated_at,
			deleted_at,
			confirmed,
			image_id
		FROM
			"output"
		WHERE
			id = $1
		AND
			image_id = $2
		AND
			deleted_at IS NULL
	`

	return database.QueryRowx[output.Output](ctx, o.db, query, id, imageID)
}

func (o *OutputRepository) OutputsByImageID(ctx context.Context, imageID uuid.UUID) ([]*output.Output, error) {
	query := `
		SELECT
			id,
			name,
			format,
			path,
			created_at,
			updated_at,
			deleted_at,
			confirmed,
			image_id
		FROM
			"output"
		WHERE
			image_id = $1
		AND
			deleted_at IS NULL
	`

	return database.Queryx[output.Output](ctx, o.db, query, imageID)
}

func (o *OutputRepository) RegisterOutput(ctx context.Context, name, format, path string, imageID uuid.UUID) (*output.Output, error) {
	query := `
		INSERT INTO
			"output" (name, format, path, image_id)
		VALUES
			($1, $2, $3, $4)
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

	return database.QueryRowx[output.Output](ctx, o.db, query, name, format, path, imageID)
}

func (o *OutputRepository) UpdateOutputByID(ctx context.Context, id uuid.UUID, confirmed bool) (*output.Output, error) {
	query := `
		UPDATE
			"output"
		SET
			confirmed = $1,
			updated_at = now()
		WHERE
			id = $2
		AND
			deleted_at IS NULL
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

	return database.QueryRowx[output.Output](ctx, o.db, query, confirmed, id)
}

func (o *OutputRepository) DeleteOutputByID(ctx context.Context, id uuid.UUID) (int, error) {
	query := `
		UPDATE
			"output"
		SET
			deleted_at = now()
		WHERE
			id = $1
		AND
            deleted_at IS NULL
	`

	return database.ExecuteContext(ctx, o.db, query, id)
}

func (o *OutputRepository) DeleteOutputsByImageID(ctx context.Context, imageID uuid.UUID) (int, error) {
	query := `
        UPDATE
            "output"
        SET
            deleted_at = now()
        WHERE
            image_id = $1
        AND
            deleted_at IS NULL
    `

	return database.ExecuteContext(ctx, o.db, query, imageID)
}

package output

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/output"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SQLOutputQuerier struct {
	db *sqlx.DB
}

func NewSQLOutputQuerier(db *sqlx.DB) *SQLOutputQuerier {
	return &SQLOutputQuerier{
		db: db,
	}
}

func (q *SQLOutputQuerier) RegisterOutput(ctx context.Context, name, format, path string, imageID, userID uuid.UUID) (*output.Output, error) {
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

	var o output.Output
	err := q.db.QueryRowxContext(ctx, query, name, format, path, imageID, userID).StructScan(&o)

	return &o, err
}

func (q *SQLOutputQuerier) OutputByID(ctx context.Context, outputID, imageID, userID uuid.UUID) (*output.Output, error) {
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

	var o output.Output
	err := q.db.QueryRowxContext(ctx, query, outputID, imageID, userID).StructScan(&o)

	return &o, err
}

func (q *SQLOutputQuerier) OutputsByImageID(ctx context.Context, imageID, userID uuid.UUID) ([]*output.Output, error) {
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

	rows, err := q.db.QueryxContext(ctx, query, imageID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var outputs []*output.Output
	for rows.Next() {
		var o output.Output
		if err := rows.StructScan(&o); err != nil {
			return nil, err
		}
		outputs = append(outputs, &o)
	}

	return outputs, err
}

func (q *SQLOutputQuerier) UpdateOutputByID(ctx context.Context, confirmed bool, outputID, imageID, userID uuid.UUID) (*output.Output, error) {
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

	var o output.Output
	err := q.db.QueryRowxContext(ctx, query, confirmed, outputID, imageID, userID).StructScan(&o)

	return &o, err
}

func (q *SQLOutputQuerier) DeleteOutputByID(ctx context.Context, outputID, imageID, userID uuid.UUID) error {
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

	_, err := q.db.ExecContext(ctx, query, outputID, imageID, userID)
	return err
}

func (q *SQLOutputQuerier) DeleteOutputsByImageID(ctx context.Context, imageID, userID uuid.UUID) error {
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

	_, err := q.db.ExecContext(ctx, query, imageID, userID)
	return err
}

func (q *SQLOutputQuerier) DeleteUserOutputs(ctx context.Context, userID uuid.UUID) error {
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

	_, err := q.db.ExecContext(ctx, query, userID)
	return err
}

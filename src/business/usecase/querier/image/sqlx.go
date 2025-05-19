package image

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/image"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SQLImageQuerier struct {
	db *sqlx.DB
}

func NewSQLImageQuerier(db *sqlx.DB) *SQLImageQuerier {
	return &SQLImageQuerier{
		db: db,
	}
}

func (q *SQLImageQuerier) RegisterImage(ctx context.Context, name, format, path string, userID uuid.UUID) (*image.Image, error) {
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

	var i image.Image
	err := q.db.QueryRowxContext(ctx, query, name, format, path, userID).StructScan(&i)

	return &i, err
}

func (q *SQLImageQuerier) ImageByID(ctx context.Context, imageID, userID uuid.UUID) (*image.Image, error) {
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

	var i image.Image
	err := q.db.QueryRowxContext(ctx, query, imageID, userID).StructScan(&i)

	return &i, err
}

func (q *SQLImageQuerier) ImagesByUserID(ctx context.Context, userID uuid.UUID) ([]*image.Image, error) {
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

	rows, err := q.db.QueryxContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	var images []*image.Image
	for rows.Next() {
		var i image.Image
		if err := rows.StructScan(&i); err != nil {
			return nil, err
		}

		images = append(images, &i)
	}

	return images, err
}

func (q *SQLImageQuerier) DeleteImageByID(ctx context.Context, imageID, userID uuid.UUID) error {
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

	_, err := q.db.ExecContext(ctx, query, imageID, userID)
	return err
}

func (q *SQLImageQuerier) DeleteUserImages(ctx context.Context, userID uuid.UUID) error {
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

	_, err := q.db.ExecContext(ctx, query, userID)
	return err
}

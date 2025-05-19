package session

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/session"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SQLSessionQuerier struct {
	db *sqlx.DB
}

func NewSQLSessionQuerier(db *sqlx.DB) *SQLSessionQuerier {
	return &SQLSessionQuerier{
		db: db,
	}
}

func (q *SQLSessionQuerier) CreateSession(ctx context.Context, userID uuid.UUID) (*session.Session, error) {
	query := `
		INSERT INTO
			"session" (user_id, expires_at)
		VALUES
			($1, now() + interval '1 hours')
		RETURNING token
	`

	var s session.Session
	err := q.db.QueryRowxContext(ctx, query, userID).StructScan(&s)

	return &s, err
}

func (q *SQLSessionQuerier) GetValidSession(ctx context.Context, token uuid.UUID) (*session.Session, error) {
	query := `
		SELECT
			token,
			user_id,
			created_at,
			expires_at
		FROM
			"session"
		WHERE
			token = $1
		AND
			expires_at >= now()
		ORDER BY
			created_at
		DESC LIMIT
			1
	`

	var s session.Session
	err := q.db.QueryRowxContext(ctx, query, token).StructScan(&s)

	return &s, err
}

func (q *SQLSessionQuerier) DeleteSession(ctx context.Context, token, userID uuid.UUID) error {
	query := `
		DELETE FROM
			"session"
		WHERE
			token = $1
		AND
			user_id = $2
	`

	_, err := q.db.ExecContext(ctx, query, token, userID)
	return err
}

func (q *SQLSessionQuerier) ClearSessionsByUserID(ctx context.Context, userID uuid.UUID) error {
	query := `
		DELETE FROM
			"session"
		WHERE
			user_id = $1
	`

	_, err := q.db.ExecContext(ctx, query, userID)
	return err
}

func (q *SQLSessionQuerier) ClearExpiredSessions(ctx context.Context) error {
	query := `
		DELETE FROM
			"session"
		WHERE
			expires_at < now()
	`

	_, err := q.db.ExecContext(ctx, query)
	return err
}

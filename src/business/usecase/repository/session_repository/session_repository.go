package sessionrepository

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/session"
	"github.com/erlendromo/forsete-atr/src/database"
	"github.com/erlendromo/forsete-atr/src/querier"
	"github.com/erlendromo/forsete-atr/src/querier/sqlx"
	"github.com/google/uuid"
)

type SessionRepository struct {
	querier querier.Querier[session.Session]
}

func NewSessionRepository(db database.Database) *SessionRepository {
	return &SessionRepository{
		querier: sqlx.NewSqlxQuerier[session.Session](db),
	}
}

func (s *SessionRepository) CreateSession(ctx context.Context, userID uuid.UUID) (*session.Session, error) {
	query := `
		INSERT INTO
			"session" (user_id, expires_at)
		VALUES
			($1, now() + interval '1 hours')
		RETURNING token
	`

	return s.querier.QueryRowx(ctx, query, userID)
}

func (s *SessionRepository) DeleteSession(ctx context.Context, token, userID uuid.UUID) error {
	query := `
		DELETE FROM
			"session"
		WHERE
			token = $1
		AND
			user_id = $2
	`

	return s.querier.Executex(ctx, query, token, userID)
}

func (s *SessionRepository) GetValidSession(ctx context.Context, token uuid.UUID) (*session.Session, error) {
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

	return s.querier.QueryRowx(ctx, query, token)
}

func (s *SessionRepository) ClearSessionsByUserID(ctx context.Context, userID uuid.UUID) error {
	query := `
		DELETE FROM
			"session"
		WHERE
			user_id = $1
	`

	return s.querier.Executex(ctx, query, userID)
}

func (s *SessionRepository) ClearExpiredSessions(ctx context.Context) error {
	query := `
		DELETE FROM
			"session"
		WHERE
			expires_at < now()
	`

	return s.querier.Executex(ctx, query)
}

package sessionrepository

import (
	"context"
	"fmt"
	"time"

	"github.com/erlendromo/forsete-atr/src/business/domain/session"
	"github.com/erlendromo/forsete-atr/src/database"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SessionRepository struct {
	db *sqlx.DB
}

func NewSessionRepository(db *sqlx.DB) *SessionRepository {
	return &SessionRepository{
		db: db,
	}
}

func (s *SessionRepository) CreateSession(ctx context.Context, userID uuid.UUID) (*session.Session, error) {
	query := `
		INSERT INTO
			"session" (user_id, expires_at)
		VALUES
			($1, $2)
		RETURNING token
	`

	return database.QueryRowx[session.Session](ctx, s.db, query, userID, time.Now().UTC().Add(24*time.Hour))
}

func (s *SessionRepository) DeleteSession(ctx context.Context, token uuid.UUID) error {
	query := `
		DELETE FROM
			"session"
		WHERE
			token = $1
	`

	rowsAffected, err := database.ExecuteContext(ctx, s.db, query, token)
	if err != nil {
		return err
	}

	fmt.Printf("\nDeleted %d session(s)\n", rowsAffected)

	return nil
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
			expires_at >= $2
		ORDER BY
			created_at
		DESC LIMIT
			1
	`

	return database.QueryRowx[session.Session](ctx, s.db, query, token, time.Now().UTC())
}

func (s *SessionRepository) ClearExpiredSessions(ctx context.Context) (int, error) {
	query := `
		DELETE FROM
			"session"
		WHERE
			expires_at < $1
	`

	result, err := s.db.ExecContext(ctx, query, time.Now().UTC())
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsAffected), nil
}

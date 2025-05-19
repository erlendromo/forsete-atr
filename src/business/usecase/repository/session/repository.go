package session

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/session"
	querier "github.com/erlendromo/forsete-atr/src/business/usecase/querier/session"
	"github.com/google/uuid"
)

type SessionRepository struct {
	querier querier.SessionQuerier
}

func NewSessionRepository(q querier.SessionQuerier) *SessionRepository {
	return &SessionRepository{
		querier: q,
	}
}

func (s *SessionRepository) CreateSession(ctx context.Context, userID uuid.UUID) (*session.Session, error) {
	return s.querier.CreateSession(ctx, userID)
}

func (s *SessionRepository) GetValidSession(ctx context.Context, token uuid.UUID) (*session.Session, error) {
	return s.querier.GetValidSession(ctx, token)
}

func (s *SessionRepository) DeleteSession(ctx context.Context, token, userID uuid.UUID) error {
	return s.querier.DeleteSession(ctx, token, userID)
}

func (s *SessionRepository) ClearSessionsByUserID(ctx context.Context, userID uuid.UUID) error {
	return s.querier.ClearSessionsByUserID(ctx, userID)
}

func (s *SessionRepository) ClearExpiredSessions(ctx context.Context) error {
	return s.querier.ClearExpiredSessions(ctx)
}

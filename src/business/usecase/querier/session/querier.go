package session

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/session"
	"github.com/google/uuid"
)

type SessionQuerier interface {
	CreateSession(ctx context.Context, userID uuid.UUID) (*session.Session, error)
	GetValidSession(ctx context.Context, token uuid.UUID) (*session.Session, error)
	DeleteSession(ctx context.Context, token, userID uuid.UUID) error
	ClearSessionsByUserID(ctx context.Context, userID uuid.UUID) error
	ClearExpiredSessions(ctx context.Context) error
}

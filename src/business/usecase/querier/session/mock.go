package session

import (
	"context"
	"errors"
	"time"

	"github.com/erlendromo/forsete-atr/src/business/domain/session"
	"github.com/google/uuid"
)

type MockSessionQuerier struct {
	sessions map[uuid.UUID]*session.Session
}

func NewMockSessionQuerier() *MockSessionQuerier {
	return &MockSessionQuerier{
		sessions: make(map[uuid.UUID]*session.Session),
	}
}

func (m *MockSessionQuerier) CreateSession(ctx context.Context, userID uuid.UUID) (*session.Session, error) {
	token := uuid.New()
	s := &session.Session{
		Token:     token,
		UserID:    userID,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(1 * time.Hour),
	}
	m.sessions[token] = s
	return s, nil
}

func (m *MockSessionQuerier) GetValidSession(ctx context.Context, token uuid.UUID) (*session.Session, error) {
	s, ok := m.sessions[token]
	if !ok {
		return nil, errors.New("session not found")
	}
	if s.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("session expired")
	}
	return s, nil
}

func (m *MockSessionQuerier) DeleteSession(ctx context.Context, token, userID uuid.UUID) error {
	s, ok := m.sessions[token]
	if !ok {
		return errors.New("session not found")
	}
	if s.UserID != userID {
		return errors.New("user ID mismatch")
	}
	delete(m.sessions, token)
	return nil
}

func (m *MockSessionQuerier) ClearSessionsByUserID(ctx context.Context, userID uuid.UUID) error {
	found := false
	for k, s := range m.sessions {
		if s.UserID == userID {
			found = true
			delete(m.sessions, k)
		}
	}

	if !found {
		return errors.New("not found")
	}

	return nil
}

func (m *MockSessionQuerier) ClearExpiredSessions(ctx context.Context) error {
	now := time.Now()
	for k, s := range m.sessions {
		if s.ExpiresAt.Before(now) {
			delete(m.sessions, k)
		}
	}
	return nil
}

func (m *MockSessionQuerier) Seed(sessions []*session.Session) {
	for _, s := range sessions {
		m.sessions[s.Token] = s
	}
}

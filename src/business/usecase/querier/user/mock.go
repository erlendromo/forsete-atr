package user

import (
	"context"
	"errors"
	"time"

	"github.com/erlendromo/forsete-atr/src/business/domain/user"
	"github.com/google/uuid"
)

type MockUserQuerier struct {
	users map[uuid.UUID]*user.User
}

func NewMockUserQuerier() *MockUserQuerier {
	return &MockUserQuerier{
		users: make(map[uuid.UUID]*user.User),
	}
}

func (m *MockUserQuerier) RegisterUser(ctx context.Context, email, hashedPassword string) (*user.User, error) {
	for _, u := range m.users {
		if u.Email == email && u.DeletedAt == nil {
			return nil, errors.New("user already exists")
		}
	}

	newUser := &user.User{
		ID:        uuid.New(),
		Email:     email,
		Password:  hashedPassword,
		CreatedAt: time.Now().UTC(),
		RoleID:    2,
		RoleName:  "Default",
	}

	m.users[newUser.ID] = newUser
	return newUser, nil
}

func (m *MockUserQuerier) GetUserByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	if u, ok := m.users[id]; ok {
		return u, nil
	}

	return nil, errors.New("not found")
}

func (m *MockUserQuerier) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	for _, u := range m.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, errors.New("not found")
}

func (m *MockUserQuerier) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	if u, ok := m.users[id]; ok {
		now := time.Now().UTC()
		u.DeletedAt = &now
		return nil
	}

	return errors.New("not found")
}

func (m *MockUserQuerier) Seed(users []*user.User) {
	for _, u := range users {
		m.users[u.ID] = u
	}
}

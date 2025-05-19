package user

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/user"
	"github.com/google/uuid"
)

type UserQuerier interface {
	RegisterUser(ctx context.Context, email, hashedPassword string) (*user.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*user.User, error)
	GetUserByEmail(ctx context.Context, email string) (*user.User, error)
	DeleteUserByID(ctx context.Context, id uuid.UUID) error
}

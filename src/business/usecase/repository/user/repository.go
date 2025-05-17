package userrepository

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/user"
	querier "github.com/erlendromo/forsete-atr/src/business/usecase/querier/user"
	"github.com/google/uuid"
)

type UserRepository struct {
	querier querier.UserQuerier
}

func NewUserRepository(q querier.UserQuerier) *UserRepository {
	return &UserRepository{
		querier: q,
	}
}

func (u *UserRepository) RegisterUser(ctx context.Context, email, hashedPassword string) (*user.User, error) {
	return u.querier.RegisterUser(ctx, email, hashedPassword)
}

func (u *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	return u.querier.GetUserByID(ctx, id)
}

func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	return u.querier.GetUserByEmail(ctx, email)
}

func (u *UserRepository) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	return u.querier.DeleteUserByID(ctx, id)
}

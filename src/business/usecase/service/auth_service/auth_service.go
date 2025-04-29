package authservice

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/session"
	"github.com/erlendromo/forsete-atr/src/business/domain/user"
	sessionrepository "github.com/erlendromo/forsete-atr/src/business/usecase/repository/session_repository"
	userrepository "github.com/erlendromo/forsete-atr/src/business/usecase/repository/user_repository"
	"github.com/erlendromo/forsete-atr/src/util"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AuthService struct {
	UserRepo    *userrepository.UserRepository
	SessionRepo *sessionrepository.SessionRepository
}

func NewAuthService(db *sqlx.DB) *AuthService {
	return &AuthService{
		UserRepo:    userrepository.NewUserRepository(db),
		SessionRepo: sessionrepository.NewSessionRepository(db),
	}
}

func (a *AuthService) RegisterUser(ctx context.Context, email, password string) (*user.User, error) {
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return nil, err
	}

	return a.UserRepo.RegisterUser(ctx, email, hashedPassword)
}

func (a *AuthService) Login(ctx context.Context, email, password string) (*session.Session, error) {
	user, err := a.UserRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if err := user.ComparePassword(password); err != nil {
		return nil, err
	}

	return a.SessionRepo.CreateSession(ctx, user.ID)
}

func (a *AuthService) Logout(ctx context.Context, token uuid.UUID) error {
	if err := a.SessionRepo.DeleteSession(ctx, token); err != nil {
		return err
	}

	return nil
}

func (a *AuthService) IsAuthorized(ctx context.Context, token uuid.UUID) (*user.User, error) {
	session, err := a.SessionRepo.GetValidSession(ctx, token)
	if err != nil {
		return nil, err
	}

	return a.UserRepo.GetByID(ctx, session.UserID)
}

func (a *AuthService) RefreshToken(ctx context.Context, oldToken uuid.UUID) (*session.Session, error) {
	session, err := a.SessionRepo.GetValidSession(ctx, oldToken)
	if err != nil {
		return nil, err
	}

	if err := a.SessionRepo.DeleteSession(ctx, oldToken); err != nil {
		return nil, err
	}

	newSession, err := a.SessionRepo.CreateSession(ctx, session.UserID)
	if err != nil {
		return nil, err
	}

	return newSession, nil
}

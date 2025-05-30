package auth

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/session"
	"github.com/erlendromo/forsete-atr/src/business/domain/user"
	sessionrepository "github.com/erlendromo/forsete-atr/src/business/usecase/repository/session"
	userrepository "github.com/erlendromo/forsete-atr/src/business/usecase/repository/user"
	"github.com/erlendromo/forsete-atr/src/util"
	"github.com/google/uuid"
)

type AuthService struct {
	UserRepo    *userrepository.UserRepository
	SessionRepo *sessionrepository.SessionRepository
}

func NewAuthService(u *userrepository.UserRepository, s *sessionrepository.SessionRepository) *AuthService {
	return &AuthService{
		UserRepo:    u,
		SessionRepo: s,
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
	user, err := a.UserRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if err := user.ComparePassword(password); err != nil {
		return nil, err
	}

	// Clear ghost sessions
	if err := a.SessionRepo.ClearSessionsByUserID(ctx, user.ID); err != nil {
		return nil, err
	}

	return a.SessionRepo.CreateSession(ctx, user.ID)
}

func (a *AuthService) Logout(ctx context.Context, token, userID uuid.UUID) error {
	if err := a.SessionRepo.DeleteSession(ctx, token, userID); err != nil {
		return err
	}

	return nil
}

func (a *AuthService) IsAuthorized(ctx context.Context, token uuid.UUID) (*user.User, error) {
	session, err := a.SessionRepo.GetValidSession(ctx, token)
	if err != nil {
		return nil, err
	}

	return a.UserRepo.GetUserByID(ctx, session.UserID)
}

func (a *AuthService) RefreshToken(ctx context.Context, oldToken, userID uuid.UUID) (*session.Session, error) {
	session, err := a.SessionRepo.GetValidSession(ctx, oldToken)
	if err != nil {
		return nil, err
	}

	if err := a.SessionRepo.DeleteSession(ctx, oldToken, userID); err != nil {
		return nil, err
	}

	return a.SessionRepo.CreateSession(ctx, session.UserID)
}

func (a *AuthService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	if err := a.SessionRepo.ClearSessionsByUserID(ctx, userID); err != nil {
		return err
	}

	if err := a.UserRepo.DeleteUserByID(ctx, userID); err != nil {
		return err
	}

	return nil
}

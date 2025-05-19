package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/erlendromo/forsete-atr/src/business/domain/user"
	authservice "github.com/erlendromo/forsete-atr/src/business/usecase/service/auth"
	"github.com/erlendromo/forsete-atr/src/util"
	"github.com/google/uuid"
)

type contextValuesKey string

var ContextValuesKey contextValuesKey = "context_values"

type ContextValues struct {
	Token uuid.UUID
	User  *user.User
}

func AuthMiddleware(auth *authservice.AuthService, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := extractTokenFromHeader(r)
		if err != nil {
			util.ERROR(w, http.StatusUnauthorized, err)
			return
		}

		user, err := auth.IsAuthorized(r.Context(), token)
		if err != nil {
			util.ERROR(w, http.StatusUnauthorized, fmt.Errorf("unauthorized user, log in"))
			return
		}

		contextValues := &ContextValues{
			Token: token,
			User:  user,
		}

		ctx := context.WithValue(r.Context(), ContextValuesKey, contextValues)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func extractTokenFromHeader(r *http.Request) (uuid.UUID, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return uuid.Nil, fmt.Errorf("missing auth-header")
	}

	const prefix = "Bearer "
	if len(authHeader) <= len(prefix) || authHeader[:len(prefix)] != prefix {
		return uuid.Nil, fmt.Errorf("invalid auth-header")
	}

	token, err := uuid.Parse(authHeader[len(prefix):])
	if err != nil {
		return uuid.Nil, fmt.Errorf("unable to parse token")
	}

	return token, nil
}

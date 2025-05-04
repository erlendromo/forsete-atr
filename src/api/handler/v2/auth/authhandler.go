package auth

import (
	"fmt"
	"net/http"

	"github.com/erlendromo/forsete-atr/src/api/middleware"
	authservice "github.com/erlendromo/forsete-atr/src/business/usecase/service/auth_service"
	"github.com/erlendromo/forsete-atr/src/util"
)

// RegisterAndLoginRequest
//
//	@Summary		RegisterAndLoginRequest form
//	@Description	Body containing email and password.
type RegisterAndLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register
//
//	@Summary		Register user
//	@Description	Register user with email and password.
//	@Tags			Auth
//	@Accept			json
//	@Param request body			RegisterAndLoginRequest true "Register user form"
//	@Produce		json
//	@Success		201	{object}	user.User
//	@Failure		422	{object}	util.ErrorResponse
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v2/auth/register/ [post]
func Register(authService *authservice.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		register, err := util.DecodeJSON[RegisterAndLoginRequest](r.Body)
		if err != nil {
			util.NewInternalErrorLog("REGISTER", err).PrintLog("CLIENT ERROR")
			util.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("invalid request form"))
			return
		}

		user, err := authService.RegisterUser(r.Context(), register.Email, register.Password)
		if err != nil {
			util.NewInternalErrorLog("REGISTER", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		if err := user.CreateDirs(); err != nil {
			util.NewInternalErrorLog("REGISTER", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		util.EncodeJSON(w, http.StatusCreated, user)
	}
}

// Login
//
//	@Summary		Login as user
//	@Description	Login as user with email and password.
//	@Tags			Auth
//	@Accept			json
//	@Param request body			RegisterAndLoginRequest true "Login user form"
//	@Produce		json
//	@Success		201	{object}	session.Session
//	@Header			201	{string}	Authorization	"Bearer <token>"
//	@Failure		404	{object}	util.ErrorResponse
//	@Failure		422	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v2/auth/login/ [post]
func Login(authService *authservice.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		login, err := util.DecodeJSON[RegisterAndLoginRequest](r.Body)
		if err != nil {
			util.NewInternalErrorLog("LOGIN", err).PrintLog("CLIENT ERROR")
			util.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("invalid request form"))
			return
		}

		session, err := authService.Login(r.Context(), login.Email, login.Password)
		if err != nil {
			util.NewInternalErrorLog("LOGIN", err).PrintLog("CLIENT ERROR")
			util.ERROR(w, http.StatusNotFound, fmt.Errorf("invalid email and/or password"))
			return
		}

		w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", session.Token.String()))
		util.EncodeJSON(w, http.StatusOK, session)
	}
}

// Logout
//
//	@Summary		Logout as user
//	@Description	Logout as user.
//	@Tags			Auth
//	@Param			Authorization	header	string	true	"'Bearer <token>' must be set for valid response"
//	@Produce		json
//	@Success		204
//	@Failure		401	{object}	util.ErrorResponse
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v2/auth/logout/ [post]
func Logout(authService *authservice.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contextValues, ok := r.Context().Value(middleware.ContextValuesKey).(*middleware.ContextValues)
		if !ok {
			err := fmt.Errorf("missing 'context_values' in request-context")
			util.NewInternalErrorLog("LOGOUT", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		if err := authService.Logout(r.Context(), contextValues.Token); err != nil {
			util.NewInternalErrorLog("LOGOUT", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		util.EncodeJSON(w, http.StatusNoContent, nil)
	}
}

// Refresh
//
//	@Summary		Refresh token
//	@Description	Refresh session token.
//	@Tags			Auth
//	@Param			Authorization	header	string	true	"'Bearer <token>' must be set for valid response"
//	@Produce		json
//	@Success		200	{object}	session.Session
//	@Header			200	{string}	Authorization	"Bearer <token>"
//	@Failure		401	{object}	util.ErrorResponse
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v2/auth/refresh/ [post]
func RefreshSession(authService *authservice.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contextValues, ok := r.Context().Value(middleware.ContextValuesKey).(*middleware.ContextValues)
		if !ok {
			err := fmt.Errorf("missing 'context_values' in request-context")
			util.NewInternalErrorLog("REFRESH SESSION", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		session, err := authService.RefreshToken(r.Context(), contextValues.Token)
		if err != nil {
			util.NewInternalErrorLog("REFRESH SESSION", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", session.Token.String()))
		util.EncodeJSON(w, http.StatusOK, session)
	}
}

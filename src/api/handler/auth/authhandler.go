package auth

import (
	"fmt"
	"net/http"

	"github.com/erlendromo/forsete-atr/src/api/middleware"
	authservice "github.com/erlendromo/forsete-atr/src/business/usecase/auth_service"
	"github.com/erlendromo/forsete-atr/src/util"
)

type RegisterAndLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

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

		util.EncodeJSON(w, http.StatusCreated, user)
	}
}

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

		w.Header().Set("Authorization", "Bearer "+session.Token.String())
		util.EncodeJSON(w, http.StatusOK, session)
	}
}

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

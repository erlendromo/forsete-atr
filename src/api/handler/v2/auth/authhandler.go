package auth

import (
	"fmt"
	"net/http"
	"regexp"
	"unicode"

	"github.com/erlendromo/forsete-atr/src/api/middleware"
	atrservice "github.com/erlendromo/forsete-atr/src/business/usecase/service/atr_service"
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

// One character before @, valid domain format, no whitespace or invalid symbols, trailing format atleast 2 characters
func (rlf *RegisterAndLoginRequest) validEmail() error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if ok := emailRegex.MatchString(rlf.Email); !ok {
		return fmt.Errorf("invalid email: must be a valid address like 'name@example.com'")
	}

	return nil
}

// Atleast: one uppercase letter, one lowercase letter, one number, min length of 8
func (rlf *RegisterAndLoginRequest) validPassword() error {
	passwordRegex := regexp.MustCompile(`^[a-zA-Z0-9!@#\$%\^&\*\(\)_\+\-=\[\]{};':"\\|,.<>\/?` + "`" + `~]{8,}$`)
	if ok := passwordRegex.MatchString(rlf.Password); !ok {
		return fmt.Errorf("malformed password: must be at least 8 characters and contain only valid characters")
	}

	var hasUpper, hasLower, hasNumber bool
	for _, c := range rlf.Password {
		switch {
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsDigit(c):
			hasNumber = true
		}
	}

	if !hasUpper {
		return fmt.Errorf("invalid password: must contain atleast one uppercase letter")
	} else if !hasLower {
		return fmt.Errorf("invalid password: must contain atleast one lowercase letter")
	} else if !hasNumber {
		return fmt.Errorf("invalid password: must contain atleast one number")
	}

	return nil
}

// If both email and password are valid -> return nil
func (rlf *RegisterAndLoginRequest) Validate() error {
	if err := rlf.validEmail(); err != nil {
		return err
	}

	if err := rlf.validPassword(); err != nil {
		return err
	}

	return nil
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
//	@Failure		400	{object}	util.ErrorResponse
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

		if err := register.Validate(); err != nil {
			util.ERROR(w, http.StatusBadRequest, err)
			return
		}

		user, err := authService.RegisterUser(r.Context(), register.Email, register.Password)
		if err != nil {
			// TODO Should check if error is constraint violation or not...
			util.NewInternalErrorLog("REGISTER", err).PrintLog("CLIENT ERROR")
			util.ERROR(w, http.StatusBadRequest, fmt.Errorf("email already in use"))
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
//	@Header			201	{string}	Authorization	"Bearer token"
//	@Failure		400	{object}	util.ErrorResponse
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

		if err := login.Validate(); err != nil {
			util.ERROR(w, http.StatusBadRequest, err)
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
//	@Param			Authorization	header	string	true	"'Bearer token' must be set for valid response"
//	@Produce		json
//	@Success		204
//	@Failure		401	{object}	util.ErrorResponse
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v2/auth/logout/ [post]
func Logout(authService *authservice.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxValues, ok := r.Context().Value(middleware.ContextValuesKey).(*middleware.ContextValues)
		if !ok {
			err := fmt.Errorf("missing 'context_values' in request-context")
			util.NewInternalErrorLog("LOGOUT", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		if err := authService.Logout(r.Context(), ctxValues.Token, ctxValues.User.ID); err != nil {
			util.NewInternalErrorLog("LOGOUT", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		util.EncodeJSON(w, http.StatusNoContent, nil)
	}
}

// RefreshSession
//
//	@Summary		Refresh token
//	@Description	Refresh session token.
//	@Tags			Auth
//	@Param			Authorization	header	string	true	"'Bearer token' must be set for valid response"
//	@Produce		json
//	@Success		200	{object}	session.Session
//	@Header			200	{string}	Authorization	"Bearer <token>"
//	@Failure		401	{object}	util.ErrorResponse
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v2/auth/refresh/ [post]
func RefreshSession(authService *authservice.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxValues, ok := r.Context().Value(middleware.ContextValuesKey).(*middleware.ContextValues)
		if !ok {
			err := fmt.Errorf("missing 'context_values' in request-context")
			util.NewInternalErrorLog("REFRESH SESSION", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		session, err := authService.RefreshToken(r.Context(), ctxValues.Token, ctxValues.User.ID)
		if err != nil {
			util.NewInternalErrorLog("REFRESH SESSION", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", session.Token.String()))
		util.EncodeJSON(w, http.StatusOK, session)
	}
}

// DeleteUser
//
//	@Summary		Delete user
//	@Description	Delete user and all its data.
//	@Tags			Auth
//	@Param			Authorization	header	string	true	"'Bearer token' must be set for valid response"
//	@Produce		json
//	@Success		204
//	@Failure		401	{object}	util.ErrorResponse
//	@Failure		500	{object}	util.ErrorResponse
//	@Router			/forsete-atr/v2/auth/delete/ [delete]
func DeleteUser(authService *authservice.AuthService, atrService *atrservice.ATRService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxValues, ok := r.Context().Value(middleware.ContextValuesKey).(*middleware.ContextValues)
		if !ok {
			err := fmt.Errorf("missing 'context_values' in request-context")
			util.NewInternalErrorLog("DELETE USER", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		if err := atrService.DeleteUserOutputsAndImages(r.Context(), ctxValues.User.ID); err != nil {
			util.NewInternalErrorLog("DELETE USER (OUTPUTS/IMAGES)", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		if err := authService.DeleteUser(r.Context(), ctxValues.User.ID); err != nil {
			util.NewInternalErrorLog("DELETE USER (USER)", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		if err := ctxValues.User.RemoveData(); err != nil {
			util.NewInternalErrorLog("DELETE USER (FILES)", err).PrintLog("SERVER ERROR")
			util.ERROR(w, http.StatusInternalServerError, fmt.Errorf(util.INTERNAL_SERVER_ERROR))
			return
		}

		util.EncodeJSON(w, http.StatusNoContent, nil)
	}
}

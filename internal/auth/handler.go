package auth

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/naresh-official/ecommerce_go/configs"
	"github.com/naresh-official/ecommerce_go/internal/response"
	"github.com/naresh-official/ecommerce_go/internal/validator"
)

type Handler struct {
	service *Service
	cfg     *configs.AppConfig
}

func NewHandler(service *Service, cfg *configs.AppConfig) *Handler {
	return &Handler{
		service: service,
		cfg:     cfg,
	}
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {

	role := chi.URLParam(r, "role")

	if role != "user" && role != "admin" && role != "seller" {
		http.NotFound(w, r)
		return
	}

	var req SignUpRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"Invalid request body",
		)
		return
	}

	if err := validator.ValidateRequest(req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	signupResponse, err := h.service.SignUp(r.Context(), req, role)

	if err != nil {

		if errors.Is(err, ErrEmailAlreadyExists) {
			response.Error(
				w,
				http.StatusConflict,
				err.Error(),
			)
			return
		}
		response.Error(
			w,
			http.StatusInternalServerError,
			"Internal server error",
		)
		slog.Error("Error signing up user", "error", err)
		return
	}
	response.Json(
		w,
		http.StatusCreated,
		"User created successfully",
		signupResponse,
	)
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	role := chi.URLParam(r, "role")

	if role != "user" && role != "admin" && role != "seller" {
		http.NotFound(w, r)
		return
	}

	var req SignInRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"Invalid request body",
		)
		return
	}

	if err := validator.ValidateRequest(req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	signInResult, err := h.service.SignIn(r.Context(), req, role)

	if err != nil {

		if errors.Is(err, ErrUserNotFound) {
			response.Error(
				w,
				http.StatusNotFound,
				err.Error(),
			)
			return
		}

		if errors.Is(err, ErrInvalidCredentials) {
			response.Error(
				w,
				http.StatusUnauthorized,
				err.Error(),
			)
			return
		}

		response.Error(
			w,
			http.StatusInternalServerError,
			"Internal server error",
		)
		slog.Error("Error signing in user", "error", err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    signInResult.AccessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   h.cfg.Env != "development",
		SameSite: http.SameSiteStrictMode,
		MaxAge:   int(h.service.cfg.AccessTokenExpiration.Seconds()),
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    signInResult.RefreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   h.cfg.Env != "development",
		SameSite: http.SameSiteStrictMode,
		MaxAge:   int(h.service.cfg.RefreshTokenExpiration.Seconds()),
	})

	response.Json(
		w,
		http.StatusOK,
		"User signed in successfully",
		signInResult.User,
	)
}

func (h *Handler) GetMe(w http.ResponseWriter, r *http.Request) {
	userID := GetUserIdFromContext(r.Context())

	getMeResponse, err := h.service.GetMeUser(r.Context(), userID)

	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			response.Error(
				w,
				http.StatusNotFound,
				err.Error(),
			)
			return
		}

		response.Error(
			w,
			http.StatusInternalServerError,
			"Internal server error",
		)
		slog.Error("Error getting user", "error", err)
		return
	}

	response.Json(
		w,
		http.StatusOK,
		"User fetched successfully",
		getMeResponse,
	)
}

func (h *Handler) SignOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   h.cfg.Env != "development",
		SameSite: http.SameSiteStrictMode,
		MaxAge:   -1,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   h.cfg.Env != "development",
		SameSite: http.SameSiteStrictMode,
		MaxAge:   -1,
	})

	response.Json(
		w,
		http.StatusOK,
		"User signed out successfully",
		nil,
	)
}

func (h *Handler) RefreshAccessToken(w http.ResponseWriter, r *http.Request) {
	refreshTokenCookie, err := r.Cookie("refresh_token")
	if err != nil || refreshTokenCookie.Value == "" {
		response.Error(
			w,
			http.StatusUnauthorized,
			"Refresh token not found",
		)
		return
	}

	updateTokenResult, err := h.service.UpdateAccessToken(r.Context(), refreshTokenCookie.Value)
	if err != nil {
		response.Error(
			w,
			http.StatusUnauthorized,
			"Invalid refresh token",
		)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    updateTokenResult.AccessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   h.cfg.Env != "development",
		SameSite: http.SameSiteStrictMode,
		MaxAge:   int(h.service.cfg.AccessTokenExpiration.Seconds()),
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    updateTokenResult.RefreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   h.cfg.Env != "development",
		SameSite: http.SameSiteStrictMode,
		MaxAge:   int(h.service.cfg.RefreshTokenExpiration.Seconds()),
	})

	response.Json(
		w,
		http.StatusOK,
		"Token refreshed successfully",
		nil,
	)
}

package user

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/naresh-official/ecommerce_go/internal/auth"
	"github.com/naresh-official/ecommerce_go/internal/response"
	"github.com/naresh-official/ecommerce_go/internal/validator"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIdFromContext(r.Context())

	profile, err := h.service.GetProfile(r.Context(), userID)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}

		response.Error(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Error fetching user profile", "error", err)
		return
	}

	response.Json(w, http.StatusOK, "Profile fetched successfully", profile)
}

func (h *Handler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIdFromContext(r.Context())

	var req UpdateProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := validator.ValidateRequest(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	profile, err := h.service.UpdateProfile(r.Context(), userID, req)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		if errors.Is(err, ErrEmailAlreadyExists) {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}

		response.Error(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Error updating user profile", "error", err)
		return
	}

	response.Json(w, http.StatusOK, "Profile updated successfully", profile)
}

func (h *Handler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIdFromContext(r.Context())

	var req ChangePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := validator.ValidateRequest(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.ChangePassword(r.Context(), userID, req)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		if errors.Is(err, ErrInvalidCurrentPassword) {
			response.Error(w, http.StatusUnauthorized, err.Error())
			return
		}

		response.Error(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Error changing password", "error", err)
		return
	}

	response.Json(w, http.StatusOK, "Password changed successfully", nil)
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response.Json(w, http.StatusOK, "User module is healthy", nil)
}

var _ = chi.URLParam

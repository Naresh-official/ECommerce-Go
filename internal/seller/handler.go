package seller

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/naresh-official/ecommerce_go/internal/response"
	"github.com/naresh-official/ecommerce_go/internal/validator"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetMySeller(w http.ResponseWriter, r *http.Request) {
	seller, err := h.service.GetMySeller(r.Context())
	if err != nil {
		if errors.Is(err, ErrSellerNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Error fetching seller", "error", err)
		return
	}

	response.Json(w, http.StatusOK, "Seller fetched successfully", seller)
}

func (h *Handler) CreateSeller(w http.ResponseWriter, r *http.Request) {
	var req CreateSellerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := validator.ValidateRequest(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	seller, err := h.service.CreateSeller(r.Context(), req)
	if err != nil {
		if errors.Is(err, ErrSellerAlreadyExists) || errors.Is(err, ErrStoreNameAlreadyTaken) {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Error creating seller", "error", err)
		return
	}

	response.Json(w, http.StatusCreated, "Seller created successfully", seller)
}

func (h *Handler) UpdateSeller(w http.ResponseWriter, r *http.Request) {
	var req UpdateSellerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := validator.ValidateRequest(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	seller, err := h.service.UpdateSeller(r.Context(), req)
	if err != nil {
		if errors.Is(err, ErrSellerNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Error updating seller", "error", err)
		return
	}

	response.Json(w, http.StatusOK, "Seller updated successfully", seller)
}

func (h *Handler) VerifySeller(w http.ResponseWriter, r *http.Request) {
	sellerID := chi.URLParam(r, "sellerId")

	seller, err := h.service.VerifySeller(r.Context(), sellerID)
	if err != nil {
		if errors.Is(err, ErrSellerNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Error verifying seller", "error", err)
		return
	}

	response.Json(w, http.StatusOK, "Seller verified successfully", seller)
}

func (h *Handler) UnverifySeller(w http.ResponseWriter, r *http.Request) {
	sellerID := chi.URLParam(r, "sellerId")

	seller, err := h.service.UnverifySeller(r.Context(), sellerID)
	if err != nil {
		if errors.Is(err, ErrSellerNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Error unverifying seller", "error", err)
		return
	}

	response.Json(w, http.StatusOK, "Seller unverified successfully", seller)
}

func (h *Handler) ActivateSeller(w http.ResponseWriter, r *http.Request) {
	sellerID := chi.URLParam(r, "sellerId")

	seller, err := h.service.ActivateSeller(r.Context(), sellerID)
	if err != nil {
		if errors.Is(err, ErrSellerNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Error activating seller", "error", err)
		return
	}

	response.Json(w, http.StatusOK, "Seller activated successfully", seller)
}

func (h *Handler) DeactivateSeller(w http.ResponseWriter, r *http.Request) {
	sellerID := chi.URLParam(r, "sellerId")

	seller, err := h.service.DeactivateSeller(r.Context(), sellerID)
	if err != nil {
		if errors.Is(err, ErrSellerNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Error deactivating seller", "error", err)
		return
	}

	response.Json(w, http.StatusOK, "Seller deactivated successfully", seller)
}

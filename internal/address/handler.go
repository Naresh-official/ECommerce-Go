package address

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/naresh-official/ecommerce_go/internal/response"
	"github.com/naresh-official/ecommerce_go/internal/validator"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetAllAddresses(w http.ResponseWriter, r *http.Request) {
	addresses, err := h.service.GetAllAddresses(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch addresses")
		return
	}
	response.Json(w, http.StatusOK, "Addresses fetched successfully", addresses)
}

func (h *Handler) CreateAddress(w http.ResponseWriter, r *http.Request) {
	var req CreateAddressRequest
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

	req.City = strings.ToLower(req.City)
	req.State = strings.ToLower(req.State)
	req.Country = strings.ToLower(req.Country)

	address, err := h.service.CreateAddress(r.Context(), req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to create address")
		return
	}
	response.Json(w, http.StatusCreated, "Address created successfully", address)
}

func (h *Handler) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	addressId := chi.URLParam(r, "addressId")

	err := h.service.DeleteAddress(r.Context(), addressId)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to delete address")
		return
	}
	response.Json(w, http.StatusOK, "Address deleted successfully", nil)
}

func (h *Handler) GetAddressById(w http.ResponseWriter, r *http.Request) {
	addressId := chi.URLParam(r, "addressId")

	address, err := h.service.GetAddressById(r.Context(), addressId)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch address")
		return
	}
	response.Json(w, http.StatusOK, "Address fetched successfully", address)
}

func (h *Handler) UpdateAddress(w http.ResponseWriter, r *http.Request) {
	addressId := chi.URLParam(r, "addressId")

	var req CreateAddressRequest
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

	address, err := h.service.UpdateAddress(r.Context(), addressId, req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to update address")
		return
	}
	response.Json(w, http.StatusOK, "Address updated successfully", address)
}

func (h *Handler) SetDefaultAddress(w http.ResponseWriter, r *http.Request) {
	addressId := chi.URLParam(r, "addressId")

	err := h.service.SetDefaultAddress(r.Context(), addressId)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to set default address")
		return
	}
	response.Json(w, http.StatusOK, "Default address set successfully", nil)
}

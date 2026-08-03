package product

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

func (h *Handler) GetMyProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetMyProducts(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Error fetching products", "error", err)
		return
	}

	response.Json(w, http.StatusOK, "Products fetched successfully", products)
}

func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "productId")

	product, err := h.service.GetProductByID(r.Context(), productID)
	if err != nil {
		if errors.Is(err, ErrProductNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Error fetching product", "error", err)
		return
	}

	response.Json(w, http.StatusOK, "Product fetched successfully", product)
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := validator.ValidateRequest(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.service.CreateProduct(r.Context(), req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Error creating product", "error", err)
		return
	}

	response.Json(w, http.StatusCreated, "Product created successfully", product)
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "productId")

	var req UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := validator.ValidateRequest(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.service.UpdateProduct(r.Context(), productID, req)
	if err != nil {
		if errors.Is(err, ErrProductNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Error updating product", "error", err)
		return
	}

	response.Json(w, http.StatusOK, "Product updated successfully", product)
}

func (h *Handler) UpdateProductPrice(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "productId")

	var req UpdateProductPriceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := validator.ValidateRequest(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.service.UpdateProductPrice(r.Context(), productID, req.Price)
	if err != nil {
		if errors.Is(err, ErrProductNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Error updating product price", "error", err)
		return
	}

	response.Json(w, http.StatusOK, "Product price updated successfully", product)
}

func (h *Handler) UpdateProductStockQuantity(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "productId")

	var req UpdateProductStockQuantityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := validator.ValidateRequest(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.service.UpdateProductStockQuantity(r.Context(), productID, req.StockQuantity)
	if err != nil {
		if errors.Is(err, ErrProductNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Error updating product stock quantity", "error", err)
		return
	}

	response.Json(w, http.StatusOK, "Product stock quantity updated successfully", product)
}
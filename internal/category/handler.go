package category

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

func (h *Handler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetAllCategories(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch categories")
		return
	}

	response.Json(w, http.StatusOK, "Categories fetched successfully", categories)
}

func (h *Handler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	categoryID := chi.URLParam(r, "categoryId")

	category, err := h.service.GetCategoryByID(r.Context(), categoryID)
	if err != nil {
		if errors.Is(err, ErrCategoryNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}

		response.Error(w, http.StatusInternalServerError, "Failed to fetch category")
		slog.Error("Error fetching category", "error", err)
		return
	}

	response.Json(w, http.StatusOK, "Category fetched successfully", category)
}

func (h *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var req CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := validator.ValidateRequest(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	category, err := h.service.CreateCategory(r.Context(), req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to create category")
		slog.Error("Error creating category", "error", err)
		return
	}

	response.Json(w, http.StatusCreated, "Category created successfully", category)
}

func (h *Handler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	categoryID := chi.URLParam(r, "categoryId")

	var req CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := validator.ValidateRequest(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	category, err := h.service.UpdateCategory(r.Context(), categoryID, req)
	if err != nil {
		if errors.Is(err, ErrCategoryNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}

		response.Error(w, http.StatusInternalServerError, "Failed to update category")
		slog.Error("Error updating category", "error", err)
		return
	}

	response.Json(w, http.StatusOK, "Category updated successfully", category)
}

func (h *Handler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	categoryID := chi.URLParam(r, "categoryId")

	if err := h.service.DeleteCategory(r.Context(), categoryID); err != nil {
		if errors.Is(err, ErrCategoryNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}

		response.Error(w, http.StatusInternalServerError, "Failed to delete category")
		slog.Error("Error deleting category", "error", err)
		return
	}

	response.Json(w, http.StatusOK, "Category deleted successfully", nil)
}

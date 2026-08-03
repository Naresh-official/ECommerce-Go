package category

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(
	r chi.Router,
	authMiddleware func(http.Handler) http.Handler,
	adminMiddleware func(http.Handler) http.Handler,
	h *Handler,
) {
	r.Route("/category", func(r chi.Router) {
		r.Get("/", h.GetAllCategories)
		r.Get("/{categoryId}", h.GetCategoryByID)

		r.With(authMiddleware, adminMiddleware).Post("/", h.CreateCategory)
		r.With(authMiddleware, adminMiddleware).Put("/{categoryId}", h.UpdateCategory)
		r.With(authMiddleware, adminMiddleware).Delete("/{categoryId}", h.DeleteCategory)
	})
}

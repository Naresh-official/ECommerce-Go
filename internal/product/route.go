package product

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(
	r chi.Router,
	authMiddleware func(http.Handler) http.Handler,
	sellerMiddleware func(http.Handler) http.Handler,
	h *Handler,
) {
	r.Route("/product", func(r chi.Router) {
		r.With(authMiddleware, sellerMiddleware).Get("/me", h.GetMyProducts)
		r.With(authMiddleware, sellerMiddleware).Post("/", h.CreateProduct)
		r.With(authMiddleware, sellerMiddleware).Put("/{productId}", h.UpdateProduct)
		r.With(authMiddleware, sellerMiddleware).Put("/{productId}/price", h.UpdateProductPrice)
		r.With(authMiddleware, sellerMiddleware).Put("/{productId}/stock", h.UpdateProductStockQuantity)
		r.With(authMiddleware, sellerMiddleware).Get("/{productId}", h.GetProductByID)
	})
}

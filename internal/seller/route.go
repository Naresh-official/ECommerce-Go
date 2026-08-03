package seller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(
	r chi.Router,
	authMiddleware func(http.Handler) http.Handler,
	sellerMiddleware func(http.Handler) http.Handler,
	adminMiddleware func(http.Handler) http.Handler,
	h *Handler,
) {
	r.Route("/seller", func(r chi.Router) {
		r.With(authMiddleware, sellerMiddleware).Get("/me", h.GetMySeller)
		r.With(authMiddleware, sellerMiddleware).Post("/", h.CreateSeller)
		r.With(authMiddleware, sellerMiddleware).Put("/", h.UpdateSeller)

		r.With(authMiddleware, adminMiddleware).Put("/{sellerId}/verify", h.VerifySeller)
		r.With(authMiddleware, adminMiddleware).Put("/{sellerId}/unverify", h.UnverifySeller)
		r.With(authMiddleware, adminMiddleware).Put("/{sellerId}/activate", h.ActivateSeller)
		r.With(authMiddleware, adminMiddleware).Put("/{sellerId}/deactivate", h.DeactivateSeller)
	})
}

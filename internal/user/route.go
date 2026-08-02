package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(
	r chi.Router,
	authMiddleware func(http.Handler) http.Handler,
	h *Handler,
) {
	r.With(authMiddleware).Route("/user", func(r chi.Router) {
		r.Get("/profile", h.GetProfile)
		r.Put("/profile", h.UpdateProfile)
		r.Put("/change-password", h.ChangePassword)
	})
}

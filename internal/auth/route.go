package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(
	r chi.Router,
	authMiddleware func(http.Handler) http.Handler,
	h *Handler,
) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/signup/{role}", h.SignUp)
		r.Post("/signin/{role}", h.SignIn)
		r.Post("/signout", h.SignOut)
		r.Post("/refresh", h.RefreshAccessToken)
		r.With(authMiddleware).Get("/me", h.GetMe)
	})
}

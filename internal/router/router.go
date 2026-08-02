package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/naresh-official/ecommerce_go/internal/auth"
	"github.com/naresh-official/ecommerce_go/internal/response"
	userpkg "github.com/naresh-official/ecommerce_go/internal/user"
)

func Register(router chi.Router, authMiddleware func(http.Handler) http.Handler, authHandler *auth.Handler, userHandler *userpkg.Handler) {
	router.Route("/api/v1", func(r chi.Router) {
		auth.RegisterRoutes(r, authMiddleware, authHandler)
		userpkg.RegisterRoutes(r, authMiddleware, userHandler)

		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			response.Json(w, http.StatusOK, "Server is Running ", nil)
		})
	})
}

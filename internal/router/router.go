package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/naresh-official/ecommerce_go/internal/address"
	"github.com/naresh-official/ecommerce_go/internal/app"
	"github.com/naresh-official/ecommerce_go/internal/auth"
	"github.com/naresh-official/ecommerce_go/internal/response"
	"github.com/naresh-official/ecommerce_go/internal/user"
)

func Register(router chi.Router, application app.App) {
	router.Route("/api/v1", func(r chi.Router) {
		auth.RegisterRoutes(r, application.Middleware.Auth, application.Handlers.Auth)
		user.RegisterRoutes(r, application.Middleware.Auth, application.Handlers.User)
		address.RegisterRoutes(r, application.Middleware.Auth, application.Middleware.User, application.Handlers.Address)

		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			response.Json(w, http.StatusOK, "Server is Running ", nil)
		})
	})
}

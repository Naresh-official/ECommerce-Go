package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/naresh-official/ecommerce_go/internal/response"
)

func Register(router chi.Router) {
	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			response.Json(w, http.StatusOK, "Server is Running ", nil)
		})
	})

}

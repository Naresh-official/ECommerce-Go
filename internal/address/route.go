package address

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, authMiddleware func(http.Handler) http.Handler, userAuthorizationMiddleware func(http.Handler) http.Handler, handler *Handler) {
	r.
		With(authMiddleware).
		With(userAuthorizationMiddleware).
		Route("/address", func(r chi.Router) {
			r.Get("/", handler.GetAllAddresses)
			r.Post("/", handler.CreateAddress)
			r.Delete("/{addressId}", handler.DeleteAddress)
			r.Get("/{addressId}", handler.GetAddressById)
			r.Put("/{addressId}", handler.UpdateAddress)
			r.Put("/{addressId}/set-default", handler.SetDefaultAddress)
		})
}

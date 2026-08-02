package middleware

import (
	"context"
	"net/http"

	"github.com/naresh-official/ecommerce_go/configs"
	"github.com/naresh-official/ecommerce_go/internal/auth"
	"github.com/naresh-official/ecommerce_go/internal/response"
)

var UserContextKey = "user_id"

func Auth(cfg *configs.JWTConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("access_token")
			if err != nil || cookie.Value == "" {
				response.Error(w, http.StatusUnauthorized, "Unauthenticated")
				return
			}

			claims, err := auth.ValidateAccessToken(cfg, cookie.Value)
			if err != nil {
				response.Error(w, http.StatusUnauthorized, "Unauthenticated")
				return
			}

			ctx := context.WithValue(r.Context(), UserContextKey, claims.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

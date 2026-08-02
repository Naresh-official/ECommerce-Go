package middleware

import (
	"net/http"

	"github.com/naresh-official/ecommerce_go/configs"
	"github.com/naresh-official/ecommerce_go/internal/auth"
	"github.com/naresh-official/ecommerce_go/internal/response"
)

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

			ctx := auth.WithUser(r.Context(), *claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func AuthorizeUser() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			role := auth.GetUserRoleFromContext(r.Context())
			if role != auth.RoleUser {
				response.Error(w, http.StatusForbidden, "Forbidden")
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func AuthorizeAdmin() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			role := auth.GetUserRoleFromContext(r.Context())
			if role != auth.RoleAdmin {
				response.Error(w, http.StatusForbidden, "Forbidden")
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func AuthorizeSeller() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			role := auth.GetUserRoleFromContext(r.Context())
			if role != auth.RoleSeller {
				response.Error(w, http.StatusForbidden, "Forbidden")
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

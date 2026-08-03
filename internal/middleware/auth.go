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

func AuthorizeRoles(allowedRoles ...auth.Role) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			role := auth.GetUserRoleFromContext(r.Context())
			for _, allowedRole := range allowedRoles {
				if role == allowedRole {
					next.ServeHTTP(w, r)
					return
				}
			}

			response.Error(w, http.StatusForbidden, "Forbidden")
		})
	}
}

func AuthorizeUser() func(http.Handler) http.Handler {
	return AuthorizeRoles(auth.RoleUser)
}

func AuthorizeAdmin() func(http.Handler) http.Handler {
	return AuthorizeRoles(auth.RoleAdmin)
}

func AuthorizeSeller() func(http.Handler) http.Handler {
	return AuthorizeRoles(auth.RoleSeller)
}

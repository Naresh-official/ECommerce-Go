package auth

import (
	"context"
)

type userContextKey string

const UserContextKey userContextKey = "user_id"

func WithUser(ctx context.Context, authClaims AccessTokenClaims) context.Context {
	return context.WithValue(ctx, UserContextKey, authClaims)
}

func GetUserIdFromContext(ctx context.Context) string {
	userId, ok := ctx.Value(UserContextKey).(AccessTokenClaims)
	if !ok {
		return ""
	}
	return userId.UserID
}

func GetUserRoleFromContext(ctx context.Context) Role {
	userRole, ok := ctx.Value(UserContextKey).(AccessTokenClaims)
	if !ok {
		return ""
	}
	return Role(userRole.Role)
}

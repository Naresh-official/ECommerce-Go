package auth

type Role string

const (
	RoleUser   Role = "user"
	RoleSeller Role = "seller"
	RoleAdmin  Role = "admin"
)

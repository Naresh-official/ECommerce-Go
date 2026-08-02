package auth

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/naresh-official/ecommerce_go/internal/database/sqlc"
)

type Repository struct {
	q *sqlc.Queries
}

var NoRowsError = pgx.ErrNoRows

func NewRepository(q *sqlc.Queries) *Repository {
	return &Repository{
		q: q,
	}
}

func (r *Repository) GetUserById(ctx context.Context, id uuid.UUID) (*sqlc.User, error) {
	user, err := r.q.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*sqlc.User, error) {
	user, err := r.q.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) CreateUser(ctx context.Context, input SignUpRequest, role string) (*sqlc.User, error) {

	arg := sqlc.CreateUserParams{
		Name:         input.Name,
		Email:        input.Email,
		Phone:        input.Phone,
		PasswordHash: input.Password,
		Role:         sqlc.Role(role),
	}

	user, err := r.q.CreateUser(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) GetUserByEmailAndRole(ctx context.Context, email string, role string) (*sqlc.User, error) {
	user, err := r.q.GetUserByEmailAndRole(ctx, sqlc.GetUserByEmailAndRoleParams{
		Email: email,
		Role:  sqlc.Role(role),
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) UpdateRefreshToken(ctx context.Context, userID uuid.UUID, refreshToken string) error {
	err := r.q.UpdateRefreshToken(ctx, sqlc.UpdateRefreshTokenParams{
		ID:           userID,
		RefreshToken: pgtype.Text{String: refreshToken, Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

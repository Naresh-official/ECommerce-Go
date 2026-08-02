package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/naresh-official/ecommerce_go/internal/database/sqlc"
)

type Repository struct {
	q *sqlc.Queries
}

var NoRowsError = pgx.ErrNoRows

func NewRepository(q *sqlc.Queries) *Repository {
	return &Repository{q: q}
}

func (r *Repository) GetUserByID(ctx context.Context, id uuid.UUID) (*sqlc.User, error) {
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

func (r *Repository) UpdateProfile(ctx context.Context, id uuid.UUID, req UpdateProfileRequest, role string) (*sqlc.User, error) {
	user, err := r.q.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:    id,
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
		Role:  sqlc.Role(role),
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) UpdatePassword(ctx context.Context, id uuid.UUID, passwordHash string) error {
	_, err := r.q.UpdateUserPassword(ctx, sqlc.UpdateUserPasswordParams{
		ID:           id,
		PasswordHash: passwordHash,
	})
	return err
}

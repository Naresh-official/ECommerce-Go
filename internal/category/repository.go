package category

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
	return &Repository{q: q}
}

func (r *Repository) GetAllCategories(ctx context.Context) ([]sqlc.Category, error) {
	return r.q.GetAllCategories(ctx)
}

func (r *Repository) GetCategoryByID(ctx context.Context, categoryID uuid.UUID) (*sqlc.Category, error) {
	category, err := r.q.GetCategoryById(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *Repository) CreateCategory(ctx context.Context, input CreateCategoryRequest) (*sqlc.Category, error) {
	var description pgtype.Text
	if input.Description != "" {
		description = pgtype.Text{String: input.Description, Valid: true}
	}

	category, err := r.q.CreateCategory(ctx, sqlc.CreateCategoryParams{
		Name:        input.Name,
		Description: description,
	})
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *Repository) UpdateCategory(ctx context.Context, categoryID uuid.UUID, input CreateCategoryRequest) (*sqlc.Category, error) {
	var description pgtype.Text
	if input.Description != "" {
		description = pgtype.Text{String: input.Description, Valid: true}
	}

	category, err := r.q.UpdateCategory(ctx, sqlc.UpdateCategoryParams{
		ID:          categoryID,
		Name:        input.Name,
		Description: description,
	})
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *Repository) DeleteCategory(ctx context.Context, categoryID uuid.UUID) error {
	if _, err := r.q.GetCategoryById(ctx, categoryID); err != nil {
		return err
	}

	return r.q.DeleteCategory(ctx, categoryID)
}

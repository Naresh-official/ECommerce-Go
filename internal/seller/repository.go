package seller

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

func (r *Repository) GetSellerByID(ctx context.Context, id uuid.UUID) (*sqlc.Seller, error) {
	seller, err := r.q.GetSellerById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &seller, nil
}

func (r *Repository) GetSellerByOwnerID(ctx context.Context, ownerID uuid.UUID) (*sqlc.Seller, error) {
	seller, err := r.q.GetSellerByOwnerId(ctx, ownerID)
	if err != nil {
		return nil, err
	}
	return &seller, nil
}

func (r *Repository) GetSellerByStoreName(ctx context.Context, storeName string) (*sqlc.Seller, error) {
	seller, err := r.q.GetSellerByStoreName(ctx, storeName)
	if err != nil {
		return nil, err
	}
	return &seller, nil
}

func (r *Repository) CreateSeller(ctx context.Context, ownerID uuid.UUID, input CreateSellerRequest) (*sqlc.Seller, error) {
	var description pgtype.Text
	if input.Description != "" {
		description = pgtype.Text{String: input.Description, Valid: true}
	}

	seller, err := r.q.CreateSeller(ctx, sqlc.CreateSellerParams{
		OwnerID:     ownerID,
		StoreName:   input.StoreName,
		Description: description,
	})
	if err != nil {
		return nil, err
	}
	return &seller, nil
}

func (r *Repository) UpdateSeller(ctx context.Context, sellerID uuid.UUID, input UpdateSellerRequest) (*sqlc.Seller, error) {
	var description pgtype.Text
	if input.Description != "" {
		description = pgtype.Text{String: input.Description, Valid: true}
	}

	seller, err := r.q.UpdateSeller(ctx, sqlc.UpdateSellerParams{
		ID:          sellerID,
		StoreName:   input.StoreName,
		Description: description,
	})
	if err != nil {
		return nil, err
	}
	return &seller, nil
}

func (r *Repository) MarkSellerAsVerified(ctx context.Context, sellerID uuid.UUID) (*sqlc.Seller, error) {
	seller, err := r.q.MarkSellerAsVerified(ctx, sellerID)
	if err != nil {
		return nil, err
	}
	return &seller, nil
}

func (r *Repository) MarkSellerAsUnverified(ctx context.Context, sellerID uuid.UUID) (*sqlc.Seller, error) {
	seller, err := r.q.MarkSellerAsUnverified(ctx, sellerID)
	if err != nil {
		return nil, err
	}
	return &seller, nil
}

func (r *Repository) MarkSellerAsActive(ctx context.Context, sellerID uuid.UUID) (*sqlc.Seller, error) {
	seller, err := r.q.MarkSellerAsActive(ctx, sellerID)
	if err != nil {
		return nil, err
	}
	return &seller, nil
}

func (r *Repository) MarkSellerAsInactive(ctx context.Context, sellerID uuid.UUID) (*sqlc.Seller, error) {
	seller, err := r.q.MarkSellerAsInactive(ctx, sellerID)
	if err != nil {
		return nil, err
	}
	return &seller, nil
}

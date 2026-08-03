package product

import (
	"context"
	"fmt"

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

func (r *Repository) GetProductByID(ctx context.Context, id uuid.UUID) (*sqlc.Product, error) {
	product, err := r.q.GetProductById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *Repository) GetSellerByOwnerID(ctx context.Context, ownerID uuid.UUID) (*sqlc.Seller, error) {
	seller, err := r.q.GetSellerByOwnerId(ctx, ownerID)
	if err != nil {
		return nil, err
	}
	return &seller, nil
}

func (r *Repository) GetProductsBySellerID(ctx context.Context, sellerID uuid.UUID) ([]sqlc.Product, error) {
	return r.q.GetProductsBySellerId(ctx, sellerID)
}

func (r *Repository) CreateProduct(ctx context.Context, sellerID uuid.UUID, input CreateProductRequest) (*sqlc.Product, error) {
	var categoryID pgtype.UUID
	if input.CategoryID != "" {
		parsedCategoryID, err := uuid.Parse(input.CategoryID)
		if err != nil {
			return nil, err
		}
		categoryID = pgtype.UUID{Bytes: parsedCategoryID, Valid: true}
	}

	var description pgtype.Text
	if input.Description != "" {
		description = pgtype.Text{String: input.Description, Valid: true}
	}

	var price pgtype.Numeric
	if err := price.Scan(input.Price); err != nil {
		return nil, fmt.Errorf("invalid price: %w", err)
	}

	product, err := r.q.CreateProduct(ctx, sqlc.CreateProductParams{
		SellerID:      sellerID,
		CategoryID:    categoryID,
		Name:          input.Name,
		Description:   description,
		Price:         price,
		StockQuantity: input.StockQuantity,
		Images:        input.Images,
	})
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *Repository) UpdateProduct(ctx context.Context, productID uuid.UUID, input UpdateProductRequest) (*sqlc.Product, error) {
	var categoryID pgtype.UUID
	if input.CategoryID != "" {
		parsedCategoryID, err := uuid.Parse(input.CategoryID)
		if err != nil {
			return nil, err
		}
		categoryID = pgtype.UUID{Bytes: parsedCategoryID, Valid: true}
	}

	var description pgtype.Text
	if input.Description != "" {
		description = pgtype.Text{String: input.Description, Valid: true}
	}

	product, err := r.q.UpdateProduct(ctx, sqlc.UpdateProductParams{
		ID:          productID,
		CategoryID:  categoryID,
		Name:        input.Name,
		Description: description,
		Images:      input.Images,
	})
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *Repository) UpdateProductPrice(ctx context.Context, productID uuid.UUID, price string) (*sqlc.Product, error) {
	var numericPrice pgtype.Numeric
	if err := numericPrice.Scan(price); err != nil {
		return nil, fmt.Errorf("invalid price: %w", err)
	}

	product, err := r.q.UpdateProductPrice(ctx, sqlc.UpdateProductPriceParams{
		ID:    productID,
		Price: numericPrice,
	})
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *Repository) UpdateProductStockQuantity(ctx context.Context, productID uuid.UUID, stockQuantity int32) (*sqlc.Product, error) {
	product, err := r.q.UpdateProductStockQuantity(ctx, sqlc.UpdateProductStockQuantityParams{
		ID:            productID,
		StockQuantity: stockQuantity,
	})
	if err != nil {
		return nil, err
	}
	return &product, nil
}

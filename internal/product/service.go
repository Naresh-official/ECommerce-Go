package product

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/naresh-official/ecommerce_go/internal/auth"
	"github.com/naresh-official/ecommerce_go/internal/database/sqlc"
)

var ErrProductNotFound = errors.New("Product not found")

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetMyProducts(ctx context.Context) ([]ProductResponse, error) {
	sellerID, err := s.currentSellerID(ctx)
	if err != nil {
		return nil, err
	}

	products, err := s.repo.GetProductsBySellerID(ctx, sellerID)
	if err != nil {
		return nil, err
	}

	responses := make([]ProductResponse, 0, len(products))
	for _, product := range products {
		responses = append(responses, mapProduct(product))
	}

	return responses, nil
}

func (s *Service) GetProductByID(ctx context.Context, productID string) (*ProductResponse, error) {
	product, err := s.repo.GetProductByID(ctx, uuid.MustParse(productID))
	if err != nil {
		if errors.Is(err, NoRowsError) {
			return nil, ErrProductNotFound
		}
		return nil, err
	}

	response := mapProduct(*product)
	return &response, nil
}

func (s *Service) CreateProduct(ctx context.Context, input CreateProductRequest) (*ProductResponse, error) {
	sellerID, err := s.currentSellerID(ctx)
	if err != nil {
		return nil, err
	}

	product, err := s.repo.CreateProduct(ctx, sellerID, input)
	if err != nil {
		return nil, err
	}

	response := mapProduct(*product)
	return &response, nil
}

func (s *Service) UpdateProduct(ctx context.Context, productID string, input UpdateProductRequest) (*ProductResponse, error) {
	updatedProduct, err := s.repo.UpdateProduct(ctx, uuid.MustParse(productID), input)
	if err != nil {
		if errors.Is(err, NoRowsError) {
			return nil, ErrProductNotFound
		}
		return nil, err
	}

	response := mapProduct(*updatedProduct)
	return &response, nil
}

func (s *Service) UpdateProductPrice(ctx context.Context, productID string, price string) (*ProductResponse, error) {
	updatedProduct, err := s.repo.UpdateProductPrice(ctx, uuid.MustParse(productID), price)
	if err != nil {
		if errors.Is(err, NoRowsError) {
			return nil, ErrProductNotFound
		}
		return nil, err
	}

	response := mapProduct(*updatedProduct)
	return &response, nil
}

func (s *Service) UpdateProductStockQuantity(ctx context.Context, productID string, stockQuantity int32) (*ProductResponse, error) {
	updatedProduct, err := s.repo.UpdateProductStockQuantity(ctx, uuid.MustParse(productID), stockQuantity)
	if err != nil {
		if errors.Is(err, NoRowsError) {
			return nil, ErrProductNotFound
		}
		return nil, err
	}

	response := mapProduct(*updatedProduct)
	return &response, nil
}

func (s *Service) currentSellerID(ctx context.Context) (uuid.UUID, error) {
	seller, err := s.repo.GetSellerByOwnerID(ctx, uuid.MustParse(auth.GetUserIdFromContext(ctx)))
	if err != nil {
		if errors.Is(err, NoRowsError) {
			return uuid.Nil, fmt.Errorf("seller profile not found")
		}
		return uuid.Nil, err
	}

	return seller.ID, nil
}

func mapProduct(product sqlc.Product) ProductResponse {
	return ProductResponse{
		ID:            product.ID.String(),
		SellerID:      product.SellerID.String(),
		CategoryID:    uuidFromPgtype(product.CategoryID),
		Name:          product.Name,
		Description:   product.Description.String,
		Price:         numericToString(product.Price),
		StockQuantity: product.StockQuantity,
		Images:        product.Images,
		CreatedAt:     formatTimestamp(product.CreatedAt),
		UpdatedAt:     formatTimestamp(product.UpdatedAt),
	}
}

func formatTimestamp(timestamp pgtype.Timestamp) string {
	if !timestamp.Valid {
		return ""
	}
	return timestamp.Time.Format(time.RFC3339)
}

func uuidFromPgtype(value pgtype.UUID) string {
	if !value.Valid {
		return ""
	}
	return uuid.UUID(value.Bytes).String()
}

func numericToString(value pgtype.Numeric) string {
	result, err := value.Value()
	if err != nil || result == nil {
		return ""
	}
	return fmt.Sprint(result)
}

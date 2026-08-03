package seller

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/naresh-official/ecommerce_go/internal/auth"
	"github.com/naresh-official/ecommerce_go/internal/database/sqlc"
)

var (
	ErrSellerNotFound        = errors.New("Seller not found")
	ErrSellerAlreadyExists   = errors.New("Seller already exists")
	ErrStoreNameAlreadyTaken = errors.New("Store name already exists")
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetMySeller(ctx context.Context) (*SellerResponse, error) {
	seller, err := s.repo.GetSellerByOwnerID(ctx, uuid.MustParse(auth.GetUserIdFromContext(ctx)))
	if err != nil {
		if errors.Is(err, NoRowsError) {
			return nil, ErrSellerNotFound
		}
		return nil, err
	}

	return mapSeller(*seller), nil
}

func (s *Service) GetSellerByID(ctx context.Context, sellerID string) (*SellerResponse, error) {
	seller, err := s.repo.GetSellerByID(ctx, uuid.MustParse(sellerID))
	if err != nil {
		if errors.Is(err, NoRowsError) {
			return nil, ErrSellerNotFound
		}
		return nil, err
	}

	return mapSeller(*seller), nil
}

func (s *Service) CreateSeller(ctx context.Context, input CreateSellerRequest) (*SellerResponse, error) {
	ownerID := uuid.MustParse(auth.GetUserIdFromContext(ctx))

	if _, err := s.repo.GetSellerByOwnerID(ctx, ownerID); err != nil {
		if !errors.Is(err, NoRowsError) {
			return nil, err
		}
	} else {
		return nil, ErrSellerAlreadyExists
	}

	if _, err := s.repo.GetSellerByStoreName(ctx, input.StoreName); err != nil {
		if !errors.Is(err, NoRowsError) {
			return nil, err
		}
	} else {
		return nil, ErrStoreNameAlreadyTaken
	}

	seller, err := s.repo.CreateSeller(ctx, ownerID, input)
	if err != nil {
		return nil, err
	}

	return mapSeller(*seller), nil
}

func (s *Service) UpdateSeller(ctx context.Context, input UpdateSellerRequest) (*SellerResponse, error) {
	ownerID := uuid.MustParse(auth.GetUserIdFromContext(ctx))
	seller, err := s.repo.GetSellerByOwnerID(ctx, ownerID)
	if err != nil {
		if errors.Is(err, NoRowsError) {
			return nil, ErrSellerNotFound
		}
		return nil, err
	}

	updatedSeller, err := s.repo.UpdateSeller(ctx, seller.ID, input)
	if err != nil {
		return nil, err
	}

	return mapSeller(*updatedSeller), nil
}

func (s *Service) VerifySeller(ctx context.Context, sellerID string) (*SellerResponse, error) {
	seller, err := s.repo.MarkSellerAsVerified(ctx, uuid.MustParse(sellerID))
	if err != nil {
		if errors.Is(err, NoRowsError) {
			return nil, ErrSellerNotFound
		}
		return nil, err
	}

	return mapSeller(*seller), nil
}

func (s *Service) UnverifySeller(ctx context.Context, sellerID string) (*SellerResponse, error) {
	seller, err := s.repo.MarkSellerAsUnverified(ctx, uuid.MustParse(sellerID))
	if err != nil {
		if errors.Is(err, NoRowsError) {
			return nil, ErrSellerNotFound
		}
		return nil, err
	}

	return mapSeller(*seller), nil
}

func (s *Service) ActivateSeller(ctx context.Context, sellerID string) (*SellerResponse, error) {
	seller, err := s.repo.MarkSellerAsActive(ctx, uuid.MustParse(sellerID))
	if err != nil {
		if errors.Is(err, NoRowsError) {
			return nil, ErrSellerNotFound
		}
		return nil, err
	}

	return mapSeller(*seller), nil
}

func (s *Service) DeactivateSeller(ctx context.Context, sellerID string) (*SellerResponse, error) {
	seller, err := s.repo.MarkSellerAsInactive(ctx, uuid.MustParse(sellerID))
	if err != nil {
		if errors.Is(err, NoRowsError) {
			return nil, ErrSellerNotFound
		}
		return nil, err
	}

	return mapSeller(*seller), nil
}

func mapSeller(seller sqlc.Seller) *SellerResponse {
	return &SellerResponse{
		ID:          seller.ID.String(),
		OwnerID:     seller.OwnerID.String(),
		StoreName:   seller.StoreName,
		Description: seller.Description.String,
		IsActive:    seller.IsActive,
		IsVerified:  seller.IsVerified,
		CreatedAt:   formatTimestamp(seller.CreatedAt),
		UpdatedAt:   formatTimestamp(seller.UpdatedAt),
	}
}

func formatTimestamp(timestamp pgtype.Timestamp) string {
	if !timestamp.Valid {
		return ""
	}
	return timestamp.Time.Format(time.RFC3339)
}

package address

import (
	"context"

	"github.com/google/uuid"
	"github.com/naresh-official/ecommerce_go/internal/auth"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetAllAddresses(ctx context.Context) ([]GetAddressResponse, error) {
	addresses, err := s.repo.GetAddressesByUserId(ctx, uuid.MustParse(auth.GetUserIdFromContext(ctx)))
	if err != nil {
		return nil, err
	}
	var responses []GetAddressResponse
	for _, addr := range addresses {
		responses = append(responses, GetAddressResponse{
			ID:           addr.ID.String(),
			AddressLine1: addr.AddressLine1,
			AddressLine2: addr.AddressLine2.String,
			City:         addr.City,
			State:        addr.State,
			Country:      addr.Country,
			PostalCode:   addr.PostalCode,
			Phone:        addr.Phone,
			IsDefault:    addr.IsDefault,
		})
	}
	return responses, nil
}

func (s *Service) CreateAddress(ctx context.Context, input CreateAddressRequest) (*GetAddressResponse, error) {
	userId := uuid.MustParse(auth.GetUserIdFromContext(ctx))
	address, err := s.repo.CreateAddress(ctx, input, userId)
	if err != nil {
		return nil, err
	}
	response := &GetAddressResponse{
		ID:           address.ID.String(),
		AddressLine1: address.AddressLine1,
		AddressLine2: address.AddressLine2.String,
		City:         address.City,
		State:        address.State,
		Country:      address.Country,
		PostalCode:   address.PostalCode,
		Phone:        address.Phone,
		IsDefault:    address.IsDefault,
	}
	return response, nil
}

func (s *Service) DeleteAddress(ctx context.Context, addressId string) error {
	id, err := uuid.Parse(addressId)
	if err != nil {
		return err
	}
	err = s.repo.DeleteAddress(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAddressById(ctx context.Context, addressId string) (*GetAddressResponse, error) {
	id, err := uuid.Parse(addressId)
	if err != nil {
		return nil, err
	}
	address, err := s.repo.GetAddressById(ctx, id)
	if err != nil {
		return nil, err
	}
	response := &GetAddressResponse{
		ID:           address.ID.String(),
		AddressLine1: address.AddressLine1,
		AddressLine2: address.AddressLine2.String,
		City:         address.City,
		State:        address.State,
		Country:      address.Country,
		PostalCode:   address.PostalCode,
		Phone:        address.Phone,
		IsDefault:    address.IsDefault,
	}
	return response, nil
}

func (s *Service) UpdateAddress(ctx context.Context, addressId string, input CreateAddressRequest) (*GetAddressResponse, error) {
	id, err := uuid.Parse(addressId)
	if err != nil {
		return nil, err
	}
	address, err := s.repo.UpdateAddress(ctx, id, input)
	if err != nil {
		return nil, err
	}
	response := &GetAddressResponse{
		ID:           address.ID.String(),
		AddressLine1: address.AddressLine1,
		AddressLine2: address.AddressLine2.String,
		City:         address.City,
		State:        address.State,
		Country:      address.Country,
		PostalCode:   address.PostalCode,
		Phone:        address.Phone,
		IsDefault:    address.IsDefault,
	}
	return response, nil
}

func (s *Service) SetDefaultAddress(ctx context.Context, addressId string) error {
	id, err := uuid.Parse(addressId)
	if err != nil {
		return err
	}
	err = s.repo.SetDefaultAddress(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

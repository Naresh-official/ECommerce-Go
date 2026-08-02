package address

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/naresh-official/ecommerce_go/internal/database/sqlc"
)

type Repository struct {
	q *sqlc.Queries
}

func NewRepository(q *sqlc.Queries) *Repository {
	return &Repository{
		q: q,
	}
}

func (r *Repository) CreateAddress(ctx context.Context, input CreateAddressRequest, userId uuid.UUID) (*sqlc.Address, error) {
	var addressLine2 pgtype.Text
	if input.AddressLine2 != "" {
		addressLine2 = pgtype.Text{String: input.AddressLine2, Valid: true}
	}

	arg := sqlc.CreateAddressParams{
		UserID:       userId,
		AddressLine1: input.AddressLine1,
		AddressLine2: addressLine2,
		City:         input.City,
		State:        input.State,
		Country:      input.Country,
		PostalCode:   input.PostalCode,
		Phone:        input.Phone,
	}

	address, err := r.q.CreateAddress(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func (r *Repository) DeleteAddress(ctx context.Context, addressId uuid.UUID) error {
	err := r.q.DeleteAddress(ctx, addressId)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetAddressById(ctx context.Context, addressId uuid.UUID) (*sqlc.Address, error) {
	address, err := r.q.GetAddressById(ctx, addressId)
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func (r *Repository) GetAddressesByUserId(ctx context.Context, userId uuid.UUID) ([]sqlc.Address, error) {
	addresses, err := r.q.GetAllAddressesOfUser(ctx, userId)
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

func (r *Repository) SetDefaultAddress(ctx context.Context, addressId uuid.UUID) error {
	err := r.q.SetDefaultAddress(ctx, addressId)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateAddress(ctx context.Context, addressId uuid.UUID, input CreateAddressRequest) (*sqlc.Address, error) {
	var addressLine2 pgtype.Text
	if input.AddressLine2 != "" {
		addressLine2 = pgtype.Text{String: input.AddressLine2, Valid: true}
	}

	arg := sqlc.UpdateAddressParams{
		ID:           addressId,
		AddressLine1: input.AddressLine1,
		AddressLine2: addressLine2,
		City:         input.City,
		State:        input.State,
		Country:      input.Country,
		PostalCode:   input.PostalCode,
		Phone:        input.Phone,
	}

	address, err := r.q.UpdateAddress(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &address, nil
}

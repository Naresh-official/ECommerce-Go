package address

type CreateAddressRequest struct {
	AddressLine1 string `json:"address_line1" validate:"required"`
	AddressLine2 string `json:"address_line2"`
	City         string `json:"city" validate:"required"`
	State        string `json:"state" validate:"required"`
	PostalCode   string `json:"postal_code" validate:"required,len=6"`
	Country      string `json:"country" validate:"required"`
	Phone        string `json:"phone" validate:"required,len=10"`
}

type GetAddressResponse struct {
	ID           string `json:"id"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2,omitempty"`
	City         string `json:"city"`
	State        string `json:"state"`
	PostalCode   string `json:"postal_code"`
	Country      string `json:"country"`
	Phone        string `json:"phone"`
	IsDefault    bool   `json:"is_default"`
}

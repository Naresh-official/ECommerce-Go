package seller

type CreateSellerRequest struct {
	StoreName   string `json:"store_name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateSellerRequest struct {
	StoreName   string `json:"store_name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type SellerResponse struct {
	ID          string `json:"id"`
	OwnerID     string `json:"owner_id"`
	StoreName   string `json:"store_name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
	IsVerified  bool   `json:"is_verified"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

package product

type CreateProductRequest struct {
	CategoryID    string   `json:"category_id"`
	Name          string   `json:"name" validate:"required"`
	Description   string   `json:"description" validate:"required"`
	Price         string   `json:"price" validate:"required"`
	StockQuantity int32    `json:"stock_quantity" validate:"required"`
	Images        []string `json:"images" validate:"required"`
}

type UpdateProductRequest struct {
	CategoryID  string   `json:"category_id"`
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Images      []string `json:"images" validate:"required"`
}

type UpdateProductPriceRequest struct {
	Price string `json:"price" validate:"required"`
}

type UpdateProductStockQuantityRequest struct {
	StockQuantity int32 `json:"stock_quantity" validate:"required"`
}

type ProductResponse struct {
	ID            string   `json:"id"`
	SellerID      string   `json:"seller_id"`
	CategoryID    string   `json:"category_id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Price         string   `json:"price"`
	StockQuantity int32    `json:"stock_quantity"`
	Images        []string `json:"images"`
	CreatedAt     string   `json:"created_at"`
	UpdatedAt     string   `json:"updated_at"`
}

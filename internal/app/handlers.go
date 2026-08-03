package app

import (
	"github.com/naresh-official/ecommerce_go/internal/address"
	"github.com/naresh-official/ecommerce_go/internal/auth"
	"github.com/naresh-official/ecommerce_go/internal/category"
	"github.com/naresh-official/ecommerce_go/internal/product"
	"github.com/naresh-official/ecommerce_go/internal/seller"
	"github.com/naresh-official/ecommerce_go/internal/user"
)

type Handlers struct {
	Auth     *auth.Handler
	User     *user.Handler
	Address  *address.Handler
	Category *category.Handler
	Seller   *seller.Handler
	Product  *product.Handler
}

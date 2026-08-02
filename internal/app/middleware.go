package app

import "net/http"

type Middlewares struct {
	Auth   func(http.Handler) http.Handler
	User   func(http.Handler) http.Handler
	Admin  func(http.Handler) http.Handler
	Seller func(http.Handler) http.Handler
}

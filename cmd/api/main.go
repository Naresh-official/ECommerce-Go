package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/naresh-official/ecommerce_go/configs"
	"github.com/naresh-official/ecommerce_go/internal/address"
	"github.com/naresh-official/ecommerce_go/internal/app"
	"github.com/naresh-official/ecommerce_go/internal/auth"
	"github.com/naresh-official/ecommerce_go/internal/category"
	"github.com/naresh-official/ecommerce_go/internal/database"
	"github.com/naresh-official/ecommerce_go/internal/database/sqlc"
	appmiddleware "github.com/naresh-official/ecommerce_go/internal/middleware"
	"github.com/naresh-official/ecommerce_go/internal/router"
	"github.com/naresh-official/ecommerce_go/internal/user"
)

func main() {
	configs.LoadEnv()

	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatal("Error in Loading Config ", err)
	}

	err = database.ConnectDB(cfg.Database)
	if err != nil {
		log.Fatal("Error in Connecting to Database ", err)
	}

	r := chi.NewRouter()

	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)

	queries := sqlc.New(database.DB)

	// Initialize repositories, services, and handlers for each module

	authRepository := auth.NewRepository(queries)
	authService := auth.NewService(authRepository, &cfg.JWT)
	authHandler := auth.NewHandler(authService, &cfg.App)

	userRepository := user.NewRepository(queries)
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	addressRepository := address.NewRepository(queries)
	addressService := address.NewService(addressRepository)
	addressHandler := address.NewHandler(addressService)

	categoryRepository := category.NewRepository(queries)
	categoryService := category.NewService(categoryRepository)
	categoryHandler := category.NewHandler(categoryService)

	// Initialize the application with handlers and middleware

	application := app.App{
		Handlers: app.Handlers{
			Auth:     authHandler,
			User:     userHandler,
			Address:  addressHandler,
			Category: categoryHandler,
		},
		Middleware: app.Middlewares{
			Auth:   appmiddleware.Auth(&cfg.JWT),
			User:   appmiddleware.AuthorizeUser(),
			Admin:  appmiddleware.AuthorizeAdmin(),
			Seller: appmiddleware.AuthorizeSeller(),
		},
	}

	// Register the routes with the router
	router.Register(r, application)

	slog.Info("Server started at port " + cfg.Server.Port)
	http.ListenAndServe(":"+cfg.Server.Port, r)
}

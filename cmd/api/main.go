package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/naresh-official/ecommerce_go/configs"
	"github.com/naresh-official/ecommerce_go/internal/auth"
	"github.com/naresh-official/ecommerce_go/internal/database"
	"github.com/naresh-official/ecommerce_go/internal/database/sqlc"
	appmiddleware "github.com/naresh-official/ecommerce_go/internal/middleware"
	"github.com/naresh-official/ecommerce_go/internal/router"
	userpkg "github.com/naresh-official/ecommerce_go/internal/user"
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

	authRepository := auth.NewRepository(queries)
	authService := auth.NewService(authRepository, &cfg.JWT)
	authHandler := auth.NewHandler(authService, &cfg.App)

	userRepository := userpkg.NewRepository(queries)
	userService := userpkg.NewService(userRepository)
	userHandler := userpkg.NewHandler(userService)

	router.Register(r, appmiddleware.Auth(&cfg.JWT), authHandler, userHandler)

	slog.Info("Server started at port " + cfg.Server.Port)
	http.ListenAndServe(":"+cfg.Server.Port, r)
}

package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/naresh-official/ecommerce_go/configs"
	"github.com/naresh-official/ecommerce_go/internal/database"
	"github.com/naresh-official/ecommerce_go/internal/database/sqlc"
	"github.com/naresh-official/ecommerce_go/internal/router"
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

	queries := sqlc.New(database.DB)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	router.Register(r)

	slog.Info("Server started at port " + cfg.Server.Port)
	http.ListenAndServe(":"+cfg.Server.Port, r)
}

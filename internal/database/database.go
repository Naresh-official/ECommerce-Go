package database

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/naresh-official/ecommerce_go/configs"
)

var DB *pgxpool.Pool

func ConnectDB(cfg configs.DatabaseConfig) error {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, cfg.DatabaseUrl)

	if err != nil {
		return err
	}

	err = pool.Ping(ctx)
	if err != nil {
		return err
	}

	slog.Info("Database connected successfully")
	DB = pool
	return nil
}

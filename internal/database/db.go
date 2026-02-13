package database

import (
	"context"
	"log/slog"
	"time"

	pgx "github.com/jackc/pgx/v5/pgxpool"
)

func CreateDBPool(ctx context.Context, connString string) (*pgx.Pool, error) {

	pool, err := pgx.ParseConfig(connString)
	if err != nil {
		slog.Info("Error connecting to Database", "db", connString)
		return nil, err
	}

	pool.MinConns = 2
	pool.MaxConns = 10
	pool.MaxConnLifetime = 1 * time.Hour
	pool.MaxConnIdleTime = 15 * time.Minute
	pool.HealthCheckPeriod = 1 * time.Minute
	pool.ConnConfig.ConnectTimeout = 5 * time.Second

	return pgx.NewWithConfig(ctx, pool)
}

package handler

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type ApiHandler struct {
	Store *pgxpool.Pool
}

package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/freecsdon30/url-shortener/internal/api"
	c "github.com/freecsdon30/url-shortener/internal/config"
	"github.com/freecsdon30/url-shortener/internal/database"
	"github.com/freecsdon30/url-shortener/internal/storage"
)

func main() {
	ctx := context.Background()
	c.LoadAllConfig()

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DB.Username, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.Name)

	db, err := database.CreateDBPool(ctx, url)
	if err != nil {
		slog.Error("Error connecting to db", "error", err)
	}

	store := storage.NewStorage(db)
	api.Server(&store)
}

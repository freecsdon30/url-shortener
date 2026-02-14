package api

import (
	"log/slog"

	"github.com/freecsdon30/url-shortener/internal/api/handler"
	"github.com/freecsdon30/url-shortener/internal/storage"
	"github.com/labstack/echo/v4"
)

func Server(store *storage.Storage) error {
	e := echo.New()
	h := handler.ApiHandler{Store: store.DB}

	e.GET("/:id", h.GetShortenURL)
	e.POST("/url", h.ShortenURL)

	if err := e.Start(":8080"); err != nil {
		slog.Error("Error starting server", "error", err)
		return err
	}
	return nil
}

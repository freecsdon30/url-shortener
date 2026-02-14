package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/freecsdon30/url-shortener/internal/api/dto"
	"github.com/freecsdon30/url-shortener/internal/shortener"
	"github.com/labstack/echo/v4"
)

func (h *ApiHandler) GetShortenURL(c echo.Context) error {
	id := c.Param("id")

	us := shortener.UrlStore{Store: h.Store}
	longURL, err := us.FetchLongURL(id)
	if err != nil {
		slog.Error("URL_FETCH_ERR", "error", err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.Redirect(http.StatusMovedPermanently, longURL)
}

func (h *ApiHandler) ShortenURL(c echo.Context) error {
	var req dto.ShortenURLRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	us := shortener.UrlStore{Store: h.Store}
	shortURLKey, err := us.ShortenURL(req.LongURl)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"short_url": fmt.Sprintf("http://localhost:8080/%s", shortURLKey),
	})
}

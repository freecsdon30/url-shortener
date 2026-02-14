package shortener

import (
	"math/rand"
	"strings"

	"github.com/freecsdon30/url-shortener/internal/storage"
	"github.com/jackc/pgx/v5/pgxpool"
)

const CHAR_SET = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type UrlStore struct {
	Store *pgxpool.Pool
}

func GenerateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = CHAR_SET[rand.Intn(len(CHAR_SET))]
	}
	return string(b)
}

func (u *UrlStore) ShortenURL(longURL string) (string, error) {
	encodedURL := GenerateRandomString(7)
	cleanedLongURl := strings.TrimSpace(longURL)

	us := storage.UrlStore{Store: u.Store}
	us.InsertURL(encodedURL, cleanedLongURl)
	return encodedURL, nil
}

func (u *UrlStore) FetchLongURL(key string) (string, error) {

	us := storage.UrlStore{Store: u.Store}
	longURL, err := us.FetchLongURL(key)
	if err != nil {
		return "", err
	}
	return longURL, nil
}

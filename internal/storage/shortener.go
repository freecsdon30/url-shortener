package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UrlStore struct {
	Store *pgxpool.Pool
}

func (s *UrlStore) InsertURL(shortURl string, longURL string) error {
	ctx := context.Background()
	sql := `
		INSERT INTO public.url_lookup(short_url,long_url) VALUES ($1,$2);
	`
	_, err := s.Store.Exec(ctx, sql, shortURl, longURL)
	if err != nil {
		return err
	}
	return nil
}

func (s *UrlStore) FetchLongURL(shorURlKey string) (string, error) {
	ctx := context.Background()
	sql := `
	    SELECT long_url FROM public.url_lookup where short_url = $1;
	`
	row := s.Store.QueryRow(ctx, sql, shorURlKey)

	var longURl string
	if err := row.Scan(&longURl); err != nil {
		return "", err
	}
	return longURl, nil
}

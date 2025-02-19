package storage

import (
	"context"

	"github.com/Njrctr/gw-exchanger/internal/storage/postgres"
	"github.com/jmoiron/sqlx"
)

type Currency interface {
	GetAllRates(ctx context.Context) (map[string]float64, error)
	GetRate(ctx context.Context, from, to string) (float64, error)
}

type Repository struct {
	Currency
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Currency: postgres.NewCurrencyRepo(db),
	}
}

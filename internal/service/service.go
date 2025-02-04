package service

import (
	"context"

	"github.com/Njrctr/gw-exchanger/internal/storage"
)

type Currency interface {
	GetAllRates(ctx context.Context) (map[string]float64, error)
	GetCurrency(ctx context.Context, key string) (float64, error)
}

type Service struct {
	Currency
}

func NewService(repos *storage.Repository) *Service {
	return &Service{
		Currency: NewCurrencyService(repos),
	}
}

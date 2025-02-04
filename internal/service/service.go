package service

import (
	"context"

	"github.com/Njrctr/gw-exchanger/internal/storage"
)

type Currency interface {
	GetAllRates(ctx context.Context) (map[string]float64, error)
	GetRate(ctx context.Context, from, to string) (float64, error)
}

type Service struct {
	Currency
}

func NewService(repos *storage.Repository) *Service {
	return &Service{
		Currency: NewCurrencyService(repos),
	}
}

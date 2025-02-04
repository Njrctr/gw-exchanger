package service

import (
	"context"

	"github.com/Njrctr/gw-exchanger/internal/storage"
)

type CurrencyService struct {
	repo storage.Currency
}

func NewCurrencyService(repo storage.Currency) *CurrencyService {
	return &CurrencyService{repo: repo}
}

func (s *CurrencyService) GetRate(ctx context.Context, from, to string) (float64, error) {
	return s.repo.GetRate(ctx, from, to)
}

func (s *CurrencyService) GetAllRates(ctx context.Context) (map[string]float64, error) {
	return s.repo.GetAllRates(ctx)
}

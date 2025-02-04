package exchange

import (
	"context"

	"github.com/Njrctr/gw-exchanger/internal/service"
	exchangev1 "github.com/Njrctr/gw-proto-exchange/gen/go/exchange"
	"google.golang.org/grpc"
)

type serverAPI struct {
	exchangev1.UnimplementedExchangeServiceServer
	services *service.Service
}

func Register(gRPC *grpc.Server, services *service.Service) {
	exchangev1.RegisterExchangeServiceServer(gRPC, &serverAPI{services: services})
}

func (s *serverAPI) GetExchangeRates(
	ctx context.Context,
	req *exchangev1.Empty,
) (*exchangev1.ExchangeRatesResponse, error) {
	allRates, _ := s.services.GetAllRates(ctx)
	return &exchangev1.ExchangeRatesResponse{
		Rates: allRates,
	}, nil
}

func (s *serverAPI) GetExchangeRateForCurrency(
	ctx context.Context,
	req *exchangev1.CurrencyRequest,
) (*exchangev1.ExchangeRateResponse, error) {
	panic(123)
}

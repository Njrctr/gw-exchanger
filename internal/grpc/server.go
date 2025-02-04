package exchange

import (
	"context"
	"strings"

	"github.com/Njrctr/gw-exchanger/internal/service"
	exchangev1 "github.com/Njrctr/gw-proto-exchange/gen/go/exchange"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	req.FromCurrency, req.ToCurrency = strings.ToUpper(req.FromCurrency), strings.ToUpper(req.ToCurrency)
	rate, err := s.services.GetRate(ctx, req.FromCurrency, req.ToCurrency)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &exchangev1.ExchangeRateResponse{
		FromCurrency: req.FromCurrency,
		ToCurrency:   req.ToCurrency,
		Rate:         rate,
	}, nil
}

package app

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/Njrctr/gw-exchanger/internal/config"
	exchange "github.com/Njrctr/gw-exchanger/internal/grpc"
	"github.com/Njrctr/gw-exchanger/internal/service"
	"github.com/Njrctr/gw-exchanger/internal/storage"
	"github.com/Njrctr/gw-exchanger/internal/storage/postgres"
	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       string
}

func NewApp(log *slog.Logger, cfg *config.Config) *App {
	gRPCServer := grpc.NewServer()
	fmt.Println(cfg.DB)
	db, err := postgres.NewDB(cfg.DB)
	if err != nil {
		panic(err)
	}
	repos := storage.NewRepository(db)
	service := service.NewService(repos)
	exchange.Register(gRPCServer, service)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       cfg.GRPC.Port,
	}
}

func (a *App) Run() error {
	const op = "app.Run"
	log := a.log.With(
		slog.String("op", op),
		slog.String("port", a.port),
	)

	l, err := net.Listen("tcp", ":"+a.port)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("gRPC server is running", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Stop() {
	const op = "app.Stop"
	a.log.With(slog.String("op", op)).Info("stoping gRPC server", slog.String("port", a.port))

	a.gRPCServer.GracefulStop()
}

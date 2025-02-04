package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	app "github.com/Njrctr/gw-exchanger/internal/app/grpc"
	"github.com/Njrctr/gw-exchanger/internal/config"
	"github.com/Njrctr/gw-exchanger/pkg/logger/slogpretty"
)

const (
	envDev   = "dev"
	envProd  = "prod"
	envLocal = "local"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		panic(err.Error())
	}

	log := setupLogger(cfg.Env)
	application := app.NewApp(log, cfg)
	go application.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	application.Stop()
	log.Info("Stoping application")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envDev:
		log = slog.New(slog.NewJSONHandler(
			os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug},
		))
	case envProd:
		log = slog.New(slog.NewJSONHandler(
			os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo},
		))
	case envLocal:
		log = slogpretty.SetupPrettySlog(slog.LevelDebug)
	}

	return log
}

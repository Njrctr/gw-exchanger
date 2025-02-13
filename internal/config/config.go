package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type Config struct {
	Env  string
	GRPC GRPC
	DB   DBConfig
}

type GRPC struct {
	Port string
	// Timeout time.Duration
}

func MustLoad() (*Config, error) {
	mode := flag.String("mode", "debug", "")
	flag.Parse()
	confFile := "config_dev.env"
	env := "local"
	if *mode != "debug" && *mode != "release" {
		return nil, fmt.Errorf("неверный режим запуска: %s", *mode)
	}
	if *mode == "release" {
		confFile = "config_relese.env"
	}

	if err := godotenv.Load(confFile); err != nil {
		return nil, fmt.Errorf("ошибка получения переменных окружения: %s", err.Error())
	}

	config := Config{
		Env: env,
		GRPC: GRPC{
			Port: os.Getenv("APP_PORT"),
		},
		DB: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			DBName:   os.Getenv("DB_DBNAME"),
			Username: os.Getenv("DB_USERNAME"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
			Password: os.Getenv("DB_PASSWORD"),
		},
	}
	return &config, nil
}

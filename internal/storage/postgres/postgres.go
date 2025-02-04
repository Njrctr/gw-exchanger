package postgres

import (
	"fmt"

	"github.com/Njrctr/gw-exchanger/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var CurrencyTable = "currency_rates"
var BaseCurrency = "RUB"

func NewDB(cfg config.DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

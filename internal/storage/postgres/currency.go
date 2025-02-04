package postgres

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type CurrencyRepo struct {
	db *sqlx.DB
}

func NewCurrencyRepo(db *sqlx.DB) *CurrencyRepo {
	return &CurrencyRepo{db: db}
}

func (r *CurrencyRepo) GetRate(ctx context.Context, from, to string) (float64, error) {
	query := fmt.Sprintf("SELECT rate FROM %s WHERE from_currency=$1 AND to_currency=$2", CurrencyTable)
	var rate float64
	err := r.db.Get(&rate, query,
		from, to)

	return rate, err
}

func (r *CurrencyRepo) GetAllRates(ctx context.Context) (map[string]float64, error) {
	resultMap := map[string]float64{}
	query := fmt.Sprintf("SELECT to_currency, rate FROM %s WHERE from_currency=$1", CurrencyTable)

	rows, err := r.db.Query(query, BaseCurrency)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for rows.Next() {
		var currency string
		var value float64
		if err := rows.Scan(&currency, &value); err != nil {
			return nil, err
		}
		resultMap[currency] = value
	}

	return resultMap, nil
}

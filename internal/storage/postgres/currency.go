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

func (r *CurrencyRepo) GetCurrency(ctx context.Context, key string) (float64, error) {
	query := fmt.Sprintf("SELECT amount FROM %s WHERE currency=%s", CurrencyTable, key)
	var amount float64
	err := r.db.Get(&amount, query)

	return amount, err
}

func (r *CurrencyRepo) GetAllRates(ctx context.Context) (map[string]float64, error) {
	resultMap := map[string]float64{}
	query := fmt.Sprintf("SELECT currency, amount FROM %s", CurrencyTable)
	rows, err := r.db.Query(query)
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

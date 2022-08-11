package bst

import "github.com/shopspring/decimal"

type Trade struct {
	Price  decimal.Decimal
	Amount decimal.Decimal
	Total  decimal.Decimal
}

package bst

import "github.com/shopspring/decimal"

type Order struct {
	Id     int
	Price  decimal.Decimal
	Amount decimal.Decimal
	Side   OrderSide
}

func CreateNewOrder(Id int, price decimal.Decimal, amount decimal.Decimal, side OrderSide) *Order {
	return &Order{
		Id:     Id,
		Price:  price,
		Amount: amount,
		Side:   side,
	}
}

// filled
func (o *Order) Filled() bool {
	return o.Amount.IsZero()
}

package bst

type Order struct {
	Price  int
	Amount int
	Side   OrderSide
}

func CreateNewOrder(price int, amount int, side OrderSide) *Order {
	return &Order{
		Price:  price,
		Amount: amount,
		Side:   side,
	}
}

package bst

import "log"

type OrderSide string

const (
	OrderSideBuy  OrderSide = "asks"
	OrderSideSell OrderSide = "bids"
)

type PriceLevel struct {
	Price       int
	TotalAmount int
	Side        OrderSide

	Orders []*Order
}

func CreatePriceLevel(order *Order) *PriceLevel {
	return &PriceLevel{
		Price:       order.Price,
		TotalAmount: order.Amount,
		Side:        order.Side,

		Orders: []*Order{order},
	}
}

func (p *PriceLevel) Add(order *Order) {
	p.TotalAmount += order.Amount

	p.Orders = append(p.Orders, order)

	log.Println("orders", p.Orders)
}

func (n *PriceLevel) Remove(order *Order) {
	for i := 0; i < len(n.Orders); i++ {
		if n.Orders[i].Amount > order.Amount {
			n.Orders[i].Amount -= order.Amount
			n.TotalAmount -= order.Amount
			log.Println("mount", n.TotalAmount)

			break
		} else if n.Orders[i].Amount <= order.Amount {
			n.Orders[i].Amount = 0
			n.Orders = append(n.Orders[:i], n.Orders[i+1:]...)
			n.TotalAmount -= order.Amount
			log.Println("mount", n.TotalAmount)
		}
	}
}

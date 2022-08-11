package bst

import (
	"log"
	"sync"

	"github.com/shopspring/decimal"
)

type OrderSide string

const (
	OrderSideBuy  OrderSide = "bids"
	OrderSideSell OrderSide = "asks"
)

type PriceLevel struct {
	Price       decimal.Decimal
	TotalAmount decimal.Decimal
	Side        OrderSide

	Orders []*Order

	sync.Mutex
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
	p.TotalAmount.Add(order.Amount)

	p.Orders = append(p.Orders, order)
}

func (p *PriceLevel) Get(id int) *Order {
	p.Lock()
	defer p.Unlock()

	for _, order := range p.Orders {
		if order.Id == id {
			return order
		}
	}

	return nil
}

func (n *PriceLevel) Remove(id int) {
	n.Lock()
	defer n.Unlock()

	for i := range n.Orders {
		if n.Orders[i].Id == id {

			log.Println("id:", id, n.Orders[i].Price)
			n.Orders = append(n.Orders[:i], n.Orders[i+1:]...)

			log.Println("orders", n.Orders)

			break
		}
	}
}

func (p *PriceLevel) Filled() bool {
	return p.TotalAmount.IsZero()
}

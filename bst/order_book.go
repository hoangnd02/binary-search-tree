package bst

import (
	"sync"

	"github.com/emirpasic/gods/trees/redblacktree"
	rbt "github.com/emirpasic/gods/trees/redblacktree"
	"github.com/shopspring/decimal"
)

type OrderBook struct {
	Bids *rbt.Tree
	Asks *rbt.Tree

	sync.Mutex
}

func CreateOrderBook() *OrderBook {
	return &OrderBook{
		Bids: rbt.NewWith(
			func(a, b interface{}) int {
				aAsserted := a.(decimal.Decimal)
				bAsserted := b.(decimal.Decimal)
				switch {
				case aAsserted.GreaterThan(bAsserted):
					return 1
				case aAsserted.LessThan(bAsserted):
					return -1
				default:
					return 0
				}
			}),
		Asks: rbt.NewWith(func(a, b interface{}) int {
			aAsserted := a.(decimal.Decimal)
			bAsserted := b.(decimal.Decimal)
			switch {
			case aAsserted.LessThan(bAsserted):
				return 1
			case aAsserted.GreaterThan(bAsserted):
				return -1
			default:
				return 0
			}
		}),
	}
}

func (o *OrderBook) Get(value decimal.Decimal, side OrderSide) *PriceLevel {
	o.Lock()
	defer o.Unlock()

	if side == OrderSideBuy {
		node := o.Bids.GetNode(value)
		return node.Value.(*PriceLevel)
	} else {
		node := o.Asks.GetNode(value)
		return node.Value.(*PriceLevel)
	}

}

func (o *OrderBook) Add(order *Order) *Trade {
	o.Lock()
	defer o.Unlock()

	var offers *rbt.Tree
	if order.Side == OrderSideBuy {
		offers = o.Bids
	} else {
		offers = o.Asks
	}

	newTrade := o.Match(order)

	if order.Amount.GreaterThan(decimal.Zero) {
		node := offers.GetNode(order.Price)
		if node != nil {
			priceLevel := node.Value.(*PriceLevel)
			priceLevel.Add(order)
		} else {
			offers.Put(order.Price, CreatePriceLevel(order))
		}
	}

	return newTrade
}

func (o *OrderBook) Match(order *Order) *Trade {
	var offers *redblacktree.Tree
	if order.Side == OrderSideBuy {
		offers = o.Asks
	} else {
		offers = o.Bids
	}

	iter := offers.Iterator()

	for iter.Next() {
		amount := decimal.NewFromFloat(0)

		if order.Amount.IsZero() {
			break
		}

		priceLevel := iter.Value().(*PriceLevel)

		if order.Side == OrderSideBuy {
			if order.Price.LessThan(priceLevel.Price) {
				continue
			}
		} else {
			if order.Price.GreaterThan(priceLevel.Price) {
				continue
			}
		}

		for _, order_offer := range priceLevel.Orders {
			if order.Filled() {
				break
			}

			min_amount := decimal.Min(order.Amount, order_offer.Amount)

			amount = amount.Add(min_amount)

			order.Amount = order.Amount.Sub(min_amount)
			order_offer.Amount = order_offer.Amount.Sub(min_amount)
			priceLevel.TotalAmount = priceLevel.TotalAmount.Sub(min_amount)

			if order_offer.Filled() {
				priceLevel.Remove(order_offer.Id)
			}
		}

		if priceLevel.Filled() {
			offers.Remove(priceLevel.Price)
		}

		return &Trade{
			Price:  order.Price,
			Amount: amount,
			Total:  order.Price.Mul(amount),
		}
	}
	return &Trade{}
}

// [[1,2], [3,4], [5,6]]
func (o *OrderBook) Depth() ([][]decimal.Decimal, [][]decimal.Decimal) {
	o.Lock()
	defer o.Unlock()

	// Asks
	ask_values := make([][]decimal.Decimal, 0)

	it_ask := o.Asks.Iterator()

	for it_ask.Next() {

		priceLevel := it_ask.Value().(*PriceLevel)

		ask_price_level := []decimal.Decimal{
			priceLevel.Price,
			priceLevel.TotalAmount,
		}

		ask_values = append(ask_values, ask_price_level)
	}

	// Bids
	bid_values := make([][]decimal.Decimal, 0)

	it_bid := o.Bids.Iterator()

	for it_bid.Next() {

		priceLevel := it_bid.Value().(*PriceLevel)

		bid_values = append(bid_values, []decimal.Decimal{
			priceLevel.Price,
			priceLevel.TotalAmount,
		})
	}

	return ask_values, bid_values
}

func (o *OrderBook) Remove(id int, side OrderSide, price decimal.Decimal) {
	o.Lock()
	defer o.Unlock()

	var tree *rbt.Tree
	if side == OrderSideBuy {
		tree = o.Bids
	} else {
		tree = o.Asks
	}

	node, bool := tree.Get(price)
	if bool {
		priceLevel := node.(*PriceLevel)
		priceLevel.Remove(id)

		if len(priceLevel.Orders) == 0 {
			tree.Remove(price)
		}
	}
}

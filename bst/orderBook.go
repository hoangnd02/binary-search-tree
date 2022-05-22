package bst

import (
	"log"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

type OrderBook struct {
	Asks *rbt.Tree
	Bids *rbt.Tree
}

func CreateOrderBook() *OrderBook {
	return &OrderBook{
		Asks: rbt.NewWithIntComparator(),
		Bids: rbt.NewWithIntComparator(),
	}
}

func (o *OrderBook) Add(order *Order) {
	price := order.Price

	if order.Side == OrderSideBuy {
		node, _ := o.Asks.Floor(price)
		if node != nil {
			priceLevel := node.Value.(*PriceLevel)
			priceLevel.Add(order)
			// o.Asks.Put(price, priceLevel)
		} else {
			newPriceLevel := CreatePriceLevel(order)
			o.Asks.Put(price, newPriceLevel)
		}
	} else {
		node, _ := o.Bids.Floor(price)
		if node != nil {
			priceLevel := node.Value.(*PriceLevel)
			priceLevel.Add(order)
			o.Bids.Put(price, priceLevel)
		} else {
			newPriceLevel := CreatePriceLevel(order)
			o.Bids.Put(price, newPriceLevel)
		}
	}
}

func (o *OrderBook) Remove(order *Order) {
	price := order.Price

	if order.Side == OrderSideBuy {
		node, _ := o.Asks.Floor(price)
		if node != nil {
			priceLevel := node.Value.(*PriceLevel)
			priceLevel.Remove(order)
			if priceLevel.TotalAmount == 0 {
				o.Asks.Remove(price)
			}
		}
	} else {
		node, _ := o.Bids.Floor(price)
		if node != nil {
			priceLevel := node.Value.(*PriceLevel)
			priceLevel.Remove(order)
			if priceLevel.TotalAmount == 0 {
				o.Bids.Remove(price)
			}
		}
	}
}

func (o *OrderBook) Match() {
	for _, ask := range o.Asks.Values() {
		log.Println("ask", ask)
		priceLevelAsks := ask.(*PriceLevel)

		for _, orderAsk := range priceLevelAsks.Orders {

			bid, bool := o.Bids.Floor(priceLevelAsks.Price)
			if !bool {
				break
			}

			priceLevelBids := bid.Value.(*PriceLevel)
			for _, orderBid := range priceLevelBids.Orders {
				priceLevelAsks.Remove(orderBid)
				priceLevelBids.Remove(orderAsk)

				if priceLevelBids.TotalAmount == 0 {
					o.Bids.Remove(priceLevelBids.Price)
					break
				}
			}
		}
	}
}

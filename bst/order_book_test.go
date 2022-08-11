package bst

import (
	"log"
	"testing"

	"github.com/shopspring/decimal"
)

func TestGet(t *testing.T) {
	orderBook := CreateOrderBook()

	orders := []*Order{
		{
			Id:     1,
			Price:  decimal.NewFromFloat(8),
			Amount: decimal.NewFromFloat(2),
			Side:   OrderSideBuy,
		},
		{
			Id:     2,
			Price:  decimal.NewFromFloat(8),
			Amount: decimal.NewFromFloat(1),
			Side:   OrderSideBuy,
		},
	}

	for _, order := range orders {
		orderBook.Add(order)
	}

	priceLevel := orderBook.Get(decimal.NewFromFloat(8), OrderSideBuy)
	if len(priceLevel.Orders) != 2 {
		t.Errorf("Expected 2 orders, got %d", len(priceLevel.Orders))
	}
}

func TestAdd(t *testing.T) {
	orderBook := CreateOrderBook()

	orders := []*Order{
		{
			Id:     1,
			Price:  decimal.NewFromFloat(8),
			Amount: decimal.NewFromFloat(2),
			Side:   OrderSideBuy,
		},
		{
			Id:     2,
			Price:  decimal.NewFromFloat(6),
			Amount: decimal.NewFromFloat(1),
			Side:   OrderSideBuy,
		},
	}

	for _, order := range orders {
		orderBook.Add(order)
	}

	asks, bids := orderBook.Depth()

	if len(bids) != 2 {
		t.Errorf("Expected 2 bids, got %d", len(bids))
	}

	if len(asks) != 0 {
		t.Errorf("Expected 0 asks, got %d", len(asks))
	}
}

func TestMatch(t *testing.T) {
	orderBook := CreateOrderBook()

	listOrders := []*Order{
		{
			Id:     1,
			Price:  decimal.NewFromFloat(6),
			Amount: decimal.NewFromFloat(1),
			Side:   OrderSideBuy,
		},
		{
			Id:     2,
			Price:  decimal.NewFromFloat(6),
			Amount: decimal.NewFromFloat(1),
			Side:   OrderSideSell,
		},
	}

	for key, order := range listOrders {
		trade := orderBook.Add(order)

		if key == 1 {
			if !trade.Amount.Equal(decimal.NewFromFloat(1)) {
				t.Errorf("Expected 1 trade amount, got %s", trade.Amount)
			}

			if !trade.Price.Equal(decimal.NewFromFloat(6)) {
				t.Errorf("Expected 6 trade price, got %s", trade.Price)
			}

			if !trade.Total.Equal(decimal.NewFromFloat(6)) {
				t.Errorf("Expected 6 trade total, got %s", trade.Total)
			}
		}
	}

	bids, asks := orderBook.Depth()
	if len(bids) != 0 {
		t.Errorf("Expected 0 bids, got %d", len(bids))
	}

	if len(asks) != 0 {
		t.Errorf("Expected 0 asks, got %d", len(asks))
	}
}

func TestRemove(t *testing.T) {
	orderBook := CreateOrderBook()

	orders := []*Order{
		{
			Id:     1,
			Price:  decimal.NewFromFloat(8),
			Amount: decimal.NewFromFloat(2),
			Side:   OrderSideBuy,
		},
		{
			Id:     2,
			Price:  decimal.NewFromFloat(6),
			Amount: decimal.NewFromFloat(1),
			Side:   OrderSideBuy,
		},
	}

	for _, order := range orders {
		orderBook.Add(order)
	}

	orderBook.Remove(2, OrderSideBuy, decimal.NewFromFloat(6))

	asks, bids := orderBook.Depth()

	log.Println("bid", bids)
	log.Println("ask", asks)

	if len(bids) != 1 {
		t.Errorf("Expected 1 bids, got %d", len(bids))
	}

	if len(asks) != 0 {
		t.Errorf("Expected 0 asks, got %d", len(asks))
	}
}

func TestDepth(t *testing.T) {
	orderBook := CreateOrderBook()

	orders := []*Order{
		{
			Id:     1,
			Price:  decimal.NewFromFloat(8),
			Amount: decimal.NewFromFloat(2),
			Side:   OrderSideBuy,
		},
		{
			Id:     2,
			Price:  decimal.NewFromFloat(6),
			Amount: decimal.NewFromFloat(1),
			Side:   OrderSideBuy,
		},
	}

	for _, order := range orders {
		orderBook.Add(order)
	}

	asks, bids := orderBook.Depth()

	log.Println("asks", asks)
	log.Println("bids", bids)

	for _, bid := range bids {
		if len(bid) != 2 {
			t.Errorf("Expected 2 bids, got %d", len(bids))
		}
	}

	for _, ask := range asks {
		if len(ask) != 0 {
			t.Errorf("Expected 0 asks, got %d", len(asks))
		}
	}
}

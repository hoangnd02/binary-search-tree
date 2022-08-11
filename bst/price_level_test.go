package bst

import (
	"log"
	"testing"

	"github.com/shopspring/decimal"
)

func TestAddOrder(t *testing.T) {

	orders := []*Order{
		{
			Id:     1,
			Price:  decimal.NewFromFloat(8),
			Amount: decimal.NewFromFloat(2),
			Side:   OrderSideBuy,
		},
	}

	priceLevel := &PriceLevel{
		Price:       decimal.NewFromFloat(8),
		TotalAmount: decimal.NewFromFloat(2),
		Orders:      orders,
		Side:        OrderSideBuy,
	}

	priceLevel.Add(&Order{
		Id:     2,
		Price:  decimal.NewFromFloat(6),
		Amount: decimal.NewFromFloat(1),
		Side:   OrderSideBuy,
	})

	if len(priceLevel.Orders) != 2 {
		t.Errorf("Expected 2 order, got %d", len(priceLevel.Orders))
	}
}

func TestGetOrder(t *testing.T) {

	orders := []*Order{
		{
			Id:     1,
			Price:  decimal.NewFromFloat(8),
			Amount: decimal.NewFromFloat(2),
			Side:   OrderSideBuy,
		},
	}

	priceLevel := &PriceLevel{
		Price:       decimal.NewFromFloat(8),
		TotalAmount: decimal.NewFromFloat(2),
		Orders:      orders,
		Side:        OrderSideBuy,
	}

	order := priceLevel.Get(1)
	log.Println("order:", order.Amount)

	if !order.Amount.Equal(decimal.NewFromFloat(2)) {
		t.Errorf("Expected 2 price level, got %d", order.Amount)
	}
}

func TestRemoveOrder(t *testing.T) {

	orders := []*Order{
		{
			Id:     1,
			Price:  decimal.NewFromFloat(8),
			Amount: decimal.NewFromFloat(2),
			Side:   OrderSideBuy,
		},
	}

	priceLevel := &PriceLevel{
		Price:       decimal.NewFromFloat(8),
		TotalAmount: decimal.NewFromFloat(2),
		Orders:      orders,
		Side:        OrderSideBuy,
	}

	priceLevel.Remove(1)

	if len(priceLevel.Orders) != 0 {
		t.Errorf("Expected 0 order, got %d", len(priceLevel.Orders))
	}
}

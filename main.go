package main

import (
	"log"

	"github.com/hoanggggg5/bst/bst"
	"github.com/shopspring/decimal"
)

func main() {
	orderBook := bst.CreateOrderBook()

	orderBook.Add(bst.CreateNewOrder(1, decimal.NewFromFloat(8), decimal.NewFromFloat(1), bst.OrderSideBuy))
	// orderBook.Add(bst.CreateNewOrder(2, decimal.NewFromFloat(2), decimal.NewFromFloat(1), bst.OrderSideBuy))

	orderBook.Add(bst.CreateNewOrder(4, decimal.NewFromFloat(4), decimal.NewFromFloat(1), bst.OrderSideBuy))
	orderBook.Add(bst.CreateNewOrder(4, decimal.NewFromFloat(5), decimal.NewFromFloat(1), bst.OrderSideBuy))

	// 	log.Println("Bids:", orderBook.Bids)
	// 	log.Println("Asks:", orderBook.Asks)

	// 	log.Println("Bids")
	// 	for _, bid := range orderBook.Bids.Values() {
	// 		priceLevel := bid.(*bst.PriceLevel)
	// 		log.Println("priceLevel:", priceLevel)
	// 	}
	// 	log.Println("Asks")
	// 	for _, bid := range orderBook.Asks.Values() {
	// 		priceLevel := bid.(*bst.PriceLevel)
	// 		log.Println("priceLevel:", priceLevel)
	// 	}

	// 	node := orderBook.Bids.GetNode(decimal.NewFromFloat(2))
	// 	log.Println("node:", node)

	orderBook.Remove(1, bst.OrderSideBuy, decimal.NewFromFloat(8))

	asks, bids := orderBook.Depth()

	log.Println("asks:", asks)
	log.Println("bids:", bids)
}

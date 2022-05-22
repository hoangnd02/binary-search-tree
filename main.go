package main

import (
	"log"

	"github.com/hoanggggg5/bst/bst"
)

func main() {
	orderBook := bst.CreateOrderBook()

	orderBook.Add(bst.CreateNewOrder(10, 1, bst.OrderSideBuy))
	orderBook.Add(bst.CreateNewOrder(5, 1, bst.OrderSideBuy))
	orderBook.Add(bst.CreateNewOrder(13, 1, bst.OrderSideBuy))
	orderBook.Add(bst.CreateNewOrder(8, 1, bst.OrderSideBuy))

	orderBook.Add(bst.CreateNewOrder(10, 1, bst.OrderSideSell))
	orderBook.Add(bst.CreateNewOrder(8, 1, bst.OrderSideSell))
	orderBook.Add(bst.CreateNewOrder(13, 2, bst.OrderSideSell))

	// orderBook.Match()

	log.Println("Asks:", orderBook.Asks)
	log.Println("Bids:", orderBook.Bids)

	for _, ask := range orderBook.Asks.Values() {
		priceLevel := ask.(*bst.PriceLevel)
		for _, order := range priceLevel.Orders {
			log.Println("order asks", order)
		}
	}
}

package main

import (
	"che/dt/week2/orders/business"
	"che/dt/week2/orders/types"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <json file with orders>")
		os.Exit(1)
	}

	var orders []types.Order
	business.ReadOrders(os.Args[1], &orders)

	fmt.Printf("Aantal orders: %v\n", len(orders))
	fmt.Printf("Total order price: %.2f\n", business.CalculateTotal(orders))
}

func GetMostExpensiveOrder(orders []types.Order) types.Order {
	max = 0
	maxOrder = nil
	foreach order in orders
	  orderPrice = CalculateTotalOrderPrice(order)
	  if orderPrice > max {
		max = orderPrice
		maxOrder = order
	  }
	
	return maxOrder
}
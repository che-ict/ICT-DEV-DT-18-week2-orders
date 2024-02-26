package business

import (
	"che/dt/week2/orders/types"
)

func GetMostExpensiveOrder(orders []types.Order) types.Order {

	var max float32
	var maxOrder types.Order
	var orderPrice float32
	for _, order := range orders {
		orderPrice = CalculateTotalOrderPrice(order)
		if orderPrice > max {
			max = orderPrice
			maxOrder = order
		}
	}

	return maxOrder
}

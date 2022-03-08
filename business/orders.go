package business

import (
	"che/dt/week2/orders/types"
)

func GetMostExpensiveOrder(orders []types.Order) types.Order {

	max := 0
	maxOrder := nil
	orderPrice := 0
	for order := range orders {
		orderPrice = CalculateTotalOrderPrice(order)
		if orderPrice > max {
			max = orderPrice
			maxOrder = order
		}
	}

	return maxOrder
}

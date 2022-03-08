package business

import (
	"che/dt/week2/orders/types"
	utils "che/dt/week2/orders/util"
)

func ReadOrders(filename string, orders *[]types.Order) {
	err := utils.LoadData(filename, orders)

	if err != nil {
		panic(err)
	}
}

func GetMostExpensiveOrder(orders []types.Order) types.Order {
	max = 0
	maxOrder = nil
	for _, order := range orders.json {
		orderPrice = CalculateTotalOrderPrice(order)

		if orderPrice > max {
			max = orderPrice
			maxOrder = order

			return maxOrder
		}

	}

}

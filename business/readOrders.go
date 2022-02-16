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

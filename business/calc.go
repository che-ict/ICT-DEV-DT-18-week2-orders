package business

import (
	"che/dt/week2/orders/types"
	"fmt"
)

// Retourneert de totale bruto pijs (incl. BTW) van alle gegeven orders
func CalculateTotal(orders []types.Order) float32 {
	var sum float32 = 0

	for _, order := range orders {

		sum += CalculateTotalOrderPrice(order)
	}

	return sum
}

// Retourneert de totale bruto pijs (incl. BTW) van de gegeven order
func CalculateTotalOrderPrice(order types.Order) float32 {
	var btwFactor float32

	var sum float32 = 0
	for _, regel := range order.Regels {

		switch regel.Btw {
		case "21%":
			btwFactor = 1.21
		case "9%":
			btwFactor = 1.9
		case "0%":
			btwFactor = 1.0
		default:
			panic(fmt.Sprintf("Unknown btw percentage (%v) in order for %v on %v, regel %v!", regel.Btw, order.Organisatie, order.Datum, regel.Nummer))
		}

		sum += float32(regel.Aantal) * regel.Prijs * btwFactor * order.Korting
	}

	return sum
}

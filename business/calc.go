package business

import (
	"che/dt/week2/orders/types"
	"errors"
	"fmt"
)

// Retourneert de totale bruto pijs (incl. BTW) van alle gegeven orders
func CalculateTotal(orders []types.Order) float32 {
	var totaal float32 = 0

	for _, order := range orders {

		totaal += CalculateTotalOrderPrice(order)
	}

	return totaal
}

// Retourneert de totale bruto pijs (incl. BTW) van de gegeven order
func CalculateTotalOrderPrice(order types.Order) float32 {
	err := ValidateOrder(order)
	if err != nil {
		panic(fmt.Sprint("Berekening mislukt: ", err))
	}

	var btwFactor float32

	var totaal float32 = 0
	for _, regel := range order.Regels {

		switch regel.Btw {
		case "21%":
			btwFactor = 1.21
		case "9%":
			btwFactor = 1.09
		case "7%":
			btwFactor = 1.07
		case "0%":
			btwFactor = 1.0
		default:
			panic(fmt.Sprintf("Unknown btw percentage (%v) in order for %v on %v, regel %v!", regel.Btw, order.Organisatie, order.Datum, regel.Nummer))
		}

		totaal += float32(regel.Aantal) * regel.Prijs * btwFactor
	}

	return totaal

}

func ValidateOrder(order types.Order) error {
	if order.Id < 0 {
		return errors.New("id is negatief")
	} else if order.Korting > float32(1) || order.Korting < float32(0) {
		return errors.New("korting is geen juist getal")
	} else {
		for i, regel := range order.Regels {
			if regel.Aantal < 0 {
				return errors.New(fmt.Sprint("Ongeldig veld Aantal in Order nummer ", i))
			} else if regel.Prijs < float32(0) {
				return errors.New(fmt.Sprint("Ongeldig veld Prijs in Order nummer ", i))
			}
		}
	}

	return nil
}

func GetMostExpensiveOrder(orders []types.Order) types.Order {
	var max = 0
	MaxOrder := types.Order{}
	for _, order := range orders {
		var orderPrice = CalculateTotalOrderPrice(order)
		if orderPrice > float32(max) {
			max = int(orderPrice)
			MaxOrder = order
		}
	}
	return MaxOrder

func PrintOrders(orders []types.Order) {
	for _, order := range orders {
		fmt.Printf("- %v: %v regels, %.2f euro\n", order.Organisatie, len(order.Regels), CalculateTotalOrderPrice(order))
	}
}

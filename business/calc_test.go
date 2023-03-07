package business

import (
	"testing"
)

func TestCalculateTotalOrderPrice_MultipleOrderLines_GivesExpectedTotal(t *testing.T) {

	// honkyplankie

	//hallo
	//EINDELUK
	//a
	// order := types.Order{
	// 	Id:          0,
	// 	Organisatie: "Centuria",
	// 	Regels: []types.OrderRegel{
	// 		{
	// 			Nummer:       0,
	// 			Btw:          "21%",
	// 			Aantal:       4,
	// 			Prijs:        94.85,
	// 			Omschrijving: "aute aute adipisicing",
	// 		},
	// 		{
	// 			Nummer:       1,
	// 			Btw:          "9%",
	// 			Aantal:       2,
	// 			Prijs:        163.07,
	// 			Omschrijving: "aute aute adipisicing",
	// 		},
	// 	},
	// 	Korting:    0,
	// 	Definitief: true,
	// 	Datum:      time.Now(),
	// }

	// When

	// Then

}

func TestCalculateTotalOrderPrice_NoOrderLines_GivesTotalZero(t *testing.T) {

	// Given

	// When

	// Then

}

func TestCalculateTotalOrderPrice_SingleOrderLine_GivesExpectedTotal(t *testing.T) {

	// Given

	// When

	// Then

}

func TestCalculateTotalOrderPrice_OrderLineWithDiscount_GivesExpectedTotal(t *testing.T) {

	// Given

	// When

	// Then

}

func TestCalculateTotalOrderPrice_OrderLineWith21PercentVat_GivesExpectedTotal(t *testing.T) {

	// Given

	// When

	// Then

}

func TestCalculateTotalOrderPrice_OrderLineWith9PercentVat_GivesExpectedTotal(t *testing.T) {

	// Given

	// When

	// Then

}

func TestCalculateTotalOrderPrice_NegativeQuantity_GivesError(t *testing.T) {

	// Given

	// When

	// Then

}

func TestCalculateTotalOrderPrice_NegativePrice_GivesError(t *testing.T) {

	// Given

	// When

	// Then

}

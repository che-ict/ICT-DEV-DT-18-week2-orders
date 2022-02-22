package business

import (
	"che/dt/week2/orders/types"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTotalOrderPrice_MultipleOrderLines_GivesExpectedTotal(t *testing.T) {

	// Given

	order := types.Order{
		Id:          0,
		Organisatie: "Centuria",
		Regels: []types.OrderRegel{
			{
				Nummer:       0,
				Btw:          "21%",
				Aantal:       4,
				Prijs:        93.85,
				Omschrijving: "aute aute adipisicing",
			},
			{
				Nummer:       1,
				Btw:          "9%",
				Aantal:       2,
				Prijs:        163.07,
				Omschrijving: "aute aute adipisicing",
			},
		},
		Korting:    0,
		Definitief: true,
		Datum:      time.Now(),
	}

	// When, haha fck u

	total := CalculateTotalOrderPrice(order)

	// Then

	assert.Equal(t, float32(809.7266), total)

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

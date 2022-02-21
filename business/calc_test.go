package business

import (
	"che/dt/week2/orders/types"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTests(t *testing.T) {

	assert.Equal(t, true, true)
}

func TestCalculateTotalOrderPrice_SingleOrderLine_GivesExpectedTotal(t *testing.T) {

	// Given

	order := types.Order{
		Id:          1,
		Organisatie: "Made by Bob",
		Regels: []types.OrderRegel{
			{
				Nummer:       3,
				Btw:          "21%",
				Aantal:       1,
				Prijs:        20,
				Omschrijving: "Mooi!!",
			},
		},
		Korting:    0,
		Definitief: true,
		Datum:      time.Now(),
	}

	// When

	total := CalculateTotalOrderPrice(order)

	// Then

	assert.Equal(t, float32(24.20), total)

}

func TestCalculateTotalOrderPrice_NoOrderLine_GivesZeroTotal(t *testing.T) {

	// Given

	// When

	// Then

}

func Sum(a int, b int) int {
	return a + b
}

func TestSum_TenPlusOne_EqualsEleven(t *testing.T) {

	// Given

	a := 10
	b := 1

	// When

	total := Sum(a, b)

	// Then

	assert.Equal(t, 11, total)
}

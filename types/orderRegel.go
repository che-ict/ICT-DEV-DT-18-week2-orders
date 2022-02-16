package types

type OrderRegel struct {
	Nummer       int     `json:"nummer"`
	Btw          string  `json:"btw"`
	Aantal       int     `json:"aantal"`
	Prijs        float32 `json:"prijs"`
	Omschrijving string  `json:"omschrijving"`
}

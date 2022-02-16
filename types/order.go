package types

import "time"

type Order struct {
	Id          int          `json:"id"`
	Organisatie string       `json:"organisatie"`
	Korting     float32      `json:"korting"`
	Definitief  bool         `json:"definitief"`
	Datum       time.Time    `json:"datum"`
	Regels      []OrderRegel `json:"regels"`
}

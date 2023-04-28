package models

type Location struct {
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
	Country string  `json:"country"`
	City    string  `json:"city"`
	Type    string  `json:"type"` // if it is urban or country side
}

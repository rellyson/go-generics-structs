package models

import "time"

type Accessory struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type Base struct {
	ID              int         `json:"id"`
	Brand           string      `json:"brand"`
	Model           string      `json:"model"`
	FabricationYear int64       `json:"fabrication_year"`
	ModelYear       int64       `json:"model_year"`
	Value           float64     `json:"value"`
	Accessories     []Accessory `json:"accessories"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

package models

type Truck struct {
	Base
	LoadCapacity float64 `json:"load_capacity"`
}

package models

type Motorcycle struct {
	Base
	HasTrunk bool `json:"has_trunk"`
}

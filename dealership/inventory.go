package dealership

import (
	"errors"

	"github.com/rellyson/go-generics-structs/dealership/models"
)

type Vehicle interface {
	models.Car | models.Motorcycle | models.Truck
}

type Inventory[T Vehicle] struct {
	items map[int]T
}

func (inventory *Inventory[T]) AddItem(item T) {
	if inventory.items == nil {
		inventory.items = make(map[int]T)
	}

	index := len(inventory.items)
	inventory.items[index] = item
}

func (inventory Inventory[T]) GetItem(index int) (T, error) {
	if v, ok := inventory.items[index]; ok {
		return v, nil
	}

	return *new(T), errors.New("index out of range")
}

func (inventory Inventory[T]) GetItems() map[int]T {
	if inventory.items == nil {
		return nil
	}

	return inventory.items
}

package main

import (
	"fmt"
	"time"

	"github.com/rellyson/go-generics-structs/dealership"
	"github.com/rellyson/go-generics-structs/dealership/models"
)

func main() {
	carsInventory := dealership.Inventory[models.Car]{}
	motorcycleInventory := dealership.Inventory[models.Motorcycle]{}

	carsInventory.AddItem(models.Car{
		Base: models.Base{
			ID:              1,
			Brand:           "Chevrolet",
			Model:           "Onix LTZ 1.0 Turbo",
			FabricationYear: 2021,
			ModelYear:       2022,
			Value:           88.643,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			Accessories: []models.Accessory{
				{
					Name:  "Jogo de tapetes esportivo",
					Value: 108.50,
				},
			},
		},
		IsSportive: true,
	})

	carsInventory.AddItem(models.Car{
		Base: models.Base{
			ID:              2,
			Brand:           "Fiat",
			Model:           "Argo S Design 1.3 2022",
			FabricationYear: 2021,
			ModelYear:       2022,
			Value:           82.264,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			Accessories: []models.Accessory{
				{
					Name:  "Retrovisor fotocrômico",
					Value: 203.57,
				},
			},
		},
		IsSportive: true,
	})

	motorcycleInventory.AddItem(models.Motorcycle{
		Base: models.Base{
			ID:              1,
			Brand:           "Triumph",
			Model:           "Tiger 900 Rally",
			FabricationYear: 2020,
			ModelYear:       2021,
			Value:           65.943,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			Accessories: []models.Accessory{
				{
					Name:  "Baú Esportivo 200l",
					Value: 809.27,
				},
			},
		},
		HasTrunk: true,
	})

	for i, v := range carsInventory.GetItems() {
		fmt.Printf("Car number %v: %v \n", i+1, v)
	}

	for i, v := range motorcycleInventory.GetItems() {
		fmt.Printf("Motorcycle number %v: %v \n", i+1, v)
	}
}

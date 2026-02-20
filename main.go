package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}
type RegularTruck struct {
	id    string
	cargo int
}

func (t *RegularTruck) LoadCargo() error {
	return nil
}
func (t *RegularTruck) UnloadCarg() error {
	return nil
}

type EletricTruck struct {
	id      string
	cargo   int
	battery float64
}

func processTruck(truck RegularTruck) error {

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %s", err)
	}
	return ErrNotImplemented
}

func main() {
	trucks := []RegularTruck{
		{id: "Truck-1"},
		{id: "Truck-2"},
		{id: "Truck-3"},
	}
	/*
		eTrucks := []EletricTruck{
			{id: "Eletric-truck-1"},
		}*/

	for i, truck := range trucks {
		fmt.Printf("Truck %s has arrived.\n", truck.id)

		/*
			if err := processTruck(trucks[i]); err != nil {
				if errors.Is(err, ErrNotImplemented) {
					continue
				}

				log.Fatalf("Error processing truck %d: %s\n", i, err)
			}

		*/
		err := processTruck(truck)
		if err != nil {
			log.Fatalf("Error processing truck %d: %s\n", i, err)
		}

		fmt.Printf("Truck %s has departured.\n", truck.id)

	}
}

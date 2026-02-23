package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}
type NormalTruck struct {
	id    string
	cargo int
}

func (t *NormalTruck) LoadCargo() error {
	t.cargo += 1
	return nil
}
func (t *NormalTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

func (e *ElectricTruck) LoadCargo() error {
	e.cargo += 1
	e.battery -= 1
	return nil
}
func (e *ElectricTruck) UnloadCargo() error {
	e.cargo = 0
	e.battery -= 1
	return nil
}

func processTruck(truck Truck) error {
	fmt.Printf("processing truck %v\n", truck)

	time.Sleep(time.Second)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %s", err)
	}

	err := truck.UnloadCargo()
	if err != nil {
		return fmt.Errorf("Error unloading cargo: %s", err)
	}

	return nil
}

func processFleet(trucks []Truck) error {
	var wg sync.WaitGroup
	// wg.Add(len(trucks))
	for _, t := range trucks {
		wg.Add(1)

		go func(t Truck) {
			if err := processTruck(t); err != nil {
				log.Println(err)
			}

			wg.Done()
		}(t)
	}
	wg.Wait()
	return nil
}

func main() {

	fleet := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, battery: 100},
	}

	if err := processFleet(fleet); err != nil {
		fmt.Printf("Error processing fleet %v\n", err)
		return
	}

	fmt.Printf("Fleet processed successfully\n")
}

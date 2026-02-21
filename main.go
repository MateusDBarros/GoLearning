package main

import (
	"errors"
	"fmt"
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
	t.cargo += 1
	return nil
}
func (t *RegularTruck) UnloadCargo() error {
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

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %s", err)
	}

	err := truck.UnloadCargo()
	if err != nil {
		return fmt.Errorf("Error unloading cargo: %s", err)
	}

	return nil
}

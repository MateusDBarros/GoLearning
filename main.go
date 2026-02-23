package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

type contextKey string

var UserIDKey contextKey = "UserID"

var (
	ErrNotImplemented = errors.New("not implemented")
	// ErrTruckNotFound  = errors.New("truck not found")
)

type Trucks interface {
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

func processTruck(ctx context.Context, truck Trucks) error {
	fmt.Printf("processing truck %v\n", truck)

	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	delay := time.Second * 1
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(delay):
		break
	}

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %s", err)
	}

	err := truck.UnloadCargo()
	if err != nil {
		return fmt.Errorf("Error unloading cargo: %s", err)
	}

	return nil
}

func processFleet(ctx context.Context, trucks []Trucks) error {
	var wg sync.WaitGroup
	// wg.Add(len(trucks))
	errorsChan := make(chan error, len(trucks))

	for _, t := range trucks {
		wg.Add(1)

		go func(t Trucks) {
			if err := processTruck(ctx, t); err != nil {
				errorsChan <- err
			}

			wg.Done()
		}(t)
	}
	wg.Wait()
	close(errorsChan)

	var errs []error
	for err := range errorsChan {
		log.Printf("Error processing truck %v\n", err)
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return fmt.Errorf("fleet processing had %d errors", len(errs))
	}
	return nil

}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, UserIDKey, "412")

	fleet := []Trucks{
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, battery: 100},
	}

	if err := processFleet(ctx, fleet); err != nil {
		fmt.Printf("Error processing fleet %v\n", err)
		return
	}

	fmt.Printf("Fleet processed successfully\n")
}

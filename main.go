package main

import (
	"errors"
	"fmt"
	"log"
)

type Truck struct {
	id string
}

func processTruck(truck Truck) error {
	fmt.Printf("Truck ID: %s\n", truck.id)

	return errors.New("my bad bruh")
}

func main() {
	trucks := []Truck{
		{id: "Truck - 1"},
		{id: "Truck - 2"},
		{id: "Truck - 3"},
	}

	for i, truck := range trucks {
		fmt.Printf("Truck %s has arrived:\n", truck.id)

		/*
			err := processTruck(truck)
			if err != nil {
				log.Fatalf("Error processing truck %d: %s\n", i, err)
			}
		*/

		if err := processTruck(trucks[i]); err != nil {
			log.Fatalf("Error processing truck %d: %s\n", i, err)
		}

	}
}

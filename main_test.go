package main

import (
	"testing"
)

func TestTruck(t *testing.T) {

	t.Run("processTruck", func(t *testing.T) {
		t.Run("should load and unload a truck cargo", func(t *testing.T) {

			rt := &NormalTruck{id: "truck1", cargo: 0}
			et := &ElectricTruck{id: "etruck2", cargo: 0}

			err := processTruck(rt)
			if err != nil {
				t.Fatalf("Error processing %s\n", err)
			}

			err = processTruck(et)
			if err != nil {
				t.Fatalf("Error processing %s\n", err)
			}

			// assert
			if rt.cargo != 0 {
				t.Fatalf("Cargo should be 0, was %d\n", rt.cargo)
			}

			if et.battery != -2 {
				t.Fatalf("Battery should be -2, was %f\n", et.battery)
			}

		})
	})
}

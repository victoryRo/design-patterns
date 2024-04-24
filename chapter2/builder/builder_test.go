package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	t.Run("build a car", func(t *testing.T) {
		carBuilder := &CarBuilder{}

		manufacturingComplex.SetBuilder(carBuilder)
		manufacturingComplex.Construct()

		car := carBuilder.GetVehicle()

		if car.Wheels != 4 {
			t.Errorf("Wheels on a car must be 4 and they were %d", car.Wheels)
		}
		if car.Structure != "Car" {
			t.Errorf("Structure on a car must be 'Car' and was %q", car.Structure)
		}
		if car.Seats != 5 {
			t.Errorf("Seats on a car must be 5 and they were %d", car.Seats)
		}
	})

	t.Run("build a motorbike", func(t *testing.T) {
		bikeBuilder := &BikeBuilder{}

		manufacturingComplex.SetBuilder(bikeBuilder)
		manufacturingComplex.Construct()

		bike := bikeBuilder.GetVehicle()
		bike.Seats = 1

		if bike.Wheels != 2 {
			t.Errorf("Wheels on a motorbike must be 2 and they were %d", bike.Wheels)
		}
		if bike.Structure != "MotorBike" {
			t.Errorf("Structure on a motorbike must be 'MotorBike' and was %q", bike.Structure)
		}
	})
}

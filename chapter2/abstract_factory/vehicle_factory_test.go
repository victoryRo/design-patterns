package abstract_factory

import "testing"

// Test Factory of factory
func TestGetVehicleFactory_abs(t *testing.T) {
	t.Run("Type test for GetVehicleFactory 'error type'", func(t *testing.T) {
		_, err := GetVehicleFactory(3)
		if err == nil {
			t.Fatal("Car factory with id 3 should not be recognized")
		}
	})
}

func TestMotorbikeFactory(t *testing.T) {
	// abs factory
	motorbikeF, err := GetVehicleFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Get factory sport motorbike", func(t *testing.T) {
		// motorbike factory
		motorbikeVehicle, err := motorbikeF.GetVehicle(SportMotorbikeType)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Motorbike vehicle has %d wheels and %d seats\n", motorbikeVehicle.GetWheels(), motorbikeVehicle.GetSeats())

		sportBike, ok := motorbikeVehicle.(Motorbike)
		if !ok {
			t.Fatal("Struct assertion has failed")
		}
		t.Logf("Sport motorbike has type %d\n", sportBike.GetType())
	})

	t.Run("Get factory cruise motorbike", func(t *testing.T) {
		motorbikeVehicle, err := motorbikeF.GetVehicle(CruiseMotorbikeType)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Motorbike vehicle has %d wheels\n", motorbikeVehicle.GetWheels())

		cruiseBike, ok := motorbikeVehicle.(Motorbike)
		if !ok {
			t.Fatal("Struct assertion has failed")
		}
		t.Logf("Cruise motorbike has type %d\n", cruiseBike.GetType())
	})

	t.Run("Motorbike vehicle error", func(t *testing.T) {
		_, err := motorbikeF.GetVehicle(3)
		if err == nil {
			t.Fatal("Motorbike of type 3 should not be recognized")
		}
	})
}

func TestCarFactory(t *testing.T) {
	// abs factory
	carF, err := GetVehicleFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Get factory luxury car", func(t *testing.T) {
		// car factory
		carVehicle, err := carF.GetVehicle(LuxuryCarType)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Car vehicle has %d wheels and %d seats\n", carVehicle.GetWheels(), carVehicle.GetSeats())

		luxuryCar, ok := carVehicle.(Car)
		if !ok {
			t.Fatal("Struct assertion has failed")
		}
		t.Logf("Luxury car has %d doors\n", luxuryCar.GetDoors())
	})

	t.Run("Get factory Familiar car", func(t *testing.T) {
		carVehicle, err := carF.GetVehicle(FamiliarCarType)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Car vehicle has %d wheels and %d seats\n", carVehicle.GetWheels(), carVehicle.GetSeats())

		familiarCar, ok := carVehicle.(Car)
		if !ok {
			t.Fatal("Struct assertion has failed")
		}
		t.Logf("Familiar car has %d doors\n", familiarCar.GetDoors())
	})

	t.Run("Car vehicle error", func(t *testing.T) {
		_, err := carF.GetVehicle(3)
		if err == nil {
			t.Fatal("Car type 3 should not be recognized")
		}
	})
}

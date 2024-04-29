package abstract_factory

type FamiliarCar struct{}

// ---- implement interface Car

func (c *FamiliarCar) GetDoors() int {
	return 5
}

// ---- implement interface Vehicle

func (c *FamiliarCar) GetWheels() int {
	return 4
}

func (c *FamiliarCar) GetSeats() int {
	return 5
}

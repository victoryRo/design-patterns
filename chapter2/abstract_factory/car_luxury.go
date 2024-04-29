package abstract_factory

type LuxuryCar struct{}

// ---- implement interface Car

func (c *LuxuryCar) GetDoors() int {
	return 4
}

// ---- implement interface Vehicle

func (c *LuxuryCar) GetWheels() int {
	return 4
}

func (c *LuxuryCar) GetSeats() int {
	return 5
}

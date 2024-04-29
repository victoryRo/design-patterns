package abstract_factory

type CruiseMotorbike struct{}

// ---- implement interface Vehicle

func (c *CruiseMotorbike) GetWheels() int {
	return 2
}

func (c *CruiseMotorbike) GetSeats() int {
	return 2
}

// ---- implement interface Motorbike

func (c *CruiseMotorbike) GetType() int {
	return CruiseMotorbikeType
}

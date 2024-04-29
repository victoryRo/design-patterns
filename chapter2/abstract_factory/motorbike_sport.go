package abstract_factory

type SportMotorbike struct{}

// ---- implement interface Vehicle

func (s *SportMotorbike) GetWheels() int {
	return 2
}

func (s *SportMotorbike) GetSeats() int {
	return 1
}

// ---- implement interface Motorbike

func (s *SportMotorbike) GetType() int {
	return SportMotorbikeType
}

package composition

import "testing"

func TestAthlete_Train(t *testing.T) {
	athlete := Athlete{}
	athlete.Train()
}

func TestSwimmer_Swim(t *testing.T) {
	localSwim := Swim
	swimmer := CompositeSwimmerA{
		MySwim: localSwim,
	}

	swimmer.MyAthlete.Train()
	swimmer.MySwim()
}

func TestAnimal_Swim(t *testing.T) {
	fish := Shark{
		Swim: Swim,
	}

	fish.Swim()
	fish.Eat()
}

func TestSwimmer_Swim2(t *testing.T) {
	swimmer := CompositeSwimmerB{
		new(Athlete),
		new(SwimmerImplementor),
	}

	swimmer.Swim()
	swimmer.Train()
}

func TestSon_GetParentField(t *testing.T) {
	son := Son{}
	GetParentField(&son.P)
}

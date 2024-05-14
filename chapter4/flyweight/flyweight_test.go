package flyweight

import (
	"fmt"
	"testing"
)

func TestTeamFlyweightFactory_GetTeam(t *testing.T) {
	factory := NewTeamFactory()

	teamA1 := factory.GetTeam(TEAM_A)
	if teamA1 == nil {
		t.Error("the pointer to the TEAM_A was nil")
	}

	teamA2 := factory.GetTeam(TEAM_A)
	if teamA2 == nil {
		t.Error("the pointer to the TEAM_A was nil")
	}

	if teamA1 != teamA2 {
		t.Error("TEAM_A pointer weren't the same")
	}

	if factory.GetNumberOfObjects() != 1 {
		t.Errorf("the number of objects created was not 1: %d\n", factory.GetNumberOfObjects())
	}
}

func TestHighVolume(t *testing.T) {
	factory := NewTeamFactory()
	five_hundred_thousand := 500000

	teams := make([]*Team, five_hundred_thousand*2)

	for i := 0; i < five_hundred_thousand; i++ {
		teams[i] = factory.GetTeam(TEAM_A)
	}
	for i := five_hundred_thousand; i < 2*five_hundred_thousand; i++ {
		teams[i] = factory.GetTeam(TEAM_B)
	}

	if factory.GetNumberOfObjects() != 2 {
		t.Errorf("the number of objects created was not 2: %d\n", factory.GetNumberOfObjects())
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("Pointer %d points to %p and is located in %p\n", i, teams[i], &teams[i])
	}
}

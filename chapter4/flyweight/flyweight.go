package flyweight

import "time"

const (
	TEAM_A = iota
	TEAM_B
)

type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}

type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}

type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}

type Team struct {
	ID             uint64
	Name           string
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

// ---------------------------------------------------------------

type teamFlyweightFactory struct {
	createTeams map[int]*Team
}

func NewTeamFactory() teamFlyweightFactory {
	return teamFlyweightFactory{
		createTeams: make(map[int]*Team),
	}
}

func getTeamFactory(team int) Team {
	switch team {
	case TEAM_B:
		return Team{
			ID:   2,
			Name: "TEAM_B",
		}
	default:
		return Team{
			ID:   1,
			Name: "TEAM_A",
		}
	}
}

func (t *teamFlyweightFactory) GetTeam(teamName int) *Team {
	if t.createTeams[teamName] != nil {
		return t.createTeams[teamName]
	}

	team := getTeamFactory(teamName)
	t.createTeams[teamName] = &team

	return t.createTeams[teamName]
}

func (t *teamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createTeams)
}

// ---------------------------------------------------------------

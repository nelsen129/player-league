package store

// PlayerStore records and stores the scores for players
type PlayerStore interface {
	// Should return a player's score
	GetPlayerScore(name string) (int, error)
	// Should increment a player's score
	RecordWin(name string)
	// Should return a league containing all players
	GetLeague() []Player
}

// Player represents an individual player
type Player struct {
	Name string `json:"name"`
	Wins int    `json:"wins"`
}

// PlayerSlice type. Primarily for sorting Players
type PlayerSlice []Player

func (p PlayerSlice) Less(i, j int) bool {
	return p[i].Wins > p[j].Wins
}

func (p PlayerSlice) Len() int {
	return len(p)
}

func (p PlayerSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

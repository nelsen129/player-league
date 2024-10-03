package store

// PlayerStore records and stores the scores for players
type PlayerStore interface {
	// Should return a player's score
	GetPlayerScore(name string) (int, error)
	// Should increment a player's score
	RecordWin(name string)
}

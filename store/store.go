// Package store implements stores for tracking players in a league
//
// The primary interface is the [PlayerStore], which contains methods
// for interacting with the players, including recording wins and
// getting the scores. This package includes multiple structs that
// implement [PlayerStore] to support multiple modes of operation.
// All implemented player stores are safe for concurrent use by
// multiple goroutines
//
// [InMemoryPlayerStore] is a [PlayerStore] that is stored in memory.
// This should only be used for testing or for when persistent storage
// is not required
//
// [FileSystemPlayerStore] is a [PlayerStore] that writes to a file
// on disk.
package store

// PlayerStore records and stores the scores for players
type PlayerStore interface {
	// Should return a player's score
	GetPlayerScore(name string) (int, error)
	// Should increment a player's score
	RecordWin(name string) error
	// Should return a league containing all players
	GetLeague() League
}

// Player represents an individual player
type Player struct {
	Name string `json:"name"`
	Wins int    `json:"wins"`
}

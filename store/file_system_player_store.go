package store

import (
	"encoding/json"
	"io"
	"os"
	"sync"

	"github.com/nelsen129/player-league/store/tape"
)

// FileSystemPlayerStore represents a PlayerStore that is stored in the
// file system as a JSON file
type FileSystemPlayerStore struct {
	mu       sync.Mutex
	database *json.Encoder
	league   League
}

// NewFileSystemPlayerStore returns an FileSystemPlayerStore with an empty store
func NewFileSystemPlayerStore(database *os.File) *FileSystemPlayerStore {
	database.Seek(0, io.SeekStart)
	league, _ := NewLeague(database)
	f := new(FileSystemPlayerStore)
	f.database = json.NewEncoder(tape.NewTape(database))
	f.league = league
	return f
}

// GetPlayerScore returns a score if the player exists or an error if
// they don't
func (f *FileSystemPlayerStore) GetPlayerScore(name string) (int, error) {
	player := f.league.Find(name)

	if player == nil {
		return 0, ErrPlayerNotFound
	}
	return player.Wins, nil
}

// RecordWin increments a player's score by 1. If the player doesn't exist,
// it will create the player and increment their score to 1.
func (f *FileSystemPlayerStore) RecordWin(name string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	player := f.league.Find(name)

	if player == nil {
		f.league = append(f.league, Player{Name: name, Wins: 1})
	} else {
		player.Wins++
	}

	f.database.Encode(f.league)
}

// GetLeague returns an ordered slice containing every Player in the league
// sorted by score, descending
func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}

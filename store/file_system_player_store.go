package store

import (
	"encoding/json"
	"io"
	"sync"
)

// FileSystemPlayerStore represents a PlayerStore that is stored in the
// file system as a JSON file
type FileSystemPlayerStore struct {
	mu       sync.Mutex
	database io.ReadWriteSeeker
}

// NewFileSystemPlayerStore returns an FileSystemPlayerStore with an empty store
func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	f := new(FileSystemPlayerStore)
	f.database = database
	return f
}

// GetPlayerScore returns a score if the player exists or an error if
// they don't
func (f *FileSystemPlayerStore) GetPlayerScore(name string) (int, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	player := f.getLeague().Find(name)

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
	league := f.getLeague()
	player := league.Find(name)

	if player == nil {
		league = append(league, Player{Name: name, Wins: 1})
	} else {
		player.Wins++
	}

	f.database.Seek(0, io.SeekStart)
	json.NewEncoder(f.database).Encode(league)
}

// GetLeague returns an ordered slice containing every Player in the league
// sorted by score, descending
func (f *FileSystemPlayerStore) GetLeague() League {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.getLeague()
}

func (f *FileSystemPlayerStore) getLeague() League {
	f.database.Seek(0, io.SeekStart)
	league, _ := NewLeague(f.database)
	return league
}

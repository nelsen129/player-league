package store

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"sync"

	"github.com/nelsen129/player-league/store/tape"
)

// FileSystemPlayerStore represents a PlayerStore that is stored in the
// file system as a JSON file
// FileSystemPlayerStore is safe for concurrent use by multiple goroutines
type FileSystemPlayerStore struct {
	mu       sync.Mutex
	database *json.Encoder
	league   League
}

// NewFileSystemPlayerStore returns an FileSystemPlayerStore with an empty store
func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initializePlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("problem initializing player db file, %v", err)
	}

	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	f := new(FileSystemPlayerStore)
	f.database = json.NewEncoder(tape.NewTape(file))
	f.league = league
	return f, nil
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
	sort.Sort(sort.Reverse(f.league))
	return f.league
}

func initializePlayerDBFile(file *os.File) error {
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("problem seeking file, %v", err)
	}

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		_, err = file.Write([]byte("[]"))
		if err != nil {
			return fmt.Errorf("problem writing file, %v", err)
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			return fmt.Errorf("problem seeking file, %v", err)
		}
	}

	return nil
}

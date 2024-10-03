package store

import (
	"errors"
)

// InMemoryPlayerStore represents a PlayerStore that is stored in memory
type InMemoryPlayerStore struct {
	store map[string]int
}

// NewInMemoryPlayerStore returns an InMemoryPlayer with an empty store
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// GetPlayerScore returns a score if the player exists or an error if
// they don't
func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, error) {
	if _, ok := i.store[name]; !ok {
		return 0, errors.New("player not found")
	}
	return i.store[name], nil
}

// RecordWin increments a player's score by 1. If the player doesn't exist,
// it will create the player and increment their score to 1.
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

package store

import (
	"sort"
	"sync"
)

// InMemoryPlayerStore represents a PlayerStore that is stored in memory
// InMemoryPlayerStore is safe for concurrent use by multiple goroutines
type InMemoryPlayerStore struct {
	mu    sync.Mutex
	store map[string]int
}

// NewInMemoryPlayerStore returns an InMemoryPlayerStore with an empty store
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	i := new(InMemoryPlayerStore)
	i.store = map[string]int{}
	return i
}

// GetPlayerScore returns a score if the player exists or an error if
// they don't
func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, error) {
	i.mu.Lock()
	defer i.mu.Unlock()
	if _, ok := i.store[name]; !ok {
		return 0, ErrPlayerNotFound
	}
	return i.store[name], nil
}

// RecordWin increments a player's score by 1. If the player doesn't exist,
// it will create the player and increment their score to 1.
func (i *InMemoryPlayerStore) RecordWin(name string) error {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.store[name]++

	return nil
}

// GetLeague returns an ordered slice containing every Player in the league
// sorted by score, descending
func (i *InMemoryPlayerStore) GetLeague() League {
	league := i.getUnsortedLeague()

	sort.Sort(sort.Reverse(league))
	return league
}

func (i *InMemoryPlayerStore) getUnsortedLeague() League {
	league := make(League, len(i.store))
	idx := 0
	i.mu.Lock()
	defer i.mu.Unlock()
	for k, v := range i.store {
		league[idx] = Player{Name: k, Wins: v}
		idx++
	}
	return league
}

package main

import (
	"errors"
)

// InMemoryPlayerStore represents a PlayerStore that is stored in
// memory. In a future release, this will be replaced with a
// persistent store
type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, error) {
	if _, ok := i.store[name]; !ok {
		return 0, errors.New("player not found")
	}
	return i.store[name], nil
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	if _, ok := i.store[name]; !ok {
		i.store[name] = 0
	}
	i.store[name]++
}

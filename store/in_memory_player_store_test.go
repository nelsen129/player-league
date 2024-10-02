package store_test

import (
	"errors"
	"testing"

	"github.com/nelsen129/player-league/store"
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
	i.store[name]++
}

func TestInMemoryPlayerStore(t *testing.T) {
	t.Run("returns an error on get if player doesn't exist", func(t *testing.T) {
		playerStore := store.NewInMemoryPlayerStore()
		_, err := playerStore.GetPlayerScore("Bob")
		if err == nil {
			t.Error("want error, got none")
		}
	})

	t.Run("records and returns a score for a new player", func(t *testing.T) {
		playerStore := store.NewInMemoryPlayerStore()
		playerStore.RecordWin("Neil")
		got, err := playerStore.GetPlayerScore("Neil")
		want := 1

		if err != nil {
			t.Errorf("want no error, got %v", err)
		}
		
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

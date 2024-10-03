package store_test

import (
	"testing"

	"github.com/nelsen129/player-league/store"
)

func TestInMemoryPlayerStore(t *testing.T) {
	playerStore := store.NewInMemoryPlayerStore()
	testStore(t, playerStore)
}

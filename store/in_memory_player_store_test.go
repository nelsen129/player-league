package store_test

import (
	"reflect"
	"testing"

	"github.com/nelsen129/player-league/store"
)

func TestInMemoryPlayerStore(t *testing.T) {
	t.Run("standard player store suite", func(t *testing.T) {
		playerStore := store.NewInMemoryPlayerStore()
		testStore(t, playerStore)
	})

	t.Run("sorts for GetLeague", func(t *testing.T) {
		playerStore := store.NewInMemoryPlayerStore()
		sortedLeague := []store.Player{
			{Name: "Alice", Wins: 5},
			{Name: "Bob", Wins: 4},
			{Name: "Charlie", Wins: 3},
			{Name: "Dave", Wins: 2},
			{Name: "Eve", Wins: 1},
		}
		unsortedLeague := []store.Player{
			{Name: "Dave", Wins: 2},
			{Name: "Bob", Wins: 4},
			{Name: "Charlie", Wins: 3},
			{Name: "Eve", Wins: 1},
			{Name: "Alice", Wins: 5},
		}

		for _, player := range unsortedLeague {
			for range player.Wins {
				playerStore.RecordWin(player.Name)
			}
		}

		got := playerStore.GetLeague()
		if !reflect.DeepEqual(got, sortedLeague) {
			t.Errorf("got %v, want %v", got, sortedLeague)
		}
	})
}

package store_test

import (
	"reflect"
	"sync"
	"testing"

	"github.com/nelsen129/player-league/store"
)

// testStore isn't run on its own. It is a helper function
// to test other stores
func testStore(t *testing.T, playerStore store.PlayerStore) {
	t.Run("returns an error on get if player doesn't exist", func(t *testing.T) {
		_, err := playerStore.GetPlayerScore("Bob")
		if err == nil {
			t.Error("want error, got none")
		}
		if err != store.ErrPlayerNotFound {
			t.Errorf("got %q, want %q", err, store.ErrPlayerNotFound)
		}
	})

	t.Run("records and returns a score for a new player", func(t *testing.T) {
		playerStore.RecordWin("Neil")
		playerStore.RecordWin("Neil")
		playerStore.RecordWin("Neil")
		got, err := playerStore.GetPlayerScore("Neil")
		want := 3

		if err != nil {
			t.Errorf("want no error, got %v", err)
		}

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("can return a league of players", func(t *testing.T) {
		wantedLeague := []store.Player{
			{"Neil", 3},
			{"Bob", 2},
		}
		playerStore.RecordWin("Bob")
		playerStore.RecordWin("Bob")
		got := playerStore.GetLeague()

		if !reflect.DeepEqual(got, wantedLeague) {
			t.Errorf("got %v, want %v", got, wantedLeague)
		}
	})

	t.Run("handles concurrent operations", func(t *testing.T) {
		winCount := 1000
		scoreCount := 1000
		player := "Karen"

		var wg sync.WaitGroup
		wg.Add(winCount)
		wg.Add(scoreCount)

		for range winCount {
			go func() {
				playerStore.RecordWin(player)
				wg.Done()
			}()
		}
		for range scoreCount {
			go func() {
				playerStore.GetPlayerScore(player)
				wg.Done()
			}()
		}
		wg.Wait()

		got, _ := playerStore.GetPlayerScore(player)

		if got != winCount {
			t.Errorf("got %d, want %d", got, winCount)
		}
	})
}

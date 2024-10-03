package store_test

import (
	"fmt"

	"github.com/nelsen129/player-league/store"
)

func ExamplePlayerServer_ServeHTTP() {
	playerStore := store.NewInMemoryPlayerStore()

	playerStore.RecordWin("Pepper")
	playerStore.RecordWin("Pepper")
	score, err := playerStore.GetPlayerScore("Pepper")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(score)
	// Output: 2
}

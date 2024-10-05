package store_test

import (
	"fmt"
	"os"

	"github.com/nelsen129/player-league/store"
)

func ExampleInMemoryPlayerStore() {
	playerStore := store.NewInMemoryPlayerStore()

	for range 2 {
		err := playerStore.RecordWin("Pepper")
		if err != nil {
			fmt.Println(err)
		}
	}

	score, err := playerStore.GetPlayerScore("Pepper")
	if err != nil {
		fmt.Println(err)
	}

	err = playerStore.RecordWin("Larry")
	if err != nil {
		fmt.Println(err)
	}

	league := playerStore.GetLeague()

	fmt.Println(score)
	fmt.Println(league)
	// Output:
	// 2
	// [{Pepper 2} {Larry 1}]
}

func ExampleFileSystemPlayerStore() {
	file, err := os.CreateTemp("", "db")
	if err != nil {
		fmt.Println(err)
	}

	_, err = file.Write([]byte(""))
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()

	playerStore, err := store.NewFileSystemPlayerStore(file)
	if err != nil {
		fmt.Println(err)
	}

	for range 2 {
		err := playerStore.RecordWin("Pepper")
		if err != nil {
			fmt.Println(err)
		}
	}

	score, err := playerStore.GetPlayerScore("Pepper")
	if err != nil {
		fmt.Println(err)
	}

	err = playerStore.RecordWin("Larry")
	if err != nil {
		fmt.Println(err)
	}

	league := playerStore.GetLeague()

	fmt.Println(score)
	fmt.Println(league)
	// Output:
	// 2
	// [{Pepper 2} {Larry 1}]
}

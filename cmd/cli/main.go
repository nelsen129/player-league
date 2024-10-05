package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nelsen129/player-league/cli"
	"github.com/nelsen129/player-league/store"
)

const dbFileName = "game.db.json"

func main() {
	playerStore, closeFunc, err := store.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer closeFunc()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	game := cli.NewCLI(playerStore, os.Stdin)
	game.PlayPoker()
}

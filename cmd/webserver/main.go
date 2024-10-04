package main

import (
	"log"
	"net/http"

	"github.com/nelsen129/player-league/server"
	"github.com/nelsen129/player-league/store"
)

const dbFileName = "game.db.json"

func main() {
	playerStore, closeFunc, err := store.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer closeFunc()

	server := server.NewPlayerServer(playerStore)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

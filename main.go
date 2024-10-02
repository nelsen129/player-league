package main

import (
	"log"
	"net/http"

	"github.com/nelsen129/player-league/server"
	"github.com/nelsen129/player-league/store"
)

func main() {
	server := server.NewPlayerServer(store.NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5000", server))
}

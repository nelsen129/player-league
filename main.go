package main

import (
	"log"
	"net/http"

	"github.com/nelsen129/player-league/server"
)

func main() {
	server := server.NewPlayerServer(NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5000", server))
}

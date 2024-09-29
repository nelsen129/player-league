package main

import (
	"log"
	"net/http"

	"github.com/nelsen129/player-league/server"
)

func main() {
	handler := http.HandlerFunc(server.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}

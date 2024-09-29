package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/nelsen129/player-league/server"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, error) {
	if name == "dummy" {
		return 0, errors.New("test")
	}
	return 123, nil
}

func main() {
	server := server.NewPlayerServer(&InMemoryPlayerStore{})
	log.Fatal(http.ListenAndServe(":5000", server))
}

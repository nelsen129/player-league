package server

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerStore records and stores the scores for players
type PlayerStore interface {
	// Should return a player's score
	GetPlayerScore(name string) (int, error)
	// Should increment a player's score
	RecordWin(name string)
}

// PlayerServer handles the HTTP routing for requests that
// interact with the PlayerStore
type PlayerServer struct {
	store PlayerStore
}

// NewPlayerServer initializes a PlayerServer with a PlayerStore
func NewPlayerServer(store PlayerStore) *PlayerServer {
	return &PlayerServer{
		store: store,
	}
}

// ServeHTTP handles the HTTP server for player requests.
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()

	router.Handle("/league", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	router.Handle("/players/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		player := strings.TrimPrefix(r.URL.Path, "/players/")

		switch r.Method {
		case http.MethodGet:
			p.getScore(w, player)
		case http.MethodPost:
			p.incrementScore(w, player)
		}
	}))

	router.ServeHTTP(w, r)
}

func (p *PlayerServer) getScore(w http.ResponseWriter, player string) {
	score, err := p.store.GetPlayerScore(player)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) incrementScore(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

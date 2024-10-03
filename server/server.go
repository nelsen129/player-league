// Package server implements the HTTP server for interacting with a league of players
package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/nelsen129/player-league/store"
)

const jsonContentType = "application/json"

// PlayerServer handles the HTTP routing for requests that
// interact with the PlayerStore
type PlayerServer struct {
	store store.PlayerStore
	http.Handler
}

// NewPlayerServer initializes a PlayerServer with a PlayerStore
func NewPlayerServer(store store.PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playerHandler))
	p.Handler = router

	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	_ = json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodGet:
		p.getScore(w, player)
	case http.MethodPost:
		p.incrementScore(w, player)
	}
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
	err := p.store.RecordWin(player)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

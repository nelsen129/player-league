package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/nelsen129/player-league/store"
)

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
	w.WriteHeader(http.StatusOK)
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
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

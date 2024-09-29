package server

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, GetPlayerScore(player))
}

func GetPlayerScore(player string) int {
	switch player {
	case "Pepper":
		return 20
	case "Floyd":
		return 10
	}
	return 0
}

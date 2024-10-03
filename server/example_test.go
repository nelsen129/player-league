package server_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/nelsen129/player-league/server"
	"github.com/nelsen129/player-league/store"
)

func ExamplePlayerServer() {
	playerStore := store.NewInMemoryPlayerStore()
	playerServer := server.NewPlayerServer(playerStore)

	postRequest, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
	getRequest, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
	leagueRequest, _ := http.NewRequest(http.MethodGet, "/league", nil)

	playerServer.ServeHTTP(httptest.NewRecorder(), postRequest)

	response := httptest.NewRecorder()
	playerServer.ServeHTTP(response, getRequest)
	fmt.Println(response.Body.String())

	response = httptest.NewRecorder()
	playerServer.ServeHTTP(response, leagueRequest)
	fmt.Println(response.Body.String())
	// Output:
	// 1
	// [{"name":"Pepper","wins":1}]
}

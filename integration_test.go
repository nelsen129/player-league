package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nelsen129/player-league/server"
	"github.com/nelsen129/player-league/store"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := store.NewInMemoryPlayerStore()
	playerServer := server.NewPlayerServer(store)
	player := "Pepper"
	wins := 3

	for range wins {
		playerServer.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	}

	response := httptest.NewRecorder()
	playerServer.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")
}

func newGetScoreRequest(player string) *http.Request {
	r, _ := http.NewRequest(http.MethodGet, "/players/"+player, nil)
	return r
}

func newPostWinRequest(player string) *http.Request {
	r, _ := http.NewRequest(http.MethodPost, "/players/"+player, nil)
	return r
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

package server_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/nelsen129/player-league/server"
	"github.com/nelsen129/player-league/store"
)

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		score: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		winCalls: []string{},
	}
	playerServer := server.NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
		assertResponseBody(t, response.Body.String(), "")
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		score:    map[string]int{},
		winCalls: []string{},
	}
	playerServer := server.NewPlayerServer(&store)

	t.Run("it records wins when POST", func(t *testing.T) {
		request := newPostWinRequest("Pepper")
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
		assertWinCalls(t, store.winCalls, []string{"Pepper"})
	})
}

func TestLeague(t *testing.T) {
	wantedLeague := []store.Player{
		{Name: "Cleo", Wins: 32},
		{Name: "Chris", Wins: 20},
		{Name: "Tiest", Wins: 14},
	}

	playerStore := StubPlayerStore{
		score:    nil,
		winCalls: nil,
		league:   wantedLeague,
	}
	playerServer := server.NewPlayerServer(&playerStore)

	request := newLeagueRequest()
	response := httptest.NewRecorder()

	playerServer.ServeHTTP(response, request)

	got := getLeagueFromResponse(t, response.Body)

	assertStatus(t, response.Code, http.StatusOK)
	assertLeague(t, got, wantedLeague)
	assertContentType(t, response, "application/json")
}

type StubPlayerStore struct {
	score    map[string]int
	winCalls []string
	league   []store.Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) (int, error) {
	score, ok := s.score[name]
	if !ok {
		return 0, errors.New("Player not found")
	}
	return score, nil
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() []store.Player {
	return s.league
}

func newGetScoreRequest(player string) *http.Request {
	r, _ := http.NewRequest(http.MethodGet, "/players/"+player, nil)
	return r
}

func newPostWinRequest(player string) *http.Request {
	r, _ := http.NewRequest(http.MethodPost, "/players/"+player, nil)
	return r
}

func newLeagueRequest() *http.Request {
	r, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return r
}

func getLeagueFromResponse(t testing.TB, body io.Reader) []store.Player {
	t.Helper()
	var league []store.Player
	err := json.NewDecoder(body).Decode(&league)
	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of player", err)
		return nil
	}

	return league
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

func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("got response content-type %q, want %q", response.Result().Header.Get("content-type"), want)
	}
}

func assertWinCalls(t testing.TB, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v calls to RecordWin want %v", got, want)
	}
}

func assertLeague(t testing.TB, got, want []store.Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v league, want %v", got, want)
	}
}

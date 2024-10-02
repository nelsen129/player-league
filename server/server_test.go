package server_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/nelsen129/player-league/server"
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

type StubPlayerStore struct {
	score    map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) (int, error) {
	score, ok := s.score[name]
	if !ok {
		return 0, errors.New("Player not found")
	}
	return score, nil
}

func (s *StubPlayerStore) RecordWin(name string) {
	if _, ok := s.score[name]; ok {
		s.score[name] = 0
	}
	s.score[name]++
	s.winCalls = append(s.winCalls, name)
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

func assertWinCalls(t testing.TB, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v calls to RecordWin want %v", got, want)
	}
}

package server_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nelsen129/player-league/server"
)

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
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

type StubPlayerStore struct {
	score map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) (int, error) {
	score, ok := s.score[name]
	if !ok {
		return 0, errors.New("Player not found")
	}
	return score, nil
}

func newGetScoreRequest(player string) *http.Request {
	r, _ := http.NewRequest(http.MethodGet, "/players/"+player, nil)
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

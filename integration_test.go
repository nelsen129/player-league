package main_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/nelsen129/player-league/server"
	"github.com/nelsen129/player-league/store"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "")
	defer cleanDatabase()
	playerStore := store.NewFileSystemPlayerStore(database)
	playerServer := server.NewPlayerServer(playerStore)
	player := "Pepper"
	wins := 3

	for range wins {
		playerServer.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	}

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		playerServer.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		playerServer.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)
		got := getLeagueFromResponse(t, response.Body)
		want := []store.Player{
			{Name: player, Wins: wins},
		}
		assertLeague(t, got, want)
	})
}

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
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

func assertLeague(t testing.TB, got, want []store.Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v league, want %v", got, want)
	}
}

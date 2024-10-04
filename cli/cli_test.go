package cli_test

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/nelsen129/player-league/cli"
	"github.com/nelsen129/player-league/store"
)

func TestCLI(t *testing.T) {
	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &StubPlayerStore{}
		cli := cli.NewCLI(playerStore, in)
		cli.PlayPoker()

		assertWins(t, playerStore.winCalls, []string{"Chris"})
	})

	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &StubPlayerStore{}
		cli := cli.NewCLI(playerStore, in)
		cli.PlayPoker()

		assertWins(t, playerStore.winCalls, []string{"Cleo"})
	})
}

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   store.League
}

func (s *StubPlayerStore) RecordWin(name string) error {
	s.winCalls = append(s.winCalls, name)
	return nil
}

func (s *StubPlayerStore) GetPlayerScore(name string) (int, error) {
	if _, ok := s.scores[name]; !ok {
		return 0, errors.New("player not found")
	}
	return s.scores[name], nil
}

func (s *StubPlayerStore) GetLeague() store.League {
	return s.league
}

func assertWins(t testing.TB, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got wrong win calls %v, want %v", got, want)
	}
}

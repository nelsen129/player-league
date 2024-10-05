package store_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/nelsen129/player-league/store"
)

func TestFileSystemPlayerStore(t *testing.T) {
	t.Run("standard player store suite", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[]`)
		defer cleanDatabase()
		playerStore, err := store.NewFileSystemPlayerStore(database)

		if err != nil {
			t.Fatalf("could not create file system player store %v", err)
		}

		testStore(t, playerStore)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()
		_, err := store.NewFileSystemPlayerStore(database)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("sorts the league", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"name": "Cleo", "wins": 10},
			{"name": "Chris", "wins": 33}]`)
		defer cleanDatabase()

		playerStore, err := store.NewFileSystemPlayerStore(database)

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		got := playerStore.GetLeague()

		want := store.League{
			{"Chris", 33},
			{"Cleo", 10},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}

		// read again
		got = playerStore.GetLeague()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("creates a store from a file", func(t *testing.T) {
		path := filepath.Join(os.TempDir(), "db")

		playerStore, closeFunc, err := store.FileSystemPlayerStoreFromFile(path)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		defer closeFunc()

		got := playerStore.GetLeague()

		want := store.League{}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("handles errors in creating store from file", func(t *testing.T) {
		path := ""
		_, closeFunc, err := store.FileSystemPlayerStoreFromFile(path)
		if err == nil {
			t.Error("expected error, got none")
		}
		if closeFunc != nil {
			defer closeFunc()
		}
	})
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	_, err = tmpfile.Write([]byte(initialData))
	if err != nil {
		t.Fatalf("could not write temp file %v", err)
	}

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

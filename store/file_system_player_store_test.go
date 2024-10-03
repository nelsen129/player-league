package store_test

import (
	"os"
	"testing"

	"github.com/nelsen129/player-league/store"
)

func TestFileSystemPlayerStore(t *testing.T) {
	t.Run("standard player store suite", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[]`)
		defer cleanDatabase()
		playerStore := store.NewFileSystemPlayerStore(database)
		testStore(t, playerStore)
	})
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
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

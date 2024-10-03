package tape_test

import (
	"io"
	"os"
	"testing"

	"github.com/nelsen129/player-league/store/tape"
)

func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()

	tape := tape.NewTape(file)

	_, err := tape.Write([]byte("abc"))
	if err != nil {
		t.Fatalf("could not write to file %v", err)
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		t.Fatalf("could not seek file %v", err)
	}

	newFileContents, _ := io.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	_, err = tmpfile.Write([]byte(initialData))
	if err != nil {
		t.Fatalf("could not write to temp file %v", err)
	}

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

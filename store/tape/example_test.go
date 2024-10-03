package tape_test

import (
	"fmt"
	"io"
	"os"

	"github.com/nelsen129/player-league/store/tape"
)

func ExampleTape() {
	file, _ := os.CreateTemp("", "db")
	file.Write([]byte("12345"))

	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()

	tape := tape.NewTape(file)

	tape.Write([]byte("abc"))

	file.Seek(0, io.SeekStart)
	newFileContents, _ := io.ReadAll(file)

	fmt.Println(string(newFileContents))
	// Output: abc
}

//func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
//	t.Helper()
//
//	tmpfile, err := os.CreateTemp("", "db")
//
//	if err != nil {
//		t.Fatalf("could not create temp file %v", err)
//	}
//
//	tmpfile.Write([]byte(initialData))
//
//	removeFile := func() {
//		tmpfile.Close()
//		os.Remove(tmpfile.Name())
//	}
//
//	return tmpfile, removeFile
//}

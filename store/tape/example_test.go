package tape_test

import (
	"fmt"
	"io"
	"os"

	"github.com/nelsen129/player-league/store/tape"
)

func ExampleTape() {
	file, _ := os.CreateTemp("", "db")
	_, err := file.Write([]byte("12345"))
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()

	tape := tape.NewTape(file)

	_, err = tape.Write([]byte("abc"))
	if err != nil {
		fmt.Println(err)
	}

	file.Seek(0, io.SeekStart)
	newFileContents, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(newFileContents))
	// Output: abc
}

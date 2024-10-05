package cli

import (
	"bufio"
	"io"
	"strings"

	"github.com/nelsen129/player-league/store"
)

type CLI struct {
	playerStore store.PlayerStore
	in          *bufio.Scanner
}

func NewCLI(playerStore store.PlayerStore, in io.Reader) *CLI {
	cli := new(CLI)
	cli.playerStore = playerStore
	cli.in = bufio.NewScanner(in)
	return cli
}

func (cli *CLI) PlayPoker() {
	userInput := cli.readLine()
	cli.playerStore.RecordWin(extractWinner(userInput))
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

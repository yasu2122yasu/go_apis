package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

const (
	TERMINAL = "terminal"
	PIPE     = "pipe"
)

func main() {
	isTerminal := terminal.IsTerminal(int(os.Stdin.Fd()))
	switch isTerminal {
	case true:
		var stdin string
		fmt.Scan(&stdin)
		fmt.Printf("type: %s, text: %s\n", TERMINAL, stdin)
		break
	case false:
		stdin, _ := ioutil.ReadAll(os.Stdin)
		fmt.Printf("type: %s, text: %s", PIPE, string(stdin))
		break
	default:
	}
}

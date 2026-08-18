package main

import (
	"op-cli/cmd"
	"op-cli/tui"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		tui.Home()
	} else {
		cmd.Execute()
	}
}

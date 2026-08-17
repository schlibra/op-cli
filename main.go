package main

import (
	"op-cli/cmd"
	"op-cli/tui"
	"op-cli/utils"
	"os"
)

func main() {
	utils.SetDNS()
	if len(os.Args) == 1 {
		tui.Home()
	} else {
		cmd.Execute()
	}
}

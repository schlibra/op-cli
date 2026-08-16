package tui

import (
	"fmt"
	"op-cli/utils"
)

func About() {
	fmt.Printf("Version:	%s\n", utils.Version)
	fmt.Printf("Git Commit:	%s\n", utils.GitCommit)
	fmt.Printf("Build Time:	%s\n", utils.BuildTime)
	Home()
}

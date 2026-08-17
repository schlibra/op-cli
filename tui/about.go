package tui

import (
	"fmt"
	"log"
	"op-cli/utils"
	"os"

	"github.com/charmbracelet/huh"
)

func About() {
	var description string
	var action string
	description += fmt.Sprintf("Version    : %s\n", utils.Version)
	description += fmt.Sprintf("Git Commit :	%s\n", utils.GitCommit)
	description += fmt.Sprintf("Build Time :	%s", utils.BuildTime)
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("About program").
				Description(description),
			huh.NewSelect[string]().
				Title("Select action").
				Options(
					huh.NewOption("Back", "back"),
					huh.NewOption("Quit", "quit"),
				).
				Value(&action),
		),
	).Run()
	if err != nil {
		log.Fatal(err)
	}
	switch action {
	case "back":
		Home()
	case "quit":
		os.Exit(0)
	default:
		About()
	}
}

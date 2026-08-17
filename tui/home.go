package tui

import (
	"log"
	"os"

	"github.com/charmbracelet/huh"
)

func Home() {
	var action string
	err := huh.NewSelect[string]().
		Title("Select Action").
		Options(
			huh.NewOption("Server setting", "server"),
			huh.NewOption("Auth setting", "auth"),
			huh.NewOption("File browser", "file"),
			huh.NewOption("About program", "about"),
			huh.NewOption("Quit program", "quit"),
		).
		Value(&action).
		Run()
	if err != nil {
		log.Fatal(err)
	}
	switch action {
	case "server":
		Server()
	case "auth":
		Auth()
	case "file":
		File("/")
	case "about":
		About()
	case "quit":
		os.Exit(0)
	default:
		Home()
	}
}

package tui

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
)

func Home() {
	var action = ""
	err := survey.AskOne(&survey.Select{
		Message: "Select Action",
		Options: []string{"Server", "Auth", "File", "Exit", "About", "Quit"},
		Default: "Server",
	}, &action)
	if err != nil {
		return
	}
	switch action {
	case "Server":
		Server()
	case "Auth":
		Auth()
	case "About":
		About()
	case "Quit":
		os.Exit(0)
	}
}

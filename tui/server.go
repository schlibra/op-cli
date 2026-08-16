package tui

import (
	"log"
	"op-cli/utils"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func setBaseURL() {
	var baseUrl = ""
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = survey.AskOne(&survey.Input{
		Message: "Input new base url: ",
	}, &baseUrl)
	if err != nil {
		log.Fatal(err)
	}
	baseUrl = strings.TrimSuffix(baseUrl, "/")
	config.BaseURL = baseUrl
	err = utils.SaveConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	baseURL()
}

func baseURL() {
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	var action = ""
	err = survey.AskOne(&survey.Select{
		Message: "Server Base URL: " + config.BaseURL,
		Options: []string{"Set URL", "Back", "Home", "Quit"},
		Default: "Back",
	}, &action)
	if err != nil {
		log.Fatal(err)
	}
	switch action {
	case "Set URL":
		setBaseURL()
	case "Back":
		Server()
	case "Home":
		Home()
	case "Quit":
		os.Exit(0)
	}
}

func Server() {
	var action = ""
	err := survey.AskOne(&survey.Select{
		Message: "Server setting",
		Options: []string{"Base URL", "Back", "Quit"},
		Default: "Back",
	}, &action)
	if err != nil {
		log.Fatal(err)
	}
	switch action {
	case "Base URL":
		baseURL()
	case "Back":
		Home()
	case "Quit":
		os.Exit(0)
	}
}

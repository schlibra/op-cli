package tui

import (
	"fmt"
	"log"
	"op-cli/utils"
	"os"

	"github.com/charmbracelet/huh"
)

func baseURL() {
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	baseUrl := config.BaseURL
	var action string
	err = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Base URL: ").
				Value(&baseUrl),
			huh.NewSelect[string]().
				Title("Base URL setting").
				Options(
					huh.NewOption("Save", "save"),
					huh.NewOption("Back", "back"),
					huh.NewOption("Home", "home"),
					huh.NewOption("Quit", "quit"),
				).
				Value(&action),
		),
	).Run()
	if err != nil {
		log.Fatal(err)
	}
	switch action {
	case "save":
		config.BaseURL = baseUrl
		err = utils.SaveConfig(config)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Save success")
		Server()
	case "back":
		Server()
	case "home":
		Home()
	case "quit":
		os.Exit(0)
	default:
		baseURL()
	}
}

func Server() {
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	var action string
	err = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Server setting").
				Options(
					huh.NewOption("BaseURL: "+config.BaseURL, "baseUrl"),
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
	case "baseUrl":
		baseURL()
	case "back":
		Home()
	case "quit":
		os.Exit(0)
	}
}

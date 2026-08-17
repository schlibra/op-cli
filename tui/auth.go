package tui

import (
	"fmt"
	"log"
	"op-cli/utils"
	"os"

	"github.com/charmbracelet/huh"
)

func userLogin() {
	var (
		username = ""
		password = ""
		totp     = ""
	)
	var action string
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Username: ").
				Value(&username),
			huh.NewInput().
				EchoMode(huh.EchoModePassword).
				Title("Password: ").
				Value(&password),
			huh.NewInput().
				Title("Totp(If enabled): ").
				Value(&totp),
			huh.NewSelect[string]().
				Title("Select action: ").
				Options(
					huh.NewOption("Login user", "login"),
					huh.NewOption("Back", "back"),
					huh.NewOption("Home", "home"),
					huh.NewOption("Quit", "quit"),
				).
				Value(&action),
		).Title("Login user"),
	).Run()
	if err != nil {
		log.Fatal(err)
	}
	switch action {
	case "login":
		data := utils.SendUserLogin(username, password, totp)
		if data.Code == 200 {
			fmt.Println("Login success")
			token := data.Data.Token
			config, err := utils.LoadConfig()
			if err != nil {
				log.Fatal(err)
			}
			config.Token = token
			err = utils.SaveConfig(config)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("Login failed: " + data.Message)
		}
		Auth()
	case "back":
		Auth()
	case "home":
		Home()
	case "quit":
		os.Exit(0)
	default:
		userLogin()
	}

}
func userInfo() {
	userInfoData := utils.SendUserInfo()
	if userInfoData != nil {
		if userInfoData.Code == 200 {
			fmt.Println("Username: " + userInfoData.Data.Username)
			fmt.Printf("Role: %d\n", userInfoData.Data.Role)
		} else {
			fmt.Println("User info get error: " + userInfoData.Message)
		}
	}
	Auth()
}
func userLogout() {
	utils.SendUserLogout()
	fmt.Println("Logout success")
	Auth()
}

func Auth() {
	var action string
	err := huh.NewSelect[string]().
		Title("Auth setting").
		Options(
			huh.NewOption("Login user", "login"),
			huh.NewOption("Logout user", "logout"),
			huh.NewOption("User info", "info"),
			huh.NewOption("Back", "back"),
			huh.NewOption("Quit", "quit"),
		).
		Value(&action).
		Run()
	if err != nil {
		log.Fatal(err)
	}
	switch action {
	case "login":
		userLogin()
	case "logout":
		userLogout()
	case "info":
		userInfo()
	case "back":
		Home()
	case "quit":
		os.Exit(0)
	default:
		Auth()
	}
}

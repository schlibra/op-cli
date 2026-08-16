package tui

import (
	"fmt"
	"log"
	"op-cli/utils"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

func userLogin() {
	var (
		username = ""
		password = ""
		totp     = ""
	)
	err := survey.AskOne(&survey.Input{
		Message: "Input username:",
	}, &username)
	if err != nil {
		log.Fatal(err)
	}
	err = survey.AskOne(&survey.Password{
		Message: "Input password:",
	}, &password)
	if err != nil {
		log.Fatal(err)
	}
	err = survey.AskOne(&survey.Input{
		Message: "Input TOTP Code(If enabled):",
	}, &totp)
	if err != nil {
		log.Fatal(err)
	}
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
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	if config.BaseURL == "" {
		fmt.Println("Base URL not set")
		Home()
		return
	}
	var action = ""
	err = survey.AskOne(&survey.Select{
		Message: "Auth setting ",
		Options: []string{"Login", "Logout", "Info", "Back", "Quit"},
		Default: "Login",
	}, &action)
	if err != nil {
		log.Fatal(err)
	}
	switch action {
	case "Login":
		userLogin()
	case "Logout":
		userLogout()
	case "Info":
		userInfo()
	case "Back":
		Home()
	case "Quit":
		os.Exit(0)
	}
}

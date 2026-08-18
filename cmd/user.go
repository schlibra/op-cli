package cmd

import (
	"fmt"
	"log"
	"op-cli/utils"
	"os"

	"github.com/spf13/cobra"
)

var UserCmd = &cobra.Command{
	Use:   "user",
	Short: "User actions",
}

var UserInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show user info",
	Run: func(cmd *cobra.Command, args []string) {
		userInfo()
	},
}

var UserLoginCmd = &cobra.Command{
	Use:   "login <username> <password> [totp]",
	Short: "Login to your user",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			username string
			password string
			totp     string
		)
		username = args[0]
		password = args[1]
		if len(args) > 2 {
			totp = args[2]
		}
		userLogin(username, password, totp)
	},
}

var UserLogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout your user",
	Run: func(cmd *cobra.Command, args []string) {
		userLogout()
	},
}

func userLogin(username string, password string, totp string) {
	if len(os.Args) >= 4 {
		data := utils.SendUserLogin(username, password, totp)
		if data.Code == 200 {
			config, err := utils.LoadConfig()
			if err != nil {
				log.Fatal(err)
			}
			config.Token = data.Data.Token
			err = utils.SaveConfig(config)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Login successful")
		} else {
			fmt.Println("Login failed " + data.Message)
		}
	} else {
		fmt.Println("Login to your user\n\tusage: login <username> <password>")
	}
}

func userInfo() {
	data := utils.SendUserInfo()
	if data.Code == 200 {
		fmt.Printf("Username:\t%s\nBasePath:\t%s\nRole:\t\t%d\nPermission:\t%d\n", data.Data.Username, data.Data.BasePath, data.Data.Role, data.Data.Permission)
	} else {
		fmt.Printf("User info get error: %s\n", data.Message)
	}
}

func userLogout() {
	utils.SendUserLogout()
	fmt.Println("Logout success")
}

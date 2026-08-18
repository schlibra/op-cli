package cmd

import (
	"fmt"
	"log"
	"op-cli/utils"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var UrlCmd = &cobra.Command{
	Use:   "url",
	Short: "Set or get server base-url",
}
var UrlGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get server base-url",
	Run: func(cmd *cobra.Command, args []string) {
		urlGet()
	},
}
var UrlSetCmd = &cobra.Command{
	Use:   "set <url>",
	Short: "Set server base-url",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		urlSet(args)
	},
}

func url() {
	if len(os.Args) >= 3 {
		baseUrl := os.Args[2]
		config, err := utils.LoadConfig()
		if err != nil {
			log.Fatal(err)
		}
		if strings.ToLower(baseUrl) == "get" || strings.ToLower(baseUrl) == "info" {
			fmt.Println("Current base-url is: " + config.BaseURL)
			os.Exit(0)
		}
		config.BaseURL = baseUrl
		err = utils.SaveConfig(config)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("BaseURL successfully set to " + baseUrl)
	} else {
		fmt.Println("Set server base-url\n\tusage: url <base-url>")
	}
}

func urlGet() {
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current base-url is: " + config.BaseURL)
}

func urlSet(args []string) {
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	config.BaseURL = args[0]
	err = utils.SaveConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("BaseURL successfully set to " + config.BaseURL)
}

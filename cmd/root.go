package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func Execute() {
	var rootCmd = &cobra.Command{
		Use:   filepath.Base(os.Args[0]),
		Short: "op-cli",
		Long:  "Openlist cli tool",
	}
	UrlCmd.AddCommand(UrlGetCmd)
	UrlCmd.AddCommand(UrlSetCmd)
	rootCmd.AddCommand(UrlCmd)
	UserCmd.AddCommand(UserLoginCmd)
	UserCmd.AddCommand(UserLogoutCmd)
	UserCmd.AddCommand(UserInfoCmd)
	rootCmd.AddCommand(UserCmd)
	rootCmd.AddCommand(VersionCmd)
	rootCmd.AddCommand(ListCmd)
	rootCmd.AddCommand(FileGetCmd)
	rootCmd.AddCommand(DownloadCmd)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

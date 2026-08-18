package cmd

import (
	"fmt"
	"op-cli/utils"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		version()
	},
}

func version() {
	fmt.Printf("Version:	%s\n", utils.Version)
	fmt.Printf("Git Commit:	%s\n", utils.GitCommit)
	fmt.Printf("Build Time:	%s\n", utils.BuildTime)
}

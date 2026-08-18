package cmd

import (
	"fmt"
	"op-cli/utils"
	"os"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "ls <path>",
	Short: "List files and directories",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pathname := ""
		if len(args) > 0 {
			pathname = args[0]
		}
		ls(pathname)
	},
}

func ls(pathname string) {
	if len(os.Args) >= 3 {
		pathname = os.Args[2]
	} else if len(os.Args) == 2 {
		pathname = "/"
	}
	data := utils.SendListFilePath(pathname)
	if data.Code == 200 {
		for _, item := range data.Data.Content {
			if item.IsDir {
				fmt.Printf("Folder:\t%s\n", item.Name)
			} else {
				fileSize := utils.FormatFileSize(item.Size)
				fmt.Printf("File:\t%s\t%s\n", item.Name, fileSize)
			}
		}
	} else {
		fmt.Println("File path get error: " + data.Message)
	}
}

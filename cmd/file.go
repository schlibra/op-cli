package cmd

import (
	"fmt"
	"op-cli/utils"

	"github.com/spf13/cobra"
)

var FileGetCmd = &cobra.Command{
	Use:   "file <path>",
	Short: "Get File",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fileGet(args[0])
	},
}

func fileGet(pathname string) {
	data := utils.SendGetFilePath(pathname)
	if data.Code == 200 {
		fmt.Println("Name:\t\t" + data.Data.Name)
		fmt.Println("Size:\t\t" + utils.FormatFileSize(data.Data.Size))
		fmt.Println("Type:\t\t" + []string{"Unknown", "Folder", "Video", "Audio", "Text", "Image"}[data.Data.Type])
		fmt.Println("Created:\t" + data.Data.Created.String())
		fmt.Println("Modified:\t" + data.Data.Modified.String())
		fmt.Println("URL:\t\t" + data.Data.RawUrl)
	} else {
		fmt.Println("File path get error: " + data.Message)
	}
}

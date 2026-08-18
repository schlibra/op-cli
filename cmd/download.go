package cmd

import (
	"fmt"
	"log"
	"op-cli/utils"

	"github.com/spf13/cobra"
)

var DownloadCmd = &cobra.Command{
	Use:   "download <path>",
	Short: "Download a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		download(args[0])
	},
}

func download(pathname string) {
	data := utils.SendGetFilePath(pathname)
	if data.Code == 200 {
		filename := data.Data.Name
		downUrl := data.Data.RawUrl
		err := utils.DownloadWithProgressBar(downUrl, filename)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Download finished " + filename)
	} else {
		fmt.Println("File path get error: " + data.Message)
	}
}

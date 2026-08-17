package tui

import (
	"fmt"
	"log"
	"op-cli/utils"
	"os"
	"path"

	"github.com/charmbracelet/huh"
)

type fileSelectItem struct {
	Action   string
	Filename string
}

func downloadFile(filename string, downUrl string, callbackFilename string) {
	err := utils.DownloadWithProgressBar(downUrl, filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Download finished: " + filename)
	fileInfo(callbackFilename)
}

func fileInfo(filepath string) {
	data := utils.SendGetFilePath(filepath)
	var action = "down"
	var description string
	if data.Code == 200 {
		description = fmt.Sprintf("Name    : %s\n", data.Data.Name)
		description += fmt.Sprintf("Size    : %s\n", utils.FormatFileSize(data.Data.Size))
		description += fmt.Sprintf("Type    : %s\n", []string{"Unknown", "Folder", "Video", "Audio", "Text", "Image"}[data.Data.Type])
		description += fmt.Sprintf("Created : %s\n", data.Data.Created.String())
		description += fmt.Sprintf("Modified: %s", data.Data.Modified.String())
	} else {
		description = data.Message
	}
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("File Info").
				Description(description),
			huh.NewSelect[string]().
				Title("File action").
				Options(
					huh.NewOption("Download", "down"),
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
	case "down":
		downloadFile(data.Data.Name, data.Data.RawUrl, filepath)
		fileInfo(filepath)
	case "back":
		File(path.Dir(filepath))
	case "home":
		Home()
	case "quit":
		os.Exit(0)
	default:
		fileInfo(filepath)
	}

}

func File(filepath string) {
	fileListData := utils.SendListFilePath(filepath)
	var action fileSelectItem
	options := huh.NewOptions[fileSelectItem]()
	if filepath != "/" {
		options = append(options, huh.NewOption("[..]   Parent folder", fileSelectItem{Action: "parent"}))
	}
	if fileListData.Code == 200 {
		for _, item := range fileListData.Data.Content {
			if item.IsDir {
				options = append(options, huh.NewOption("[Dir]  "+item.Name, fileSelectItem{Action: "dir", Filename: item.Name}))
			} else {
				options = append(options, huh.NewOption("[File] "+item.Name, fileSelectItem{Action: "file", Filename: item.Name}))
			}
		}
	} else {
		options = append(options, huh.NewOption(fileListData.Message, fileSelectItem{Action: "none"}))
	}
	options = append(options, huh.NewOption("Back", fileSelectItem{Action: "back"}))
	options = append(options, huh.NewOption("Quit", fileSelectItem{Action: "quit"}))
	err := huh.NewSelect[fileSelectItem]().
		Title(fmt.Sprintf("File browser: [%s]", filepath)).
		Options(options...).
		Value(&action).
		Run()
	if err != nil {
		log.Fatal(err)
	}
	switch action.Action {
	case "parent":
		File(path.Dir(filepath))
	case "dir":
		File(path.Join(filepath, action.Filename))
	case "file":
		fileInfo(path.Join(filepath, action.Filename))
	case "back":
		Home()
	case "quit":
		os.Exit(0)
	default:
		File(filepath)
	}
}

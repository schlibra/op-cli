package cmd

import (
	"fmt"
	"log"
	"op-cli/utils"
	"os"
	"path/filepath"
	"strings"
)

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
func login() {
	if len(os.Args) >= 4 {
		username := os.Args[2]
		password := os.Args[3]
		totp := ""
		if len(os.Args) >= 5 {
			totp = os.Args[4]
		}
		data := utils.SendUserLogin(username, password, totp)
		if data != nil {
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
		}
	} else {
		fmt.Println("Login to your user\n\tusage: login <username> <password>")
	}
}
func info() {
	data := utils.SendUserInfo()
	if data.Code == 200 {
		fmt.Printf("Username:\t%s\nBasePath:\t%s\nRole:\t\t%d\nPermission:\t%d\n", data.Data.Username, data.Data.BasePath, data.Data.Role, data.Data.Permission)
	} else {
		fmt.Printf("User info get error: %s\n", data.Message)
	}
}
func logout() {
	utils.SendUserLogout()
	fmt.Println("Logout success")
}
func help() {
	filename := filepath.Base(os.Args[0])
	fmt.Printf("Usage: %s [options]\n\turl\t<base-url>\t\tSet server base-url\n\tlogin\t<username> <password>\tLogin to your user\n\tinfo\t\t\t\tGet current login user info\n\tlogout\t\t\t\tLogout current user\n\tls\t<path>\t\t\tList file in a path\n\tget\t<path>\t\t\tGet file or dir info\n\tdownload\t<path>\t\tDownload a file\n\tversion\t\t\t\tShow program version\n\thelp\t\t\t\tShow this help menu", filename)
}

func ls() {
	pathname := ""
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
func get() {
	pathname := ""
	if len(os.Args) >= 3 {
		pathname = os.Args[2]
	} else if len(os.Args) == 2 {
		pathname = "/"
	}
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
func download() {
	pathname := ""
	if len(os.Args) >= 3 {
		pathname = os.Args[2]
	} else if len(os.Args) == 2 {
		pathname = "/"
	}
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
func version() {
	fmt.Printf("Version:	%s\n", utils.Version)
	fmt.Printf("Git Commit:	%s\n", utils.GitCommit)
	fmt.Printf("Build Time:	%s\n", utils.BuildTime)
}
func Execute() {
	arg1 := strings.ToLower(os.Args[1])
	switch arg1 {
	case "url", "-url", "--url":
		url()
	case "login", "-login", "--login":
		login()
	case "logout", "-logout", "--logout":
		logout()
	case "info", "-info", "--info":
		info()
	case "ls", "-ls", "--ls":
		ls()
	case "get", "-get", "--get":
		get()
	case "download", "dl", "down", "--download", "--dl", "--down", "-download", "-dl", "-down":
		download()
	case "version", "-version", "--version", "-v":
		version()
	case "help", "-help", "--help":
		help()
	}
}

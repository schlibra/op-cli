package utils

import (
	"fmt"
	"log"
	"op-cli/model"

	"resty.dev/v3"
)

func deferCloseClient(client *resty.Client) {
	defer func(client *resty.Client) {
		err := client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(client)
}

func e(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getBaseUrl() string {
	config, err := LoadConfig()
	e(err)
	if config.BaseURL == "" {
		fmt.Println("Base URL not set")
	}
	return config.BaseURL
}
func getToken() string {
	config, err := LoadConfig()
	e(err)
	return config.Token
}

func SendUserLogin(username string, password string, totp string) model.UserLoginResponse {
	baseUrl := getBaseUrl()
	client := resty.New()
	deferCloseClient(client)
	var result model.UserLoginResponse
	_, err := client.R().
		SetBody(model.UserLoginRequest{
			Username: username,
			Password: password,
			Totp:     totp,
		}).
		SetResult(&result).
		Post(baseUrl + "/api/auth/login")
	e(err)
	return result
}
func SendUserInfo() model.UserInfo {
	baseUrl := getBaseUrl()
	client := resty.New()
	deferCloseClient(client)
	var result model.UserInfo
	_, err := client.R().
		SetHeader("Authorization", getToken()).
		SetResult(&result).
		Get(baseUrl + "/api/me")
	e(err)
	return result
}
func SendUserLogout() {
	baseUrl := getBaseUrl()
	client := resty.New()
	deferCloseClient(client)
	_, err := client.R().
		SetHeader("Authorization", getToken()).
		Post(baseUrl + "/api/auth/logout")
	e(err)
	config, err := LoadConfig()
	e(err)
	config.Token = ""
	err = SaveConfig(config)
	e(err)
}
func SendListFilePath(filePath string) model.FileListResponse {
	baseUrl := getBaseUrl()
	client := resty.New()
	deferCloseClient(client)
	var result model.FileListResponse
	_, err := client.R().
		SetBody(model.FileListRequest{
			Path: filePath,
		}).
		SetHeader("Authorization", getToken()).
		SetResult(&result).
		Post(baseUrl + "/api/fs/list")
	e(err)
	return result
}
func SendGetFilePath(filePath string) model.FileInfoResponse {
	baseUrl := getBaseUrl()
	client := resty.New()
	deferCloseClient(client)
	var result model.FileInfoResponse
	_, err := client.R().
		SetBody(model.FileInfoRequest{
			Path: filePath,
		}).
		SetHeader("Authorization", getToken()).
		SetResult(&result).
		Post(baseUrl + "/api/fs/get")
	e(err)
	return result
}

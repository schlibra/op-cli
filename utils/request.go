package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"op-cli/model"
)

type authTransport struct {
	token     string
	transport http.RoundTripper
}

func (t *authTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	req.Header.Set("Authorization", t.token)
	return t.transport.RoundTrip(req)
}

func getBaseUrl() string {
	config, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	if config.BaseURL == "" {
		fmt.Println("Base URL not set")
	}
	return config.BaseURL
}
func getToken() string {
	config, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	return config.Token
}

func setToken(token string) {
	http.DefaultClient.Transport = &authTransport{
		token:     token,
		transport: http.DefaultTransport,
	}
}

func SendUserLogin(username string, password string, totp string) *model.UserLoginResponse {
	baseUrl := getBaseUrl()
	authData := model.UserLoginRequest{
		Username: username,
		Password: password,
	}
	if totp != "" {
		authData.Totp = totp
	}
	jsonData, err := json.Marshal(authData)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post(baseUrl+"/api/auth/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		fmt.Println("Successfully logged in")
		var result model.UserLoginResponse
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			log.Fatal(err)
		}
		return &result
	}
	fmt.Println("Failed to login")
	return nil
}
func SendUserInfo() *model.UserInfo {
	baseUrl := getBaseUrl()
	setToken(getToken())
	resp, err := http.Get(baseUrl + "/api/me")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		var result model.UserInfo
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			log.Fatal(err)
		}
		return &result
	}
	fmt.Println("Failed to get user info")
	return nil
}
func SendUserLogout() {
	baseUrl := getBaseUrl()
	setToken(getToken())
	_, _ = http.Post(baseUrl+"/api/auth/logout", "application/json", nil)
	config, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	config.Token = ""
	_ = SaveConfig(config)
}
func SendListFilePath(filePath string) model.FileListResponse {
	baseUrl := getBaseUrl()
	setToken(getToken())
	fileData := model.FileListRequest{
		Path: filePath,
	}
	jsonData, err := json.Marshal(fileData)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post(baseUrl+"/api/fs/list", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var result model.FileListResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}
	return result
}
func SendGetFilePath(filePath string) model.FileInfoResponse {
	baseUrl := getBaseUrl()
	setToken(getToken())
	fileData := model.FileInfoRequest{
		Path: filePath,
	}
	jsonData, err := json.Marshal(fileData)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post(baseUrl+"/api/fs/get", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var result model.FileInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}
	return result
}

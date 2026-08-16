package model

import "time"

type FileListRequest struct {
	Path     string `json:"path"`
	Password string `json:"password"`
	Refresh  bool   `json:"refresh"`
	Page     int    `json:"page"`
	PerPage  int    `json:"per_page"`
}
type fileListItem struct {
	Name     string    `json:"name"`
	Size     int64     `json:"size"`
	IsDir    bool      `json:"is_dir"`
	Modified time.Time `json:"modified"`
	Created  time.Time `json:"created"`
	Sign     string    `json:"sign"`
	Thumb    string    `json:"thumb"`
	Type     int       `json:"type"`
	HashInfo string    `json:"hashinfo"`
}
type fileListResponseData struct {
	Content            []fileListItem `json:"content"`
	Total              int64          `json:"total"`
	Readme             string         `json:"readme"`
	Header             string         `json:"header"`
	Write              bool           `json:"write"`
	WriteContentBypass bool           `json:"write_content_bypass"`
	Provider           string         `json:"provider"`
}
type FileListResponse struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    fileListResponseData `json:"data"`
}
type FileInfoRequest struct {
	Path     string `json:"path"`
	Password string `json:"password"`
}
type fileInfoResponseData struct {
	Id       string    `json:"id"`
	Path     string    `json:"path"`
	Name     string    `json:"name"`
	Size     int64     `json:"size"`
	IsDir    bool      `json:"is_dir"`
	Modified time.Time `json:"modified"`
	Created  time.Time `json:"created"`
	Sign     string    `json:"sign"`
	Thumb    string    `json:"thumb"`
	Type     int       `json:"type"`
	HashInfo string    `json:"hashinfo"`
	RawUrl   string    `json:"raw_url"`
}
type FileInfoResponse struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    fileInfoResponseData `json:"data"`
}

package model

type userInfoData struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	BasePath   string `json:"base_path"`
	Role       int    `json:"role"`
	Disabled   bool   `json:"disabled"`
	Permission int    `json:"permission"`
	SSOId      string `json:"sso_id"`
	Otp        bool   `json:"otp"`
}
type UserInfo struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    userInfoData `json:"data"`
}

package model

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Totp     string `json:"otp_code"`
}

type userLoginResponseData struct {
	Token string `json:"token"`
}

type UserLoginResponse struct {
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Data    userLoginResponseData `json:"data"`
}

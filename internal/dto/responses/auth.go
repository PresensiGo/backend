package responses

type LoginResponse struct {
	Token       string `json:"token"`
	AccessToken string `json:"access_token"`
}

type RegisterResponse struct {
	Token       string `json:"token"`
	AccessToken string `json:"access_token"`
}

type RefreshTokenResponse struct {
	Token       string `json:"token"`
	AccessToken string `json:"access_token"`
}

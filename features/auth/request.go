package auth

type LoginRequest struct {
	Email    string `json:"email" default:"email@email.com"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name" default:"John Doe"`
	Email    string `json:"email" default:"email@email.com"`
	Password string `json:"password" default:"password"`
}

type RefreshTokenRequest struct {
	AccessToken string `json:"access_token"`
}

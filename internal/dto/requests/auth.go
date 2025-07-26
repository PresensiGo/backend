package requests

type Login struct {
	Email    string `json:"email" default:"email@email.com"`
	Password string `json:"password"`
}

type Register struct {
	Name     string `json:"name" default:"John Doe"`
	Email    string `json:"email" default:"email@email.com"`
	Password string `json:"password" default:"password"`
}

type Logout struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenTTL struct {
	RefreshToken string `json:"refresh_token"`
} // @name RefreshTokenTTLReq

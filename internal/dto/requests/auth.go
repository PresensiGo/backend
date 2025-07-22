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

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

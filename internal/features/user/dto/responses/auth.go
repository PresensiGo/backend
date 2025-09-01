package responses

type Login struct {
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
} // @name LoginRes

type Login2 struct {
	Token string `json:"token" validate:"required"`
} // @name Login2Res

type Register struct {
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type Logout struct {
	Message string `json:"message" validate:"required"`
}

type Logout2 struct {
	Message string `json:"message" validate:"required"`
}

type RefreshToken struct {
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
} // @name RefreshTokenRes

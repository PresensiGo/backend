package requests

type Login struct {
	Email    string `json:"email" default:"email@email.com"`
	Password string `json:"password"`
} // @name LoginReq

type Login2 struct {
	Email    string `json:"email" default:"email@email.com"`
	Password string `json:"password"`
} // @name Login2Req

// type Register struct {
// 	Code     string `json:"code"`
// 	Name     string `json:"name" default:"John Doe"`
// 	Email    string `json:"email" default:"email@email.com"`
// 	Password string `json:"password" default:"password"`
// }

type Logout struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
} // @name LogoutReq

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
} // @name RefreshTokenReq

type RefreshTokenTTL struct {
	RefreshToken string `json:"refresh_token"`
} // @name RefreshTokenTTLReq

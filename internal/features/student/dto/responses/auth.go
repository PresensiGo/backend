package responses

type LoginStudent struct {
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
} // @name LoginStudentRes

type RefreshTokenStudent struct {
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
} // @name RefreshTokenStudentRes

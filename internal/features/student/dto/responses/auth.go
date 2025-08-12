package responses

type LoginStudent struct {
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
} // @name LoginStudentRes

// todo: ganti menjadi refresh student token
type RefreshTokenStudent struct {
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
} // @name RefreshTokenStudentRes

type EjectStudentToken struct {
	Message string `json:"message" validate:"required"`
}

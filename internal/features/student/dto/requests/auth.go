package requests

type LoginStudent struct {
	SchoolCode string `json:"school_code" validate:"required"`
	NIS        string `json:"nis" validate:"required"`
	DeviceId   string `json:"device_id" validate:"required"`
} // @name LoginStudentReq

type RefreshTokenStudent struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
} // @name RefreshTokenStudentReq

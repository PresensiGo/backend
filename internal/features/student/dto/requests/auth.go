package requests

type StudentLogin struct {
	SchoolCode string `json:"school_code" validate:"required"`
	NIS        string `json:"nis" validate:"required"`
	DeviceId   string `json:"device_id" validate:"required"`
}

type StudentRefreshToken struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

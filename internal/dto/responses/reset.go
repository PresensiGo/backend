package responses

type Reset struct {
	Message string `json:"message" validate:"required"`
}

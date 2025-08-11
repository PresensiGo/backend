package responses

type ImportData struct {
	Message string `json:"message" validate:"required"`
}

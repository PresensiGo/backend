package responses

type ImportTeacher struct {
	Message string `json:"message" validate:"required"`
}

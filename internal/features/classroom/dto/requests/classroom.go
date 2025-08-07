package requests

type CreateClassroom struct {
	Name string `json:"name" validate:"required"`
}

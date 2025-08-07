package requests

type CreateClassroom struct {
	Name string `json:"name" validate:"required"`
}

type UpdateClassroom struct {
	Name string `json:"name" validate:"required"`
}

package responses

type Error struct {
	Message string `json:"message" validate:"required"`
}

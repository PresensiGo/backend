package handlers

import (
	"os"

	"api/internal/features/user/services"
)

type Admin struct {
	service *services.Admin
}

func NewAdmin(service *services.Admin) *Admin {
	return &Admin{
		service: service,
	}
}

func (h *Admin) Inject() error {
	return h.service.Inject(
		os.Getenv("ADMIN_SCHOOL_NAME"), os.Getenv("ADMIN_SCHOOL_CODE"), os.Getenv("ADMIN_NAME"),
		os.Getenv("ADMIN_EMAIL"), os.Getenv("ADMIN_PASSWORD"),
	)
}

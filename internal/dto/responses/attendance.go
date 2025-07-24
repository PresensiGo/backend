package responses

import "api/internal/dto"

type GetAllAttendances struct {
	Attendances []dto.Attendance `json:"attendances"`
}

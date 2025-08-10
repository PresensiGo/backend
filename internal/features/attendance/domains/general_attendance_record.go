package domains

import "api/internal/features/attendance/models"

type GeneralAttendanceRecord struct {
	Id                  uint `json:"id" validate:"required"`
	GeneralAttendanceId uint `json:"general_attendance_id" validate:"required"`
	StudentId           uint `json:"student_id" validate:"required"`
}

func FromGeneralAttendanceRecordModel(m *models.GeneralAttendanceRecord) *GeneralAttendanceRecord {
	return &GeneralAttendanceRecord{
		Id:                  m.ID,
		GeneralAttendanceId: m.GeneralAttendanceId,
		StudentId:           m.StudentId,
	}
}

func (g *GeneralAttendanceRecord) ToModel() *models.GeneralAttendanceRecord {
	return &models.GeneralAttendanceRecord{
		GeneralAttendanceId: g.GeneralAttendanceId,
		StudentId:           g.StudentId,
	}
}

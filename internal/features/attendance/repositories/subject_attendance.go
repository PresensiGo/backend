package repositories

import (
	"api/internal/features/attendance/domains"
	"api/internal/features/attendance/models"
	"gorm.io/gorm"
)

type SubjectAttendance struct {
	db *gorm.DB
}

func NewSubjectAttendance(db *gorm.DB) *SubjectAttendance {
	return &SubjectAttendance{
		db: db,
	}
}

func (r *SubjectAttendance) Create(data domains.SubjectAttendance) (
	*domains.SubjectAttendance, error,
) {
	subjectAttendance := data.ToModel()
	if err := r.db.Create(&subjectAttendance).Error; err != nil {
		return nil, err
	}

	return domains.FromSubjectAttendanceModel(subjectAttendance), nil
}

func (r *SubjectAttendance) GetManyByClassroomIds(classroomIds []uint) (
	*[]domains.SubjectAttendance, error,
) {
	var subjectAttendances []models.SubjectAttendance
	if err := r.db.Where(
		"classroom_id in ?", classroomIds,
	).Order("date_time desc").Find(&subjectAttendances).Error; err != nil {
		return nil, err
	}

	result := make([]domains.SubjectAttendance, len(subjectAttendances))
	for i, v := range subjectAttendances {
		result[i] = *domains.FromSubjectAttendanceModel(&v)
	}

	return &result, nil
}

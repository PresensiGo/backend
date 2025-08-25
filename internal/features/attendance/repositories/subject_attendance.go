package repositories

import (
	"time"

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

func (r *SubjectAttendance) GetAllByClassroomId(classroomIds uint) (
	*[]domains.SubjectAttendance, error,
) {
	var subjectAttendances []models.SubjectAttendance
	if err := r.db.Where(
		"classroom_id = ?", classroomIds,
	).Order("date_time desc").Find(&subjectAttendances).Error; err != nil {
		return nil, err
	}

	result := make([]domains.SubjectAttendance, len(subjectAttendances))
	for i, v := range subjectAttendances {
		result[i] = *domains.FromSubjectAttendanceModel(&v)
	}

	return &result, nil
}

func (r *SubjectAttendance) GetAllTodayByClassroomId(classroomIds uint) (
	*[]domains.SubjectAttendance, error,
) {
	var subjectAttendances []models.SubjectAttendance
	if err := r.db.Where(
		"classroom_id = ? and date(date_time) = ?",
		classroomIds, time.Now().Format("2006-01-02"),
	).Order("date_time desc").Find(&subjectAttendances).Error; err != nil {
		return nil, err
	}

	result := make([]domains.SubjectAttendance, len(subjectAttendances))
	for i, v := range subjectAttendances {
		result[i] = *domains.FromSubjectAttendanceModel(&v)
	}

	return &result, nil
}

func (r *SubjectAttendance) GetAllBySubjectIdBetween(
	subjectId uint, startDate time.Time, endDate time.Time,
) (
	*[]domains.SubjectAttendance, error,
) {
	var attendances []models.SubjectAttendance
	if err := r.db.Where(
		"subject_id = ? and date_time BETWEEN ? AND ?",
		subjectId, startDate, endDate,
	).Order("date_time asc").Find(&attendances).Error; err != nil {
		return nil, err
	} else {
		result := make([]domains.SubjectAttendance, len(attendances))
		for i, v := range attendances {
			result[i] = *domains.FromSubjectAttendanceModel(&v)
		}
		return &result, nil
	}
}

// func (r *SubjectAttendance) GetAllByClassroomIdSubjectIdBetween(
// 	classroomId uint, subjectId uint,
// 	startDate time.Time, endDate time.Time,
// ) (*[]domains.SubjectAttendance, error) {
// 	var attendances []models.SubjectAttendance
// 	if err := r.db.Where(
// 		"classroom_id = ? and subject_id = ? and date_time BETWEEN ? AND ?",
// 		classroomId, subjectId, startDate, endDate,
// 	).Order("date_time asc").Find(&attendances).Error; err != nil {
// 		return nil, err
// 	} else {
// 		result := make([]domains.SubjectAttendance, len(attendances))
// 		for i, v := range attendances {
// 			result[i] = *domains.FromSubjectAttendanceModel(&v)
// 		}
// 		return &result, nil
// 	}
// }

func (r *SubjectAttendance) Get(subjectAttendanceId uint) (
	*domains.SubjectAttendance, error,
) {
	var subjectAttendance models.SubjectAttendance
	if err := r.db.Where(
		"id = ?", subjectAttendanceId,
	).First(&subjectAttendance).Error; err != nil {
		return nil, err
	}

	return domains.FromSubjectAttendanceModel(&subjectAttendance), nil
}

func (r *SubjectAttendance) GetByCode(code string) (*domains.SubjectAttendance, error) {
	var subjectAttendance models.SubjectAttendance
	if err := r.db.Where("code = ?", code).First(&subjectAttendance).Error; err != nil {
		return nil, err
	} else {
		return domains.FromSubjectAttendanceModel(&subjectAttendance), nil
	}
}

func (r *SubjectAttendance) Delete(subjectAttendanceId uint) error {
	return r.db.Where(
		"id = ?", subjectAttendanceId,
	).Unscoped().Delete(&models.SubjectAttendance{}).Error
}

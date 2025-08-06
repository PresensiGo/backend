package repositories

import (
	"api/internal/features/attendance/domains"
	"api/internal/features/attendance/models"
	"gorm.io/gorm"
)

type GeneralAttendance struct {
	db *gorm.DB
}

func NewGeneralAttendance(db *gorm.DB) *GeneralAttendance {
	return &GeneralAttendance{db: db}
}

func (r *GeneralAttendance) Create(data domains.GeneralAttendance) (
	*domains.GeneralAttendance, error,
) {
	generalAttendance := data.ToModel()
	if err := r.db.Create(&generalAttendance).Error; err != nil {
		return nil, err
	}

	return domains.FromGeneralAttendanceModel(generalAttendance), nil
}

func (r *GeneralAttendance) GetAllBySchoolId(schoolId uint) (*[]domains.GeneralAttendance, error) {
	var generalAttendances []models.GeneralAttendance
	if err := r.db.Model(&models.GeneralAttendance{}).Where(
		"school_id = ?", schoolId,
	).Order("date desc").Find(&generalAttendances).Error; err != nil {
		return nil, err
	}

	result := make([]domains.GeneralAttendance, len(generalAttendances))
	for i, v := range generalAttendances {
		result[i] = *domains.FromGeneralAttendanceModel(&v)
	}

	return &result, nil
}

func (r *GeneralAttendance) Update(
	generalAttendanceId uint, data domains.GeneralAttendance,
) (*domains.GeneralAttendance, error) {
	generalAttendance := data.ToModel()
	if err := r.db.Where(
		"id = ?", generalAttendanceId,
	).Updates(&generalAttendance).Error; err != nil {
		return nil, err
	}

	return domains.FromGeneralAttendanceModel(generalAttendance), nil
}

func (r *GeneralAttendance) Delete(generalAttendanceId uint) error {
	return r.db.Where(
		"id = ?", generalAttendanceId,
	).Unscoped().Delete(&models.GeneralAttendance{}).Error
}

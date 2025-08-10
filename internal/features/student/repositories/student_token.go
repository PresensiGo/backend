package repositories

import (
	"api/internal/features/student/domains"
	"api/internal/features/student/models"
	"gorm.io/gorm"
)

type StudentToken struct {
	db *gorm.DB
}

func NewStudentToken(db *gorm.DB) *StudentToken {
	return &StudentToken{
		db: db,
	}
}

func (r *StudentToken) CreateInTx(tx *gorm.DB, data domains.StudentToken) (
	*domains.StudentToken, error,
) {
	studentToken := data.ToModel()
	if err := r.db.Create(&studentToken).Error; err != nil {
		return nil, err
	} else {
		return domains.FromStudentTokenModel(studentToken), nil
	}
}

func (r *StudentToken) GetByStudentId(studentId uint) (*domains.StudentToken, error) {
	var studentToken models.StudentToken
	if err := r.db.Where("student_id = ?", studentId).First(&studentToken).Error; err != nil {
		return nil, err
	} else {
		return domains.FromStudentTokenModel(&studentToken), nil
	}
}

func (r *StudentToken) DeleteByStudentIdInTx(tx *gorm.DB, studentId uint) error {
	return r.db.Where("student_id = ?", studentId).Unscoped().Delete(&models.StudentToken{}).Error
}

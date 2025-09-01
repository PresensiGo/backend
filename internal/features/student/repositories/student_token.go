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
	if err := tx.Create(&studentToken).Error; err != nil {
		return nil, err
	} else {
		return domains.FromStudentTokenModel(studentToken), nil
	}
}

func (r *StudentToken) GetManyByStudentIds(studentIds []uint) (*[]domains.StudentToken, error) {
	var studentTokens []models.StudentToken
	if err := r.db.Where("student_id IN ?", studentIds).Find(&studentTokens).Error; err != nil {
		return nil, err
	} else {
		result := make([]domains.StudentToken, len(studentTokens))
		for i, v := range studentTokens {
			result[i] = *domains.FromStudentTokenModel(&v)
		}
		return &result, nil
	}
}

func (r *StudentToken) GetByStudentId(studentId uint) (*domains.StudentToken, error) {
	var studentToken models.StudentToken
	if err := r.db.Where(
		"student_id = ?", studentId,
	).First(&studentToken).Error; err != nil {
		return nil, err
	} else {
		return domains.FromStudentTokenModel(&studentToken), nil
	}
}

func (r *StudentToken) GetByDeviceId(deviceId string) (*domains.StudentToken, error) {
	var studentToken models.StudentToken
	if err := r.db.Where(
		"device_id = ?", deviceId,
	).First(&studentToken).Error; err != nil {
		return nil, err
	} else {
		return domains.FromStudentTokenModel(&studentToken), nil
	}
}

func (r *StudentToken) GetByRefreshToken(refreshToken string) (*domains.StudentToken, error) {
	var studentToken models.StudentToken
	if err := r.db.Where("refresh_token = ?", refreshToken).First(&studentToken).Error; err != nil {
		return nil, err
	} else {
		return domains.FromStudentTokenModel(&studentToken), nil
	}
}

// func (r *StudentToken) UpdateDeviceId(
// 	studentTokenId uint, deviceId string,
// ) (*domains.StudentToken, error) {
// 	var studentToken models.StudentToken
// 	if err := r.db.Model(&studentToken).Where("id = ?", studentTokenId).UpdateByToken(
// 		"device_id", deviceId,
// 	).Error; err != nil {
// 		return nil, err
// 	} else {
// 		return domains.FromStudentTokenModel(&studentToken), nil
// 	}
// }

func (r *StudentToken) UpdateByRefreshToken(
	refreshToken string, data domains.StudentToken,
) (*domains.StudentToken, error) {
	studentToken := data.ToModel()
	if err := r.db.Where(
		"refresh_token = ?", refreshToken,
	).Updates(&studentToken).Error; err != nil {
		return nil, err
	} else {
		return domains.FromStudentTokenModel(studentToken), nil
	}
}

func (r *StudentToken) Delete(studentTokenId uint) error {
	return r.db.Where("id = ?", studentTokenId).Unscoped().Delete(&models.StudentToken{}).Error
}

func (r *StudentToken) DeleteByStudentIdInTx(tx *gorm.DB, studentId uint) error {
	return r.db.Where("student_id = ?", studentId).Unscoped().Delete(&models.StudentToken{}).Error
}

package repositories

import (
	"api/internal/dto"
	"api/internal/models"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db}
}

func (r *User) Create(tx *gorm.DB, data dto.User) (uint, error) {
	user := models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
	if err := tx.Create(&user).Error; err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (r *User) GetByID(id uint) (*dto.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", id).
		First(&user).
		Error; err != nil {
		return nil, err
	}

	return &dto.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
		SchoolId: user.SchoolId,
	}, nil
}

func (r *User) GetByEmail(email string) (*dto.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).
		First(&user).
		Error; err != nil {
		return nil, err
	}

	return &dto.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
		SchoolId: user.SchoolId,
	}, nil
}

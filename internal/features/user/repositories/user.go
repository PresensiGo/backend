package repositories

import (
	"api/internal/features/user/domains"
	"api/internal/features/user/models"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db}
}

// @deprecated
func (r *User) Create(tx *gorm.DB, data domains.User) (*domains.User, error) {
	user := data.ToModel()
	if err := tx.Create(&user).Error; err != nil {
		return nil, err
	} else {
		return domains.FromUserModel(user), nil
	}
}

func (r *User) CreateInTx(tx *gorm.DB, data domains.User) (*domains.User, error) {
	user := data.ToModel()
	if err := tx.Create(&user).Error; err != nil {
		return nil, err
	} else {
		return domains.FromUserModel(user), nil
	}
}

func (r *User) CreateBatch(data []domains.User) (*[]domains.User, error) {
	users := make([]models.User, len(data))
	for i, v := range data {
		users[i] = *v.ToModel()
	}

	if err := r.db.Create(&users).Error; err != nil {
		return nil, err
	} else {
		result := make([]domains.User, len(users))
		for i, v := range users {
			result[i] = *domains.FromUserModel(&v)
		}

		return &result, nil
	}
}

func (r *User) GetAll(schoolId uint) (*[]domains.User, error) {
	var users []models.User
	if err := r.db.Where("school_id = ?", schoolId).
		Find(&users).
		Error; err != nil {
		return nil, err
	} else {
		result := make([]domains.User, len(users))
		for i, v := range users {
			result[i] = *domains.FromUserModel(&v)
		}
		return &result, nil
	}
}

func (r *User) GetByID(id uint) (*domains.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", id).
		First(&user).
		Error; err != nil {
		return nil, err
	} else {
		return domains.FromUserModel(&user), nil
	}
}

func (r *User) GetByEmail(email string) (*domains.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).
		First(&user).
		Error; err != nil {
		return nil, err
	} else {
		return domains.FromUserModel(&user), nil
	}
}

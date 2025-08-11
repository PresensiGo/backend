package services

import (
	"errors"

	domains2 "api/internal/features/school/domains"
	schoolRepo "api/internal/features/school/repositories"
	"api/internal/features/user/domains"
	"api/internal/features/user/repositories"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Admin struct {
	db         *gorm.DB
	userRepo   *repositories.User
	schoolRepo *schoolRepo.School
}

func NewAdmin(db *gorm.DB, userRepo *repositories.User, schoolRepo *schoolRepo.School) *Admin {
	return &Admin{
		db:         db,
		userRepo:   userRepo,
		schoolRepo: schoolRepo,
	}
}

func (s *Admin) Inject(schoolName, schoolCode, name, email, password string) error {
	isAdminExists := true
	_, err := s.userRepo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			isAdminExists = false
		} else {
			return err
		}
	}

	if !isAdminExists {
		return s.db.Transaction(
			func(tx *gorm.DB) error {
				// cek apakah sekolah sudah ada
				isSchoolExists := true
				oldSchool, err := s.schoolRepo.GetByCodeInTx(tx, schoolCode)
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						isSchoolExists = false
					} else {
						return err
					}
				}

				var schoolId uint
				if !isSchoolExists {
					school := domains2.School{
						Name: schoolName,
						Code: schoolCode,
					}
					if result, err := s.schoolRepo.CreateInTx(tx, school); err != nil {
						return err
					} else {
						schoolId = result.Id
					}
				} else {
					schoolId = oldSchool.Id
				}

				// belum ada akun admin, maka buat akun baru admin baru
				hashedPassword, err := bcrypt.GenerateFromPassword(
					[]byte(password), bcrypt.DefaultCost,
				)
				if err != nil {
					return err
				}

				user := domains.User{
					Name:     name,
					Email:    email,
					Password: string(hashedPassword),
					Role:     domains.AdminUserRole,
					SchoolId: schoolId,
				}
				_, err = s.userRepo.CreateInTx(tx, user)
				if err != nil {
					return err
				}

				return nil
			},
		)
	}

	return nil
}

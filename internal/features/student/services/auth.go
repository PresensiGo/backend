package services

import (
	"errors"
	"fmt"
	"time"

	schoolRepo "api/internal/features/school/repositories"
	"api/internal/features/student/domains"
	"api/internal/features/student/dto/requests"
	"api/internal/features/student/dto/responses"
	"api/internal/features/student/repositories"
	"api/pkg/authentication"
	"api/pkg/authentication/claims"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StudentAuth struct {
	db               *gorm.DB
	schoolRepo       *schoolRepo.School
	studentRepo      *repositories.Student
	studentTokenRepo *repositories.StudentToken
}

func NewStudentAuth(
	db *gorm.DB,
	schoolRepo *schoolRepo.School,
	studentRepo *repositories.Student,
	studentTokenRepo *repositories.StudentToken,
) *StudentAuth {
	return &StudentAuth{
		db:               db,
		schoolRepo:       schoolRepo,
		studentRepo:      studentRepo,
		studentTokenRepo: studentTokenRepo,
	}
}

func (s *StudentAuth) Login(req requests.LoginStudent) (*responses.LoginStudent, error) {
	school, err := s.schoolRepo.GetByCode(req.SchoolCode)
	if err != nil {
		return nil, err
	}

	student, err := s.studentRepo.GetBySchoolIdNIS(school.Id, req.NIS)
	if err != nil {
		return nil, err
	}

	// mendapatkan token student
	isTokenRegistered := true
	oldStudentToken, err := s.studentTokenRepo.GetByStudentId(student.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			isTokenRegistered = false
		} else {
			return nil, err
		}
	}

	if isTokenRegistered {
		// validasi id perangkat jika token sudah terdaftar
		if req.DeviceId != oldStudentToken.DeviceId {
			return nil, errors.New("id perangkat tidak sesuai dengan yang sudah terdaftar")
		}
	}

	var response responses.LoginStudent
	if err := s.db.Transaction(
		func(tx *gorm.DB) error {
			refreshToken := uuid.NewString()

			// buat token baru
			studentToken := domains.StudentToken{
				StudentId:    student.Id,
				DeviceId:     req.DeviceId,
				RefreshToken: refreshToken,
				TTL:          time.Now().Add(time.Hour * 24 * 30),
			}

			if accessToken, err := s.generateAccessToken(
				student.Id, student.Name, student.NIS, student.SchoolId,
			); err != nil {
				return err
			} else {
				response.AccessToken = accessToken
				response.RefreshToken = refreshToken
			}

			// hapus token lama berdasarkan student id
			if err := s.studentTokenRepo.DeleteByStudentIdInTx(tx, student.Id); err != nil {
				return err
			}

			if _, err := s.studentTokenRepo.CreateInTx(tx, studentToken); err != nil {
				return err
			}

			return nil
		},
	); err != nil {
		return nil, err
	} else {
		return &response, err
	}
}

func (s *StudentAuth) RefreshToken(req requests.RefreshTokenStudent) (
	*responses.RefreshTokenStudent, error,
) {
	oldStudentToken, err := s.studentTokenRepo.GetByRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, err
	}

	if time.Now().After(oldStudentToken.TTL) {
		return nil, fmt.Errorf("refresh token expired")
	}

	currentStudent, err := s.studentRepo.Get(oldStudentToken.StudentId)
	if err != nil {
		return nil, err
	}

	// generate user token
	accessToken, err := s.generateAccessToken(
		currentStudent.Id, currentStudent.Name, currentStudent.NIS, currentStudent.SchoolId,
	)
	if err != nil {
		return nil, err
	}
	refreshToken := uuid.New().String()

	// store new token into database
	if _, err := s.studentTokenRepo.UpdateByRefreshToken(
		req.RefreshToken, domains.StudentToken{
			RefreshToken: refreshToken,
		},
	); err != nil {
		return nil, err
	}

	return &responses.RefreshTokenStudent{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *StudentAuth) generateAccessToken(
	id uint, name string, nis string, schoolId uint,
) (string, error) {
	return authentication.GenerateStudentJWT(
		claims.Student{
			Id:       id,
			Name:     name,
			NIS:      nis,
			SchoolId: schoolId,
		},
	)
}

package services

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	schoolRepo "api/internal/features/school/repositories"
	"api/internal/features/student/domains"
	"api/internal/features/student/dto/requests"
	"api/internal/features/student/dto/responses"
	"api/internal/features/student/repositories"
	"api/pkg/authentication"
	"api/pkg/authentication/claims"
	"api/pkg/http/failure"

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

func (s *StudentAuth) Login(req requests.LoginStudent) (*responses.LoginStudent, *failure.App) {
	school, err := s.schoolRepo.GetByCode(req.SchoolCode)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, failure.NewApp(
				http.StatusNotFound, "Sekolah dengan id tersebut tidak ditemukan!", nil,
			)
		}
		return nil, failure.NewInternal(err)
	}

	student, err := s.studentRepo.GetBySchoolIdNIS(school.Id, req.NIS)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, failure.NewApp(
				http.StatusNotFound, "Siswa dengan NIS tersebut tidak ditemukan!", nil,
			)
		}
		return nil, failure.NewInternal(err)
	}

	// cek berdasarkan id siswa
	if result, err := s.studentTokenRepo.GetByStudentId(student.Id); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, failure.NewInternal(err)
		}
	} else {
		if result.DeviceId != req.DeviceId {
			return nil, failure.NewApp(
				http.StatusConflict,
				"Perangkat dengan NIS tidak valid! Silahkan hubungi admin sekolah",
				nil,
			)
		}
	}

	// cek berdasarkan id perangkat
	if result, err := s.studentTokenRepo.GetByDeviceId(req.DeviceId); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, failure.NewInternal(err)
		}
	} else {
		if result.StudentId != student.Id {
			return nil, failure.NewApp(
				http.StatusConflict,
				"Perangkat dengan NIS tidak valid! Silahkan hubungi admin sekolah",
				nil,
			)
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
		return nil, failure.NewInternal(err)
	} else {
		return &response, nil
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

func (s *StudentAuth) Eject(studentTokenId uint) (*responses.EjectStudentToken, error) {
	if err := s.studentTokenRepo.Delete(studentTokenId); err != nil {
		return nil, err
	} else {
		return &responses.EjectStudentToken{
			Message: "ok",
		}, nil
	}
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

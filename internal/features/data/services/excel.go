package services

import (
	"fmt"
	"io"
	"net/http"

	"api/internal/features/batch/domains"
	batchRepo "api/internal/features/batch/repositories"
	classroomDomain "api/internal/features/classroom/domains"
	classroomRepo "api/internal/features/classroom/repositories"
	"api/internal/features/data/dto/responses"
	majorDomain "api/internal/features/major/domains"
	majorRepo "api/internal/features/major/repositories"
	studentDomain "api/internal/features/student/domains"
	"api/internal/features/student/repositories"
	"api/pkg/http/failure"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type Excel struct {
	batchRepo     *batchRepo.Batch
	majorRepo     *majorRepo.Major
	classroomRepo *classroomRepo.Classroom
	studentRepo   *repositories.Student

	db *gorm.DB
}

func NewExcel(
	batchRepo *batchRepo.Batch,
	majorRepo *majorRepo.Major,
	classroomRepo *classroomRepo.Classroom,
	studentRepo *repositories.Student,
	db *gorm.DB,
) *Excel {
	return &Excel{
		batchRepo,
		majorRepo,
		classroomRepo,
		studentRepo,
		db,
	}
}

func (s *Excel) ImportData(schoolId uint, reader io.Reader) (
	*responses.ImportData, *failure.App,
) {
	file, err := excelize.OpenReader(reader)
	if err != nil {
		return nil, failure.NewInternal(err)
	}
	defer file.Close()

	sheets := file.GetSheetList()

	customErrorCode := 0
	customErrorMessage := ""
	if err := s.db.Transaction(
		func(tx *gorm.DB) error {
			for _, sheet := range sheets {
				rows, err := file.GetRows(sheet)
				if err != nil {
					return err
				}

				if len(rows) < 5 {
					customErrorCode = http.StatusBadRequest
					customErrorMessage = fmt.Sprintf("Sheet %s memiliki jumlah baris < 5!", sheet)
					return fmt.Errorf("rows length less than 5")
				}

				var batch domains.Batch
				if result, err := s.batchRepo.FirstOrCreateInTx(
					tx, domains.Batch{
						Name:     rows[0][1],
						SchoolId: schoolId,
					},
				); err != nil {
					return err
				} else {
					batch = *result
				}

				var major majorDomain.Major
				if result, err := s.majorRepo.FirstOrCreateInTx(
					tx, majorDomain.Major{
						Name:    rows[1][1],
						BatchId: batch.Id,
					},
				); err != nil {
					return err
				} else {
					major = *result
				}

				var classroom classroomDomain.Classroom
				if result, err := s.classroomRepo.FirstOrCreateInTx(
					tx, classroomDomain.Classroom{
						Name:    rows[2][1],
						MajorId: major.Id,
					},
				); err != nil {
					return err
				} else {
					classroom = *result
				}

				for _, row := range rows[5:] {
					if _, err := s.studentRepo.GetOrCreateInTx(
						tx, studentDomain.Student{
							NIS:         row[0],
							Name:        row[1],
							SchoolId:    schoolId,
							ClassroomId: classroom.Id,
							Gender:      row[2],
						},
					); err != nil {
						return err
					}
				}
			}

			return nil
		},
	); err != nil {
		if customErrorCode != 0 {
			return nil, failure.NewApp(
				customErrorCode,
				customErrorMessage,
				err,
			)
		}
		return nil, failure.NewInternal(err)
	} else {
		return &responses.ImportData{
			Message: "ok",
		}, nil
	}
}

package services

import (
	"api/internal/dto"
	models2 "api/internal/models"
	"api/internal/repositories"
	"fmt"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"io"
)

type Excel struct {
	batchRepo     *repositories.Batch
	majorRepo     *repositories.Major
	classroomRepo *repositories.Classroom
	studentRepo   *repositories.Student

	db *gorm.DB
}

func NewExcel(
	batchRepo *repositories.Batch,
	majorRepo *repositories.Major,
	classroomRepo *repositories.Classroom,
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

func (s *Excel) ImportData(schoolId uint, reader io.Reader) error {
	file, err := excelize.OpenReader(reader)
	if err != nil {
		return err
	}
	defer file.Close()

	sheets := file.GetSheetList()

	return s.db.Transaction(func(tx *gorm.DB) error {
		mapBatch := make(map[string]uint)
		mapMajor := make(map[string]uint)
		mapClassroom := make(map[string]uint)

		for _, sheet := range sheets {
			// batch
			batchName, err := file.GetCellValue(sheet, "B1")
			if err != nil {
				return err
			}

			batchKey := batchName
			var batchId uint

			if value, ok := mapBatch[batchKey]; ok {
				batchId = value
			} else {
				newBatchId, err := s.batchRepo.CreateInTx(tx, dto.Batch{
					Name:     batchName,
					SchoolId: schoolId,
				})
				if err != nil {
					return err
				}

				mapBatch[batchKey] = *newBatchId
				batchId = *newBatchId
			}

			// major
			majorName, err := file.GetCellValue(sheet, "B2")
			if err != nil {
				return err
			}

			majorKey := fmt.Sprintf("%s-%s", batchName, majorName)
			var majorId uint

			if value, ok := mapMajor[majorKey]; ok {
				majorId = value
			} else {
				newMajorId, err := s.majorRepo.CreateInTx(tx, dto.Major{
					Name:    majorName,
					BatchId: batchId,
				})
				if err != nil {
					return err
				}
				mapMajor[majorKey] = *newMajorId
				majorId = *newMajorId
			}

			// classroom
			classroomName, err := file.GetCellValue(sheet, "B3")
			if err != nil {
				return err
			}

			classroomKey := fmt.Sprintf("%s-%s-%s", batchName, majorName, classroomName)
			var classroomId uint

			if value, ok := mapClassroom[classroomKey]; ok {
				classroomId = value
			} else {
				newClassroomId, err := s.classroomRepo.CreateInTx(tx, dto.Classroom{
					Name:    classroomName,
					MajorID: majorId,
				})
				if err != nil {
					return err
				}
				mapClassroom[classroomKey] = *newClassroomId
				classroomId = *newClassroomId
			}

			// students
			rows, err := file.GetRows(sheet)
			if err != nil {
				return err
			}

			if len(rows) < 5 {
				return fmt.Errorf("rows length less than 5")
			}

			var students []dto.Student
			for _, row := range rows[5:] {
				studentNIS := row[0]
				studentName := row[1]

				students = append(students, dto.Student{
					NIS:         studentNIS,
					Name:        studentName,
					ClassroomID: classroomId,
				})
			}
			if err := s.studentRepo.CreateBatchInTx(tx, students); err != nil {
				return err
			}
		}

		return nil
	})
}

func (s *Excel) Import(reader io.Reader) (any, error) {
	file, err := excelize.OpenReader(reader)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sheets := file.GetSheetList()
	fmt.Println("sheets", sheets)

	if err := s.db.Transaction(func(tx *gorm.DB) error {
		for _, sheet := range sheets {
			// mendapatkan data angkatan, jurusan, kelas
			batchName, err := file.GetCellValue(sheet, "B1")
			if err != nil {
				return err
			}

			batch := models2.Batch{
				Name: batchName,
			}
			if err := tx.Create(&batch).Error; err != nil {
				return err
			}

			majorName, err := file.GetCellValue(sheet, "B2")
			if err != nil {
				return err
			}

			major := models2.Major{
				Name:    majorName,
				BatchId: batch.ID,
			}
			if err := tx.Create(&major).Error; err != nil {
				return err
			}

			className, err := file.GetCellValue(sheet, "B3")
			if err != nil {
				return err
			}

			class := models2.Classroom{
				Name:    className,
				MajorId: major.ID,
			}
			if err := tx.Create(&class).Error; err != nil {
				return err
			}

			// mendapatkan data siswa
			rows, err := file.GetRows(sheet)
			if err != nil {
				return err
			}

			if len(rows) < 5 {
				return fmt.Errorf("rows length less than 5")
			}

			var students []models2.Student
			for _, row := range rows[5:] {
				studentNIS := row[0]
				studentName := row[1]

				students = append(students, models2.Student{
					NIS:         studentNIS,
					Name:        studentName,
					ClassroomId: class.ID,
				})
			}

			if err := tx.CreateInBatches(students, len(students)).Error; err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return nil, nil
}

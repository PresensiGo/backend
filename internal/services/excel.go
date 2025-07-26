package services

import (
	models2 "api/internal/models"
	"fmt"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"io"
)

type Excel struct {
	db *gorm.DB
}

func NewExcel(db *gorm.DB) *Excel {
	return &Excel{db}
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

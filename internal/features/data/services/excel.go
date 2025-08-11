package services

import (
	"errors"
	"fmt"
	"io"

	"api/internal/features/batch/domains"
	"api/internal/features/batch/models"
	repositories2 "api/internal/features/batch/repositories"
	domains4 "api/internal/features/classroom/domains"
	models4 "api/internal/features/classroom/models"
	repositories4 "api/internal/features/classroom/repositories"
	"api/internal/features/data/dto/responses"
	domains2 "api/internal/features/major/domains"
	models3 "api/internal/features/major/models"
	repositories3 "api/internal/features/major/repositories"
	domains3 "api/internal/features/student/domains"
	models2 "api/internal/features/student/models"
	"api/internal/features/student/repositories"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type Excel struct {
	batchRepo     *repositories2.Batch
	majorRepo     *repositories3.Major
	classroomRepo *repositories4.Classroom
	studentRepo   *repositories.Student

	db *gorm.DB
}

func NewExcel(
	batchRepo *repositories2.Batch,
	majorRepo *repositories3.Major,
	classroomRepo *repositories4.Classroom,
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

func (s *Excel) ImportDataV3(schoolId uint, reader io.Reader) (*responses.ImportData, error) {
	file, err := excelize.OpenReader(reader)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sheets := file.GetSheetList()

	if err := s.db.Transaction(
		func(tx *gorm.DB) error {
			for _, sheet := range sheets {
				rows, err := file.GetRows(sheet)
				if err != nil {
					return err
				}

				if len(rows) < 5 {
					return fmt.Errorf("rows length less than 5")
				}

				var batch domains.Batch
				var major domains2.Major
				var classroom domains4.Classroom
				var students []domains3.Student

				for i, row := range rows {
					if i == 0 {
						// batch
						batchName := row[1]

						if result, err := s.batchRepo.GetBySchoolIdNameInTx(
							tx, schoolId, batchName,
						); err != nil {
							if errors.Is(err, gorm.ErrRecordNotFound) {
								// jika angkatan belum ada, maka buat baru
								if result, err := s.batchRepo.CreateInTx2(
									tx, domains.Batch{
										Name:     batchName,
										SchoolId: schoolId,
									},
								); err != nil {
									return err
								} else {
									batch = *result
								}
							} else {
								return err
							}
						} else {
							batch = *result
						}
					} else if i == 1 {
						// major
						majorName := row[1]

						if result, err := s.majorRepo.GetByBatchIdNameInTx(
							tx, batch.Id, majorName,
						); err != nil {
							if errors.Is(err, gorm.ErrRecordNotFound) {
								// jika jurusan belum ada, maka buat baru
								if result, err := s.majorRepo.CreateInTx2(
									tx, domains2.Major{
										Name:    majorName,
										BatchId: batch.Id,
									},
								); err != nil {
									return err
								} else {
									major = *result
								}
							} else {
								return err
							}
						} else {
							major = *result
						}
					} else if i == 2 {
						// classroom
						// untuk classroom tidak perlu ada pengecekan duplikasi, karena
						// setiap sheet artinya setiap classroom
						classroomName := row[1]

						if result, err := s.classroomRepo.CreateInTx2(
							tx, domains4.Classroom{
								Name:    classroomName,
								MajorId: major.Id,
							},
						); err != nil {
							return err
						} else {
							classroom = *result
						}
					} else if i >= 5 {
						studentNIS := row[0]
						studentName := row[1]

						students = append(
							students, domains3.Student{
								NIS:         studentNIS,
								Name:        studentName,
								SchoolId:    schoolId,
								ClassroomId: classroom.Id,
							},
						)
					}
				}

				// simpan data siswa
				if _, err := s.studentRepo.CreateBatchInTx2(tx, students); err != nil {
					return err
				}
			}

			return nil
		},
	); err != nil {
		return nil, err
	} else {
		return &responses.ImportData{
			Message: "ok",
		}, nil
	}
}

// @deprecated
func (s *Excel) ImportData(schoolId uint, reader io.Reader) error {
	file, err := excelize.OpenReader(reader)
	if err != nil {
		return err
	}
	defer file.Close()

	sheets := file.GetSheetList()

	return s.db.Transaction(
		func(tx *gorm.DB) error {
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
					newBatchId, err := s.batchRepo.CreateInTx(
						tx, domains.Batch{
							Name:     batchName,
							SchoolId: schoolId,
						},
					)
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
					newMajorId, err := s.majorRepo.CreateInTx(
						tx, domains2.Major{
							Name:    majorName,
							BatchId: batchId,
						},
					)
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
					newClassroomId, err := s.classroomRepo.CreateInTx(
						tx, domains4.Classroom{
							Name:    classroomName,
							MajorId: majorId,
						},
					)
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

				var students []domains3.Student
				for _, row := range rows[5:] {
					studentNIS := row[0]
					studentName := row[1]

					students = append(
						students, domains3.Student{
							NIS:         studentNIS,
							Name:        studentName,
							ClassroomId: classroomId,
						},
					)
				}
				if err := s.studentRepo.CreateBatchInTx(tx, students); err != nil {
					return err
				}
			}

			return nil
		},
	)
}

// @deprecated
func (s *Excel) Import(reader io.Reader) (any, error) {
	file, err := excelize.OpenReader(reader)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sheets := file.GetSheetList()
	fmt.Println("sheets", sheets)

	if err := s.db.Transaction(
		func(tx *gorm.DB) error {
			for _, sheet := range sheets {
				// mendapatkan data angkatan, jurusan, kelas
				batchName, err := file.GetCellValue(sheet, "B1")
				if err != nil {
					return err
				}

				batch := models.Batch{
					Name: batchName,
				}
				if err := tx.Create(&batch).Error; err != nil {
					return err
				}

				majorName, err := file.GetCellValue(sheet, "B2")
				if err != nil {
					return err
				}

				major := models3.Major{
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

				class := models4.Classroom{
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

					students = append(
						students, models2.Student{
							NIS:         studentNIS,
							Name:        studentName,
							ClassroomId: class.ID,
						},
					)
				}

				if err := tx.CreateInBatches(students, len(students)).Error; err != nil {
					return err
				}
			}

			return nil
		},
	); err != nil {
		return nil, err
	}

	return nil, nil
}

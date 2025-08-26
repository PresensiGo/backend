package services

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"api/internal/features/attendance/domains"
	"api/internal/features/attendance/dto"
	"api/internal/features/attendance/dto/requests"
	"api/internal/features/attendance/dto/responses"
	"api/internal/features/attendance/repositories"
	batchRepo "api/internal/features/batch/repositories"
	classroomRepo "api/internal/features/classroom/repositories"
	majorRepo "api/internal/features/major/repositories"
	studentDomain "api/internal/features/student/domains"
	studentRepo "api/internal/features/student/repositories"
	userDomain "api/internal/features/user/domains"
	userRepo "api/internal/features/user/repositories"
	"api/pkg/authentication"
	"api/pkg/constants"
	"api/pkg/http/failure"
	"api/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type GeneralAttendance struct {
	db                          *gorm.DB
	batchRepo                   *batchRepo.Batch
	majorRepo                   *majorRepo.Major
	classroomRepo               *classroomRepo.Classroom
	studentRepo                 *studentRepo.Student
	generalAttendanceRepo       *repositories.GeneralAttendance
	generalAttendanceRecordRepo *repositories.GeneralAttendanceRecord
	userRepo                    *userRepo.User
}

func NewGeneralAttendance(
	db *gorm.DB,
	batchRepo *batchRepo.Batch,
	majorRepo *majorRepo.Major,
	classroomRepo *classroomRepo.Classroom,
	studentRepo *studentRepo.Student,
	generalAttendanceRepo *repositories.GeneralAttendance,
	generalAttendanceRecordRepo *repositories.GeneralAttendanceRecord,
	userRepo *userRepo.User,
) *GeneralAttendance {
	return &GeneralAttendance{
		db:                          db,
		batchRepo:                   batchRepo,
		majorRepo:                   majorRepo,
		classroomRepo:               classroomRepo,
		studentRepo:                 studentRepo,
		generalAttendanceRepo:       generalAttendanceRepo,
		generalAttendanceRecordRepo: generalAttendanceRecordRepo,
		userRepo:                    userRepo,
	}
}

func (s *GeneralAttendance) CreateGeneralAttendance(
	c *gin.Context, schoolId uint, req requests.CreateGeneralAttendance,
) (*responses.CreateGeneralAttendance, *failure.App) {
	user := authentication.GetAuthenticatedUser(c)
	if user.ID == 0 {
		return nil, failure.NewApp(http.StatusForbidden, "Anda tidak memiliki akses!", nil)
	}

	parsedDateTime, err := utils.GetParsedDateTime(req.DateTime)
	if err != nil {
		return nil, failure.NewApp(
			http.StatusBadRequest,
			"Tanggal dan waktu tidak valid!",
			err,
		)
	}

	generalAttendance := domains.GeneralAttendance{
		DateTime:  *parsedDateTime,
		Note:      req.Note,
		SchoolId:  schoolId,
		Code:      uuid.NewString(),
		CreatorId: user.ID,
	}

	if result, err := s.generalAttendanceRepo.Create(generalAttendance); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.CreateGeneralAttendance{
			GeneralAttendance: *result,
		}, nil
	}
}

func (s *GeneralAttendance) CreateGeneralAttendanceRecord(
	generalAttendanceId uint, req requests.CreateGeneralAttendanceRecord,
) (*responses.CreateGeneralAttendanceRecord, *failure.App) {
	parsedDateTime, err := utils.GetParsedDateTime(req.DateTime)
	if err != nil {
		return nil, failure.NewApp(http.StatusBadRequest, "Tanggal dan waktu tidak valid!", err)
	}

	if result, err := s.generalAttendanceRecordRepo.FirstOrCreate(
		domains.GeneralAttendanceRecord{
			DateTime:            *parsedDateTime,
			GeneralAttendanceId: generalAttendanceId,
			StudentId:           req.StudentId,
			Status:              req.Status,
		},
	); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.CreateGeneralAttendanceRecord{
			GeneralAttendanceRecord: *result,
		}, nil
	}
}

func (s *GeneralAttendance) CreateGeneralAttendanceRecordStudent(
	studentId uint, req requests.CreateGeneralAttendanceRecordStudent,
) (*responses.CreateGeneralAttendanceRecordStudent, *failure.App) {
	generalAttendance, err := s.generalAttendanceRepo.GetByCode(req.Code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, failure.NewApp(http.StatusNotFound, "Kode presensi tidak ditemukan!", nil)
		}
		return nil, failure.NewInternal(err)
	}

	if _, err := s.generalAttendanceRecordRepo.FirstOrCreate(
		domains.GeneralAttendanceRecord{
			GeneralAttendanceId: generalAttendance.Id,
			StudentId:           studentId,
			DateTime:            time.Now(),
			Status:              constants.AttendanceStatusPresent,
		},
	); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.CreateGeneralAttendanceRecordStudent{
			Message: "ok",
		}, nil
	}
}

func (s *GeneralAttendance) GetAllGeneralAttendances(schoolId uint) (
	*responses.GetAllGeneralAttendances, *failure.App,
) {
	if generalAttendances, err := s.generalAttendanceRepo.GetAllBySchoolId(schoolId); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		creatorIds := make([]uint, len(*generalAttendances))
		for i, v := range *generalAttendances {
			creatorIds[i] = v.CreatorId
		}

		mapCreators := make(map[uint]*userDomain.User)

		if creators, err := s.userRepo.GetMany(creatorIds); err != nil {
			return nil, failure.NewInternal(err)
		} else {
			for _, v := range *creators {
				mapCreators[v.Id] = &v
			}
		}

		result := make([]dto.GetAllGeneralAttendancesItem, len(*generalAttendances))
		for i, v := range *generalAttendances {
			var creator userDomain.User
			if v, ok := mapCreators[v.CreatorId]; ok {
				creator = *v
			} else {
				creator = userDomain.User{}
			}

			result[i] = dto.GetAllGeneralAttendancesItem{
				GeneralAttendance: v,
				Creator:           creator,
			}
		}

		return &responses.GetAllGeneralAttendances{
			Items: result,
		}, nil
	}
}

func (s *GeneralAttendance) GetAllGeneralAttendancesStudent(c *gin.Context) (
	*responses.GetAllGeneralAttendancesStudent, *failure.App,
) {
	student := authentication.GetAuthenticatedStudent(c)
	if student.Id == 0 || student.SchoolId == 0 {
		return nil, failure.NewApp(http.StatusForbidden, "Anda tidak memiliki akses!", nil)
	}

	generalAttendances, err := s.generalAttendanceRepo.GetAllTodayBySchoolId(student.SchoolId)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	generalAttendanceIds := make([]uint, len(*generalAttendances))
	for i, v := range *generalAttendances {
		generalAttendanceIds[i] = v.Id
	}

	mapGeneralAttendanceRecords := make(map[uint]*domains.GeneralAttendanceRecord)

	if generalAttendanceRecords, err := s.generalAttendanceRecordRepo.GetManyByAttendanceIdsStudentId(
		generalAttendanceIds, student.Id,
	); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		for _, record := range *generalAttendanceRecords {
			mapGeneralAttendanceRecords[record.GeneralAttendanceId] = &record
		}
	}

	result := make([]dto.GetAllGeneralAttendancesStudentItem, len(*generalAttendances))
	for i, v := range *generalAttendances {
		record := domains.GeneralAttendanceRecord{}
		if r, ok := mapGeneralAttendanceRecords[v.Id]; ok {
			record = *r
		}

		result[i] = dto.GetAllGeneralAttendancesStudentItem{
			GeneralAttendance:       v,
			GeneralAttendanceRecord: record,
		}
	}

	return &responses.GetAllGeneralAttendancesStudent{
		Items: result,
	}, nil
}

func (s *GeneralAttendance) GetAllGeneralAttendanceRecords(generalAttendanceId uint) (
	*responses.GetAllGeneralAttendanceRecords, *failure.App,
) {
	records, err := s.generalAttendanceRecordRepo.GetAll(generalAttendanceId)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	studentIds := make([]uint, len(*records))
	for i, v := range *records {
		studentIds[i] = v.StudentId
	}

	students, err := s.studentRepo.GetManyById(studentIds)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	mapStudents := make(map[uint]*studentDomain.Student)
	for _, student := range *students {
		mapStudents[student.Id] = &student
	}

	result := make([]dto.GetAllGeneralAttendanceRecordsItem, len(*records))
	for i, v := range *records {
		result[i] = dto.GetAllGeneralAttendanceRecordsItem{
			Student: *mapStudents[v.StudentId],
			Record:  v,
		}
	}

	return &responses.GetAllGeneralAttendanceRecords{
		Items: result,
	}, nil
}

func (s *GeneralAttendance) GetAllGeneralAttendanceRecordsByClassroomId(
	generalAttendanceId uint, classroomId uint,
) (
	*responses.GetAllGeneralAttendanceRecordsByClassroomId, *failure.App,
) {
	students, err := s.studentRepo.GetAllByClassroomId(classroomId)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	studentIds := make([]uint, len(students))
	for i, v := range students {
		studentIds[i] = v.Id
	}

	records, err := s.generalAttendanceRecordRepo.GetManyByAttendanceIdStudentIds(
		generalAttendanceId, studentIds,
	)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	mapRecords := make(map[uint]*domains.GeneralAttendanceRecord)
	for _, record := range *records {
		mapRecords[record.StudentId] = &record
	}

	result := make([]dto.GetAllGeneralAttendanceRecordsByClassroomIdItem, len(students))
	for i, student := range students {
		record := domains.GeneralAttendanceRecord{}
		if r, ok := mapRecords[student.Id]; ok {
			record = *r
		}

		result[i] = dto.GetAllGeneralAttendanceRecordsByClassroomIdItem{
			Student: student,
			Record:  record,
		}
	}

	return &responses.GetAllGeneralAttendanceRecordsByClassroomId{
		Items: result,
	}, nil
}

func (s *GeneralAttendance) GetGeneralAttendance(generalAttendanceId uint) (
	*responses.GetGeneralAttendance, *failure.App,
) {
	if generalAttendance, err := s.generalAttendanceRepo.Get(generalAttendanceId); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		response := responses.GetGeneralAttendance{
			GeneralAttendance: *generalAttendance,
			Creator:           userDomain.User{},
		}

		// mendapatkan creator
		if generalAttendance.CreatorId != 0 {
			if creator, err := s.userRepo.GetByID(generalAttendance.CreatorId); err != nil {
				return nil, failure.NewInternal(err)
			} else {
				response.Creator = *creator
			}
		}

		return &response, nil
	}
}

func (s *GeneralAttendance) Update(
	generalAttendanceId uint, req requests.UpdateGeneralAttendance,
) (*responses.UpdateGeneralAttendance, error) {
	parsedDateTime, err := utils.GetParsedDateTime(req.DateTime)
	if err != nil {
		return nil, err
	}

	generalAttendance := domains.GeneralAttendance{
		DateTime: *parsedDateTime,
		Note:     req.Note,
	}

	result, err := s.generalAttendanceRepo.Update(generalAttendanceId, generalAttendance)
	if err != nil {
		return nil, err
	}

	return &responses.UpdateGeneralAttendance{
		GeneralAttendance: *result,
	}, nil
}

func (s *GeneralAttendance) DeleteGeneralAttendance(generalAttendanceId uint) (
	*responses.DeleteGeneralAttendance, *failure.App,
) {
	if err := s.generalAttendanceRepo.Delete(generalAttendanceId); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.DeleteGeneralAttendance{
			Message: "ok",
		}, nil
	}
}

func (s *GeneralAttendance) DeleteGeneralAttendanceRecord(recordId uint) (
	*responses.DeleteGeneralAttendanceRecord, *failure.App,
) {
	if err := s.generalAttendanceRecordRepo.Delete(recordId); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.DeleteGeneralAttendanceRecord{
			Message: "ok",
		}, nil
	}
}

func (s *GeneralAttendance) ExportGeneralAttendance(
	c *gin.Context, req requests.ExportGeneralAttendance,
) (
	*responses.ExportGeneralAttendance, *failure.App,
) {
	parsedStartDate, err := utils.GetParsedDate(req.StartDate)
	if err != nil {
		return nil, failure.NewApp(http.StatusBadRequest, "Tanggal mulai tidak valid!", err)
	}

	parsedEndDate, err := utils.GetParsedDate(req.EndDate)
	if err != nil {
		return nil, failure.NewApp(http.StatusBadRequest, "Tanggal akhir tidak valid!", err)
	}

	endDate := time.Date(
		parsedEndDate.Year(),
		parsedEndDate.Month(),
		parsedEndDate.Day(),
		23, 59, 59, 999_999_999,
		parsedEndDate.Location(),
	)

	auth := authentication.GetAuthenticatedUser(c)
	if auth.SchoolId == 0 {
		return nil, failure.NewApp(http.StatusForbidden, "Anda tidak memiliki akses!", nil)
	}

	f := excelize.NewFile()

	batches, err := s.batchRepo.GetAllBySchoolId(auth.SchoolId)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	batchIds := make([]uint, len(*batches))
	for i, v := range *batches {
		batchIds[i] = v.Id
	}

	majors, err := s.majorRepo.GetManyByBatchIds(batchIds)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	majorIds := make([]uint, len(*majors))
	for i, v := range *majors {
		majorIds[i] = v.Id
	}

	classrooms, err := s.classroomRepo.GetManyByMajorIds(majorIds)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	centerStyle, err := f.NewStyle(
		&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "center",
				Vertical:   "center",
			},
		},
	)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	attendances, err := s.generalAttendanceRepo.GetAllBySchoolIdBetween(
		auth.SchoolId, *parsedStartDate, endDate,
	)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	for _, classroom := range classrooms {
		sheetName := classroom.Name
		sheetIndex, err := f.NewSheet(sheetName)
		if err != nil {
			return nil, failure.NewInternal(err)
		}

		f.SetActiveSheet(sheetIndex)

		// header student
		f.SetCellValue(sheetName, "A1", "No")
		f.MergeCell(sheetName, "A1", "A2")
		f.SetColWidth(sheetName, "A", "A", 4)
		f.SetCellStyle(sheetName, "A1", "A2", centerStyle)

		f.SetCellValue(sheetName, "B1", "NIS")
		f.MergeCell(sheetName, "B1", "B2")
		f.SetColWidth(sheetName, "B", "B", 8)
		f.SetCellStyle(sheetName, "B1", "B2", centerStyle)

		f.SetCellValue(sheetName, "C1", "Nama")
		f.MergeCell(sheetName, "C1", "C2")
		f.SetColWidth(sheetName, "C", "C", 32)
		f.SetCellStyle(sheetName, "C1", "C2", centerStyle)

		f.SetCellValue(sheetName, "D1", "JK")
		f.MergeCell(sheetName, "D1", "D2")
		f.SetColWidth(sheetName, "D", "D", 4)
		f.SetCellStyle(sheetName, "D1", "D2", centerStyle)

		// header attendance
		mapMonths := make(map[string]int)
		for i, attendance := range *attendances {
			columnName, err := utils.ColumnToName(i + 5)
			if err != nil {
				return nil, failure.NewInternal(err)
			}
			monthName := utils.MapMonths[attendance.DateTime.Format("January")]
			if _, ok := mapMonths[monthName]; !ok {
				mapMonths[monthName] = 1
			} else {
				mapMonths[monthName]++
			}

			f.SetCellValue(
				sheetName,
				fmt.Sprintf("%s1", columnName),
				monthName,
			)
			f.SetCellValue(
				sheetName,
				fmt.Sprintf("%s2", columnName),
				attendance.DateTime.Format("2"),
			)
			f.SetCellStyle(
				sheetName,
				fmt.Sprintf("%s2", columnName),
				fmt.Sprintf("%s2", columnName),
				centerStyle,
			)
			f.SetColWidth(sheetName, columnName, columnName, 4)
		}

		// merge month cells
		currentIndex := 5
		for _, v := range mapMonths {
			if v > 1 {
				startColumn, err := utils.ColumnToName(currentIndex)
				if err != nil {
					return nil, failure.NewInternal(err)
				}

				endColumn, err := utils.ColumnToName(currentIndex + v - 1)
				if err != nil {
					return nil, failure.NewInternal(err)
				}

				if err := f.MergeCell(
					sheetName,
					fmt.Sprintf("%s1", startColumn),
					fmt.Sprintf("%s1", endColumn),
				); err != nil {
					return nil, failure.NewInternal(err)
				}

				f.SetCellStyle(
					sheetName,
					fmt.Sprintf("%s1", startColumn),
					fmt.Sprintf("%s1", endColumn),
					centerStyle,
				)
			}
			currentIndex += v
		}

		// rekap
		presentColumn, err := utils.ColumnToName(len(*attendances) + 5)
		if err != nil {
			return nil, failure.NewInternal(err)
		}
		sickColumn, err := utils.ColumnToName(len(*attendances) + 6)
		if err != nil {
			return nil, failure.NewInternal(err)
		}
		permissionColumn, err := utils.ColumnToName(len(*attendances) + 7)
		if err != nil {
			return nil, failure.NewInternal(err)
		}
		alphaColumn, err := utils.ColumnToName(len(*attendances) + 8)
		if err != nil {
			return nil, failure.NewInternal(err)
		}
		lateColumn, err := utils.ColumnToName(len(*attendances) + 9)
		if err != nil {
			return nil, failure.NewInternal(err)
		}

		f.SetCellValue(sheetName, fmt.Sprintf("%s1", presentColumn), "Rekap")
		f.MergeCell(
			sheetName,
			fmt.Sprintf("%s1", presentColumn),
			fmt.Sprintf("%s1", lateColumn),
		)
		f.SetCellStyle(
			sheetName,
			fmt.Sprintf("%s1", presentColumn),
			fmt.Sprintf("%s1", alphaColumn),
			centerStyle,
		)

		f.SetCellValue(sheetName, fmt.Sprintf("%s2", presentColumn), "H")
		f.SetColWidth(sheetName, presentColumn, presentColumn, 6)
		f.SetCellStyle(
			sheetName,
			fmt.Sprintf("%s2", presentColumn),
			fmt.Sprintf("%s2", presentColumn),
			centerStyle,
		)

		f.SetCellValue(sheetName, fmt.Sprintf("%s2", sickColumn), "S")
		f.SetColWidth(sheetName, sickColumn, sickColumn, 6)
		f.SetCellStyle(
			sheetName,
			fmt.Sprintf("%s2", sickColumn),
			fmt.Sprintf("%s2", sickColumn),
			centerStyle,
		)

		f.SetCellValue(sheetName, fmt.Sprintf("%s2", permissionColumn), "I")
		f.SetColWidth(sheetName, permissionColumn, permissionColumn, 6)
		f.SetCellStyle(
			sheetName,
			fmt.Sprintf("%s2", permissionColumn),
			fmt.Sprintf("%s2", permissionColumn),
			centerStyle,
		)

		f.SetCellValue(sheetName, fmt.Sprintf("%s2", alphaColumn), "A")
		f.SetColWidth(sheetName, alphaColumn, alphaColumn, 6)
		f.SetCellStyle(
			sheetName,
			fmt.Sprintf("%s2", alphaColumn),
			fmt.Sprintf("%s2", alphaColumn),
			centerStyle,
		)

		f.SetCellValue(sheetName, fmt.Sprintf("%s2", lateColumn), "T")
		f.SetColWidth(sheetName, lateColumn, lateColumn, 6)
		f.SetCellStyle(
			sheetName,
			fmt.Sprintf("%s2", lateColumn),
			fmt.Sprintf("%s2", lateColumn),
			centerStyle,
		)

		// students
		students, err := s.studentRepo.GetAllByClassroomId(classroom.Id)
		if err != nil {
			return nil, failure.NewInternal(err)
		}

		for studentIndex, student := range students {
			f.SetCellValue(
				sheetName, fmt.Sprintf("A%d", studentIndex+3), studentIndex+1,
			)
			f.SetCellStyle(
				sheetName,
				fmt.Sprintf("A%d", studentIndex+3),
				fmt.Sprintf("A%d", studentIndex+3),
				centerStyle,
			)

			f.SetCellValue(
				sheetName, fmt.Sprintf("B%d", studentIndex+3), student.NIS,
			)
			f.SetCellStyle(
				sheetName,
				fmt.Sprintf("B%d", studentIndex+3),
				fmt.Sprintf("B%d", studentIndex+3),
				centerStyle,
			)

			f.SetCellValue(
				sheetName, fmt.Sprintf("C%d", studentIndex+3), student.Name,
			)

			f.SetCellValue(
				sheetName, fmt.Sprintf("D%d", studentIndex+3), "-",
			)
			f.SetCellStyle(
				sheetName,
				fmt.Sprintf("D%d", studentIndex+3),
				fmt.Sprintf("D%d", studentIndex+3),
				centerStyle,
			)

			mapRecap := make(map[string]int)
			for attendanceIndex, attendance := range *attendances {
				var status string

				if record, err := s.generalAttendanceRecordRepo.GetByAttendanceIdStudentId(
					attendance.Id, student.Id,
				); err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						status = "A"
					} else {
						return nil, failure.NewInternal(err)
					}
				} else {
					switch record.Status {
					case constants.AttendanceStatusPresent:
						recordTimeOnly := time.Date(
							0, 1, 1, record.DateTime.Hour(), record.DateTime.Minute(), 0, 0,
							record.DateTime.Location(),
						)
						attendanceTimeOnly := time.Date(
							0, 1, 1, attendance.DateTime.Hour(), attendance.DateTime.Minute(), 0, 0,
							attendance.DateTime.Location(),
						)

						if recordTimeOnly.After(attendanceTimeOnly) {
							status = "T"
						} else {
							status = "H"
						}
					case constants.AttendanceStatusSick:
						status = "S"
					case constants.AttendanceStatusPermission:
						status = "I"
					default:
						status = "A"
					}
				}

				columnName, err := utils.ColumnToName(attendanceIndex + 5)
				if err != nil {
					return nil, failure.NewInternal(err)
				}

				f.SetCellValue(
					sheetName, fmt.Sprintf("%s%d", columnName, studentIndex+3), status,
				)
				f.SetCellStyle(
					sheetName,
					fmt.Sprintf("%s%d", columnName, studentIndex+3),
					fmt.Sprintf("%s%d", columnName, studentIndex+3),
					centerStyle,
				)

				if _, ok := mapRecap[status]; !ok {
					mapRecap[status] = 1
				} else {
					mapRecap[status]++
				}
			}

			// rekap
			getCount := func(key string) string {
				count := mapRecap[key]
				if count == 0 {
					return "-"
				}
				return strconv.Itoa(count)
			}

			f.SetCellValue(
				sheetName, fmt.Sprintf("%s%d", presentColumn, studentIndex+3),
				getCount("H"),
			)
			f.SetCellStyle(
				sheetName,
				fmt.Sprintf("%s%d", presentColumn, studentIndex+3),
				fmt.Sprintf("%s%d", presentColumn, studentIndex+3),
				centerStyle,
			)

			f.SetCellValue(
				sheetName, fmt.Sprintf("%s%d", sickColumn, studentIndex+3),
				getCount("S"),
			)
			f.SetCellStyle(
				sheetName,
				fmt.Sprintf("%s%d", sickColumn, studentIndex+3),
				fmt.Sprintf("%s%d", sickColumn, studentIndex+3),
				centerStyle,
			)

			f.SetCellValue(
				sheetName,
				fmt.Sprintf("%s%d", permissionColumn, studentIndex+3),
				getCount("I"),
			)
			f.SetCellStyle(
				sheetName,
				fmt.Sprintf("%s%d", permissionColumn, studentIndex+3),
				fmt.Sprintf("%s%d", permissionColumn, studentIndex+3),
				centerStyle,
			)

			f.SetCellValue(
				sheetName,
				fmt.Sprintf("%s%d", alphaColumn, studentIndex+3),
				getCount("A"),
			)
			f.SetCellStyle(
				sheetName,
				fmt.Sprintf("%s%d", alphaColumn, studentIndex+3),
				fmt.Sprintf("%s%d", alphaColumn, studentIndex+3),
				centerStyle,
			)

			f.SetCellValue(
				sheetName,
				fmt.Sprintf("%s%d", lateColumn, studentIndex+3),
				getCount("T"),
			)
			f.SetCellStyle(
				sheetName,
				fmt.Sprintf("%s%d", lateColumn, studentIndex+3),
				fmt.Sprintf("%s%d", lateColumn, studentIndex+3),
				centerStyle,
			)
		}
	}

	// ubah file menjadi base64
	var b bytes.Buffer
	if err := f.Write(&b); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		encodedStr := base64.StdEncoding.EncodeToString(b.Bytes())
		return &responses.ExportGeneralAttendance{
			File: encodedStr,
			FileName: fmt.Sprintf(
				"Rekap Presensi Kehadiran %s - %s.xlsx",
				parsedStartDate.Format("Monday, 2 January 2006"),
				parsedEndDate.Format("Monday, 2 January 2006"),
			),
		}, nil
	}
}

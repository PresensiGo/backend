package services

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"api/internal/features/attendance/domains"
	"api/internal/features/attendance/dto"
	"api/internal/features/attendance/dto/requests"
	"api/internal/features/attendance/dto/responses"
	"api/internal/features/attendance/repositories"
	batch "api/internal/features/batch/repositories"
	classroom "api/internal/features/classroom/repositories"
	major "api/internal/features/major/repositories"
	studentRepo "api/internal/features/student/repositories"
	subjectDomain "api/internal/features/subject/domains"
	subjectRepo "api/internal/features/subject/repositories"
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

type SubjectAttendance struct {
	db                          *gorm.DB
	batchRepo                   *batch.Batch
	majorRepo                   *major.Major
	classroomRepo               *classroom.Classroom
	studentRepo                 *studentRepo.Student
	subjectAttendanceRepo       *repositories.SubjectAttendance
	subjectAttendanceRecordRepo *repositories.SubjectAttendanceRecord
	subjectRepo                 *subjectRepo.Subject
	userRepo                    *userRepo.User
}

func NewSubjectAttendance(
	db *gorm.DB,
	batchRepo *batch.Batch,
	majorRepo *major.Major,
	classroomRepo *classroom.Classroom,
	studentRepo *studentRepo.Student,
	subjectAttendanceRepo *repositories.SubjectAttendance,
	subjectAttendanceRecordRepo *repositories.SubjectAttendanceRecord,
	subjectRepo *subjectRepo.Subject,
	userRepo *userRepo.User,
) *SubjectAttendance {
	return &SubjectAttendance{
		db:                          db,
		batchRepo:                   batchRepo,
		majorRepo:                   majorRepo,
		classroomRepo:               classroomRepo,
		studentRepo:                 studentRepo,
		subjectAttendanceRepo:       subjectAttendanceRepo,
		subjectAttendanceRecordRepo: subjectAttendanceRecordRepo,
		subjectRepo:                 subjectRepo,
		userRepo:                    userRepo,
	}
}

func (s *SubjectAttendance) CreateSubjectAttendance(
	c *gin.Context, classroomId uint, req requests.CreateSubjectAttendance,
) (
	*responses.CreateSubjectAttendance, *failure.App,
) {
	user := authentication.GetAuthenticatedUser(c)
	if user.ID == 0 {
		return nil, failure.NewApp(http.StatusForbidden, "Anda tidak memiliki akses!", nil)
	}

	parsedDatetime, err := utils.GetParsedDateTime(req.DateTime)
	if err != nil {
		return nil, failure.NewApp(http.StatusBadRequest, "Tanggal dan waktu tidak valid!", err)
	}

	subjectAttendance := domains.SubjectAttendance{
		DateTime:    *parsedDatetime,
		Code:        uuid.NewString(),
		Note:        req.Note,
		ClassroomId: classroomId,
		SubjectId:   req.SubjectId,
		CreatorId:   user.ID,
	}

	if result, err := s.subjectAttendanceRepo.Create(subjectAttendance); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.CreateSubjectAttendance{
			SubjectAttendance: *result,
		}, nil
	}
}

func (s *SubjectAttendance) CreateSubjectAttendanceRecord(
	subjectAttendanceId uint, req requests.CreateSubjectAttendanceRecord,
) (*responses.CreateSubjectAttendanceRecord, *failure.App) {
	status := req.Status.ToAttendanceStatus()
	dateTime := time.Now()
	if req.Status == constants.AttendanceStatusTypePresentOnTime {
		if subjectAttendance, err := s.subjectAttendanceRepo.Get(subjectAttendanceId); err != nil {
			return nil, failure.NewInternal(err)
		} else {
			dateTime = subjectAttendance.DateTime
		}
	}

	if result, err := s.subjectAttendanceRecordRepo.FirstOrCreate(
		domains.SubjectAttendanceRecord{
			DateTime:            dateTime,
			SubjectAttendanceId: subjectAttendanceId,
			StudentId:           req.StudentId,
			Status:              status,
		},
	); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.CreateSubjectAttendanceRecord{
			SubjectAttendanceRecord: *result,
		}, nil
	}
}

func (s *SubjectAttendance) CreateSubjectAttendanceRecordStudent(
	studentId uint, req requests.CreateSubjectAttendanceRecordStudent,
) (
	*responses.CreateSubjectAttendanceRecordStudent, *failure.App,
) {
	subjectAttendance, err := s.subjectAttendanceRepo.GetByCode(req.Code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, failure.NewApp(http.StatusNotFound, "Kode akses tidak ditemukan!", err)
		}
		return nil, failure.NewInternal(err)
	}

	// validasi apakah siswa berasal dari kelas yang sama dengan subject attendance
	student, err := s.studentRepo.Get(studentId)
	if err != nil {
		return nil, failure.NewInternal(err)
	} else {
		if student.ClassroomId != subjectAttendance.ClassroomId {
			return nil, failure.NewApp(
				http.StatusForbidden,
				"Anda tidak terdaftar di kelas ini!",
				err,
			)
		}
	}

	if _, err := s.subjectAttendanceRecordRepo.FirstOrCreate(
		domains.SubjectAttendanceRecord{
			DateTime:            time.Now(),
			SubjectAttendanceId: subjectAttendance.Id,
			StudentId:           studentId,
			Status:              constants.AttendanceStatusPresent,
		},
	); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.CreateSubjectAttendanceRecordStudent{
			Message: "ok",
		}, nil
	}
}

func (s *SubjectAttendance) GetAllSubjectAttendances(classroomId uint) (
	*responses.GetAllSubjectAttendances, *failure.App,
) {
	subjectAttendances, err := s.subjectAttendanceRepo.GetAllByClassroomId(classroomId)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	subjectIds := make([]uint, len(*subjectAttendances))
	creatorIds := make([]uint, len(*subjectAttendances))
	for i, v := range *subjectAttendances {
		subjectIds[i] = v.SubjectId
		creatorIds[i] = v.CreatorId
	}

	mapSubjects := make(map[uint]*subjectDomain.Subject)
	mapCreators := make(map[uint]*userDomain.User)

	if subjects, err := s.subjectRepo.GetMany(subjectIds); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		for _, v := range *subjects {
			mapSubjects[v.Id] = &v
		}
	}

	if creators, err := s.userRepo.GetMany(creatorIds); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		for _, v := range *creators {
			mapCreators[v.Id] = &v
		}
	}

	result := make([]dto.GetAllSubjectAttendancesItem, len(*subjectAttendances))
	for i, v := range *subjectAttendances {
		var creator userDomain.User
		if v, ok := mapCreators[v.CreatorId]; ok {
			creator = *v
		} else {
			creator = userDomain.User{}
		}

		result[i] = dto.GetAllSubjectAttendancesItem{
			SubjectAttendance: v,
			Subject:           *mapSubjects[v.SubjectId],
			Creator:           creator,
		}
	}

	return &responses.GetAllSubjectAttendances{
		Items: result,
	}, nil
}

func (s *SubjectAttendance) GetAllSubjectAttendancesStudent(c *gin.Context) (
	*responses.GetAllSubjectAttendancesStudent, *failure.App,
) {
	auth := authentication.GetAuthenticatedStudent(c)
	if auth.Id == 0 || auth.SchoolId == 0 {
		return nil, failure.NewApp(http.StatusForbidden, "Anda tidak memiliki akses!", nil)
	}

	student, err := s.studentRepo.Get(auth.Id)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	attendances, err := s.subjectAttendanceRepo.GetAllTodayByClassroomId(student.ClassroomId)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	attendanceIds := make([]uint, len(*attendances))
	subjectIds := make([]uint, len(*attendances))
	creatorIds := make([]uint, len(*attendances))
	for i, v := range *attendances {
		attendanceIds[i] = v.Id
		subjectIds[i] = v.SubjectId
		creatorIds[i] = v.CreatorId
	}

	mapRecords := make(map[uint]*domains.SubjectAttendanceRecord)
	if records, err := s.subjectAttendanceRecordRepo.GetManyByAttendanceIdsStudentId(
		attendanceIds, auth.Id,
	); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		for _, record := range *records {
			mapRecords[record.SubjectAttendanceId] = &record
		}
	}

	mapSubjects := make(map[uint]*subjectDomain.Subject)
	if subjects, err := s.subjectRepo.GetMany(subjectIds); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		for _, subject := range *subjects {
			mapSubjects[subject.Id] = &subject
		}
	}

	mapCreators := make(map[uint]*userDomain.User)
	if creators, err := s.userRepo.GetMany(creatorIds); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		for _, v := range *creators {
			mapCreators[v.Id] = &v
		}
	}

	result := make([]dto.GetAllSubjectAttendancesStudentItem, len(*attendances))
	for i, v := range *attendances {
		record := domains.SubjectAttendanceRecord{}
		if r, ok := mapRecords[v.Id]; ok {
			record = *r
		}

		subject := subjectDomain.Subject{}
		if s, ok := mapSubjects[v.SubjectId]; ok {
			subject = *s
		}

		creator := userDomain.User{}
		if c, ok := mapCreators[v.CreatorId]; ok {
			creator = *c
		}

		result[i] = dto.GetAllSubjectAttendancesStudentItem{
			SubjectAttendance:       v,
			SubjectAttendanceRecord: record,
			Subject:                 subject,
			Creator:                 creator,
		}
	}

	return &responses.GetAllSubjectAttendancesStudent{
		Items: result,
	}, nil
}

func (s *SubjectAttendance) GetAllSubjectAttendanceRecords(
	classroomId uint, subjectAttendanceId uint,
) (
	*responses.GetAllSubjectAttendanceRecords, *failure.App,
) {
	students, err := s.studentRepo.GetAllByClassroomId(classroomId)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	records, err := s.subjectAttendanceRecordRepo.GetAllByAttendanceId(subjectAttendanceId)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	mapRecords := make(map[uint]*domains.SubjectAttendanceRecord)
	for _, record := range *records {
		mapRecords[record.StudentId] = &record
	}

	result := make([]dto.GetAllSubjectAttendanceRecordsItem, len(students))
	for i, student := range students {
		var record *domains.SubjectAttendanceRecord
		if r, ok := mapRecords[student.Id]; ok {
			record = r
		} else {
			record = &domains.SubjectAttendanceRecord{}
		}

		result[i] = dto.GetAllSubjectAttendanceRecordsItem{
			Student: student,
			Record:  *record,
		}
	}

	return &responses.GetAllSubjectAttendanceRecords{
		Items: result,
	}, nil
}

func (s *SubjectAttendance) GetSubjectAttendance(subjectAttendanceId uint) (
	*responses.GetSubjectAttendance, *failure.App,
) {
	if subjectAttendance, err := s.subjectAttendanceRepo.Get(subjectAttendanceId); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		response := responses.GetSubjectAttendance{
			SubjectAttendance: *subjectAttendance,
			Creator:           userDomain.User{},
		}

		// mendapatkan creator
		if subjectAttendance.CreatorId != 0 {
			if creator, err := s.userRepo.GetByID(subjectAttendance.CreatorId); err != nil {
				return nil, failure.NewInternal(err)
			} else {
				response.Creator = *creator
			}
		}

		return &response, nil
	}
}

func (s *SubjectAttendance) DeleteSubjectAttendance(attendanceId uint) (
	*responses.DeleteSubjectAttendance, *failure.App,
) {
	if err := s.subjectAttendanceRepo.Delete(attendanceId); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.DeleteSubjectAttendance{
			Message: "ok",
		}, nil
	}
}

func (s *SubjectAttendance) DeleteSubjectAttendanceRecord(recordId uint) (
	*responses.DeleteSubjectAttendanceRecord, *failure.App,
) {
	if err := s.subjectAttendanceRecordRepo.Delete(recordId); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.DeleteSubjectAttendanceRecord{
			Message: "ok",
		}, nil
	}
}

func (s *SubjectAttendance) ExportSubjectAttendance(c *gin.Context) (*excelize.File, *failure.App) {
	auth := authentication.GetAuthenticatedUser(c)
	if auth.SchoolId == 0 {
		return nil, failure.NewApp(http.StatusForbidden, "Anda tidak memiliki akses!", nil)
	}

	f := excelize.NewFile()
	if err := f.DeleteSheet("Sheet1"); err != nil {
		return nil, failure.NewInternal(err)
	}

	startDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.Local)
	endDate := time.Date(2025, 9, 8, 0, 0, 0, 0, time.Local)

	attendances, err := s.subjectAttendanceRepo.GetAllBySubjectIdBetween(4, startDate, endDate)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	mapClassroomIdAttendances := make(map[uint][]domains.SubjectAttendance)
	classroomIds := make([]uint, len(*attendances))
	for i, v := range *attendances {
		classroomIds[i] = v.ClassroomId
		mapClassroomIdAttendances[v.ClassroomId] = append(
			mapClassroomIdAttendances[v.ClassroomId], v,
		)
	}

	classrooms, err := s.classroomRepo.GetManyByIds(classroomIds)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	for _, c := range *classrooms {
		sheetIndex, err := f.NewSheet(c.Name)
		if err != nil {
			return nil, failure.NewInternal(err)
		}

		f.SetActiveSheet(sheetIndex)

		// header table
		f.SetCellValue(c.Name, "A1", "No")
		f.MergeCell(c.Name, "A1", "A2")
		f.SetColWidth(c.Name, "A", "A", 4)

		f.SetCellValue(c.Name, "B1", "NIS")
		f.MergeCell(c.Name, "B1", "B2")
		f.SetColWidth(c.Name, "B", "B", 8)

		f.SetCellValue(c.Name, "C1", "Nama")
		f.MergeCell(c.Name, "C1", "C2")
		f.SetColWidth(c.Name, "C", "C", 32)

		f.SetCellValue(c.Name, "D1", "JK")
		f.MergeCell(c.Name, "D1", "D2")
		f.SetColWidth(c.Name, "D", "D", 4)

		attendances := mapClassroomIdAttendances[c.Id]
		mapMonths := make(map[string]int)

		for i, attendance := range attendances {
			columnName, err := utils.ColumnToName(i + 5)
			if err != nil {
				return nil, failure.NewInternal(err)
			}
			monthName := attendance.DateTime.Format("January")
			if _, ok := mapMonths[monthName]; !ok {
				mapMonths[monthName] = 1
			} else {
				mapMonths[monthName]++
			}

			f.SetCellValue(
				c.Name, fmt.Sprintf("%s1", columnName),
				monthName,
			)
			f.SetCellValue(
				c.Name, fmt.Sprintf("%s2", columnName),
				attendance.DateTime.Format("2"),
			)
			f.SetColWidth(c.Name, columnName, columnName, 4)
		}

		// rekap
		presentColumn, err := utils.ColumnToName(len(attendances) + 5)
		if err != nil {
			return nil, failure.NewInternal(err)
		}
		sickColumn, err := utils.ColumnToName(len(attendances) + 6)
		if err != nil {
			return nil, failure.NewInternal(err)
		}
		permissionColumn, err := utils.ColumnToName(len(attendances) + 7)
		if err != nil {
			return nil, failure.NewInternal(err)
		}
		alphaColumn, err := utils.ColumnToName(len(attendances) + 8)
		if err != nil {
			return nil, failure.NewInternal(err)
		}
		f.SetCellValue(c.Name, fmt.Sprintf("%s2", presentColumn), "H")
		f.SetColWidth(c.Name, presentColumn, presentColumn, 4)

		f.SetCellValue(c.Name, fmt.Sprintf("%s2", sickColumn), "S")
		f.SetColWidth(c.Name, sickColumn, sickColumn, 4)

		f.SetCellValue(c.Name, fmt.Sprintf("%s2", permissionColumn), "I")
		f.SetColWidth(c.Name, permissionColumn, permissionColumn, 4)

		f.SetCellValue(c.Name, fmt.Sprintf("%s2", alphaColumn), "A")
		f.SetColWidth(c.Name, alphaColumn, alphaColumn, 4)

		f.SetCellValue(c.Name, fmt.Sprintf("%s1", presentColumn), "Rekap")
		f.MergeCell(c.Name, fmt.Sprintf("%s1", presentColumn), fmt.Sprintf("%s1", alphaColumn))

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
					c.Name, fmt.Sprintf("%s1", startColumn),
					fmt.Sprintf("%s1", endColumn),
				); err != nil {
					return nil, failure.NewInternal(err)
				}
			}
			currentIndex += v
		}

		students, err := s.studentRepo.GetAllByClassroomId(c.Id)
		if err != nil {
			return nil, failure.NewInternal(err)
		}

		for i, student := range students {
			f.SetCellValue(
				c.Name, fmt.Sprintf("A%d", i+3), i+1,
			)
			f.SetCellValue(
				c.Name, fmt.Sprintf("B%d", i+3), student.NIS,
			)
			f.SetCellValue(
				c.Name, fmt.Sprintf("C%d", i+3), student.Name,
			)
			f.SetCellValue(
				c.Name, fmt.Sprintf("D%d", i+3), "-",
			)

			recap := make(map[string]int)
			for j, attendance := range attendances {
				var status string

				if record, err := s.subjectAttendanceRecordRepo.GetByAttendanceIdStudentId(
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
						status = "H"
					case constants.AttendanceStatusSick:
						status = "S"
					case constants.AttendanceStatusPermission:
						status = "I"
					default:
						status = "A"
					}
				}

				columnName, err := utils.ColumnToName(j + 5)
				if err != nil {
					return nil, failure.NewInternal(err)
				}

				f.SetCellValue(
					c.Name, fmt.Sprintf("%s%d", columnName, i+3), status,
				)

				if _, ok := recap[status]; !ok {
					recap[status] = 1
				} else {
					recap[status]++
				}
			}

			// rekap
			f.SetCellValue(
				c.Name, fmt.Sprintf("%s%d", presentColumn, i+3),
				recap["H"],
			)
			f.SetCellValue(
				c.Name, fmt.Sprintf("%s%d", sickColumn, i+3),
				recap["S"],
			)
			f.SetCellValue(
				c.Name,
				fmt.Sprintf("%s%d", permissionColumn, i+3),
				recap["I"],
			)
			f.SetCellValue(
				c.Name,
				fmt.Sprintf("%s%d", alphaColumn, i+3),
				recap["A"],
			)

		}

	}

	return f, nil
}

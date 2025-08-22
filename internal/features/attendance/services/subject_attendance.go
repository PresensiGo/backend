package services

import (
	"errors"
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
	parsedDateTime, err := utils.GetParsedDateTime(req.DateTime)
	if err != nil {
		return nil, failure.NewApp(http.StatusBadRequest, "Tanggal dan waktu tidak valid!", err)
	}

	if result, err := s.subjectAttendanceRecordRepo.FirstOrCreate(
		domains.SubjectAttendanceRecord{
			DateTime:            *parsedDateTime,
			SubjectAttendanceId: subjectAttendanceId,
			StudentId:           req.StudentId,
			Status:              req.Status,
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

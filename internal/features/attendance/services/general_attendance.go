package services

import (
	"net/http"

	"api/internal/features/attendance/domains"
	"api/internal/features/attendance/dto"
	"api/internal/features/attendance/dto/requests"
	"api/internal/features/attendance/dto/responses"
	"api/internal/features/attendance/repositories"
	studentDomain "api/internal/features/student/domains"
	studentRepo "api/internal/features/student/repositories"
	userDomain "api/internal/features/user/domains"
	userRepo "api/internal/features/user/repositories"
	"api/pkg/authentication"
	"api/pkg/http/failure"
	"api/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GeneralAttendance struct {
	db                          *gorm.DB
	studentRepo                 *studentRepo.Student
	generalAttendanceRepo       *repositories.GeneralAttendance
	generalAttendanceRecordRepo *repositories.GeneralAttendanceRecord
	userRepo                    *userRepo.User
}

func NewGeneralAttendance(
	db *gorm.DB,
	studentRepo *studentRepo.Student,
	generalAttendanceRepo *repositories.GeneralAttendance,
	generalAttendanceRecordRepo *repositories.GeneralAttendanceRecord,
	userRepo *userRepo.User,
) *GeneralAttendance {
	return &GeneralAttendance{
		db:                          db,
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
) (*responses.CreateGeneralAttendanceRecordStudent, error) {
	generalAttendance, err := s.generalAttendanceRepo.GetByCode(req.Code)
	if err != nil {
		return nil, err
	}

	if err := s.db.Transaction(
		func(tx *gorm.DB) error {
			// hapus record lama jika sudah ada
			if err := s.generalAttendanceRecordRepo.DeleteByAttendanceIdStudentIdInTx(
				tx, generalAttendance.Id, studentId,
			); err != nil {
				return err
			}

			// create record baru
			generalAttendanceRecord := domains.GeneralAttendanceRecord{
				GeneralAttendanceId: generalAttendance.Id,
				StudentId:           studentId,
			}
			if _, err := s.generalAttendanceRecordRepo.CreateInTx(
				tx, generalAttendanceRecord,
			); err != nil {
				return err
			}

			return nil
		},
	); err != nil {
		return nil, err
	}

	return &responses.CreateGeneralAttendanceRecordStudent{
		Message: "ok",
	}, nil
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

package services

import (
	"api/internal/dto"
	"api/internal/dto/combined"
	"api/internal/dto/requests"
	"api/internal/dto/responses"
	"api/internal/repositories"
	"api/pkg/utils"
)

type Lateness struct {
	latenessRepo       *repositories.Lateness
	latenessDetailRepo *repositories.LatenessDetail
	studentRepo        *repositories.Student
	majorRepo          *repositories.Major
	classroomRepo      *repositories.Classroom
}

func NewLateness(
	latenessRepo *repositories.Lateness, latenessDetailRepo *repositories.LatenessDetail,
	studentRepo *repositories.Student, majorRepo *repositories.Major,
	classroomRepo *repositories.Classroom,
) *Lateness {
	return &Lateness{
		latenessRepo, latenessDetailRepo,
		studentRepo, majorRepo, classroomRepo,
	}
}

func (s *Lateness) Create(
	schoolId uint,
	req *requests.CreateLateness,
) error {
	parsedDate, err := utils.GetParsedDate(req.Date)
	if err != nil {
		return err
	}

	if _, err := s.latenessRepo.Create(
		&dto.Lateness{
			Date:     *parsedDate,
			SchoolId: schoolId,
		},
	); err != nil {
		return err
	}

	return nil
}

func (s *Lateness) CreateDetail(latenessId uint, req *requests.CreateLatenessDetail) error {
	latenessDetails := make([]dto.LatenessDetail, len(req.StudentIds))
	for i, item := range req.StudentIds {
		latenessDetails[i] = dto.LatenessDetail{
			LatenessId: latenessId,
			StudentId:  item,
		}
	}

	return s.latenessDetailRepo.CreateBatch(&latenessDetails)
}

func (s *Lateness) GetAllBySchoolId(schoolId uint) (*responses.GetAllLatenesses, error) {
	latenesses, err := s.latenessRepo.GetAllBySchoolId(schoolId)
	if err != nil {
		return nil, err
	}

	return &responses.GetAllLatenesses{
		Latenesses: *latenesses,
	}, nil
}

func (s *Lateness) GetDetail(latenessId uint) (*responses.GetLateness, error) {
	// get lateness
	lateness, err := s.latenessRepo.GetById(latenessId)
	if err != nil {
		return nil, err
	}

	// get lateness details
	latenessDetails, err := s.latenessDetailRepo.GetAllByLatenessId(latenessId)
	if err != nil {
		return nil, err
	}

	// get students
	studentIds := make([]uint, len(*latenessDetails))
	for i, item := range *latenessDetails {
		studentIds[i] = item.StudentId
	}
	students, err := s.studentRepo.GetManyById(studentIds)
	if err != nil {
		return nil, err
	}

	// get classrooms
	classroomIds := make([]uint, len(*students))
	mapClassrooms := make(map[uint]*dto.Classroom)
	for i, item := range *students {
		classroomIds[i] = item.ClassroomId
	}
	classrooms, err := s.classroomRepo.GetManyByIds(classroomIds)
	if err != nil {
		return nil, err
	}

	// get majors
	majorIds := make([]uint, len(*classrooms))
	mapMajors := make(map[uint]*dto.Major)
	for i, item := range *classrooms {
		majorIds[i] = item.MajorId
		mapClassrooms[item.Id] = &item
	}
	majors, err := s.majorRepo.GetManyByIds(majorIds)
	if err != nil {
		return nil, err
	}

	for _, item := range *majors {
		mapMajors[item.Id] = &item
	}

	// mapping items
	items := make([]combined.StudentMajorClassroom, len(*students))
	for i, item := range *students {
		classroom := mapClassrooms[item.ClassroomId]
		major := mapMajors[classroom.MajorId]

		items[i] = combined.StudentMajorClassroom{
			Student:   item,
			Major:     *major,
			Classroom: *classroom,
		}
	}

	return &responses.GetLateness{
		Lateness: *lateness,
		Items:    items,
	}, nil
}

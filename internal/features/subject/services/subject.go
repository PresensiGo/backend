package services

import (
	"api/internal/features/subject/domains"
	"api/internal/features/subject/dto/requests"
	"api/internal/features/subject/dto/responses"
	"api/internal/features/subject/repositories"
)

type Subject struct {
	subjectRepo *repositories.Subject
}

func NewSubjectRepo(subjectRepo *repositories.Subject) *Subject {
	return &Subject{
		subjectRepo: subjectRepo,
	}
}

func (s *Subject) Create(schoolId uint, req requests.CreateSubject) (
	*responses.CreateSubject, error,
) {
	subject := domains.Subject{
		Name:     req.Name,
		SchoolId: schoolId,
	}

	result, err := s.subjectRepo.Create(subject)
	if err != nil {
		return nil, err
	}

	return &responses.CreateSubject{
		Subject: *result,
	}, nil
}

func (s *Subject) GetAll(schoolId uint) (*responses.GetAllSubjects, error) {
	subjects, err := s.subjectRepo.GetAll(schoolId)
	if err != nil {
		return nil, err
	}

	return &responses.GetAllSubjects{
		Subjects: *subjects,
	}, nil
}

func (s *Subject) Update(subjectId uint, req requests.UpdateSubject) (
	*responses.UpdateSubject, error,
) {
	subject := domains.Subject{
		Name: req.Name,
	}

	result, err := s.subjectRepo.Update(subjectId, subject)
	if err != nil {
		return nil, err
	}

	return &responses.UpdateSubject{
		Subject: *result,
	}, nil
}

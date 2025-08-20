package services

import (
	"api/internal/features/subject/domains"
	"api/internal/features/subject/dto/requests"
	"api/internal/features/subject/dto/responses"
	"api/internal/features/subject/repositories"
	"api/pkg/http/failure"
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

func (s *Subject) GetAllSubjects(schoolId uint) (*responses.GetAllSubjects, *failure.App) {
	if subjects, err := s.subjectRepo.GetAll(schoolId); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.GetAllSubjects{
			Subjects: *subjects,
		}, nil
	}
}

func (s *Subject) GetSubject(subjectId uint) (*responses.GetSubject, *failure.App) {
	if subject, err := s.subjectRepo.Get(subjectId); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.GetSubject{
			Subject: *subject,
		}, nil
	}
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

func (s *Subject) Delete(subjectId uint) (*responses.DeleteSubject, error) {
	if err := s.subjectRepo.Delete(subjectId); err != nil {
		return nil, err
	}

	return &responses.DeleteSubject{
		Message: "ok",
	}, nil
}

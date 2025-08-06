package injector

import (
	"api/internal/features/subject/handlers"
	"api/internal/features/subject/repositories"
	"api/internal/features/subject/services"
	"api/pkg/database"
	"github.com/google/wire"
)

type SubjectHandlers struct {
	Subject *handlers.Subject
}

func NewSubjectHandlers(subject *handlers.Subject) *SubjectHandlers {
	return &SubjectHandlers{Subject: subject}
}

var (
	SubjectSet = wire.NewSet(
		handlers.NewSubject,
		services.NewSubjectRepo,
		repositories.NewSubject,
		database.New,
	)
)

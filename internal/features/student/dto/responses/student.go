package responses

import (
	batch "api/internal/features/batch/domains"
	classroom "api/internal/features/classroom/domains"
	major "api/internal/features/major/domains"
	school "api/internal/features/school/domains"
	"api/internal/features/student/domains"
	"api/internal/features/student/dto"
)

type GetAllStudentsByClassroomId struct {
	Students []domains.Student `json:"students" validate:"required"`
} // @name GetAllStudentsByClassroomIdRes

type GetAllStudentAccountsByClassroomId struct {
	Items []dto.StudentAccount `json:"items" validate:"required"`
}

// type GetAllStudents struct {
// 	Students []school.StudentMajorClassroom `json:"students" validate:"required"`
// } // @name GetAllStudentsRes

type GetProfileStudent struct {
	Student   domains.Student     `json:"student" validate:"required"`
	School    school.School       `json:"school" validate:"required"`
	Classroom classroom.Classroom `json:"classroom" validate:"required"`
	Major     major.Major         `json:"major" validate:"required"`
	Batch     batch.Batch         `json:"batch" validate:"required"`
} // @name GetProfileStudentRes

type DeleteStudent struct {
	Message string `json:"message" validate:"required"`
}

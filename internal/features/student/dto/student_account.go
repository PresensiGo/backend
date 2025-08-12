package dto

import "api/internal/features/student/domains"

type StudentAccount struct {
	Student      domains.Student      `json:"student" validate:"required"`
	StudentToken domains.StudentToken `json:"student_token" validate:"required"`
}

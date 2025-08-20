package dto

import "api/internal/features/classroom/domains"

type GetAllClassroomsByMajorIdItem struct {
	Classroom    domains.Classroom `json:"classroom" validate:"required"`
	StudentCount int64             `json:"student_count" validate:"required"`
} // @name GetAllClassroomsByMajorIdItem

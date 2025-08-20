package dto

import "api/internal/features/major/domains"

type GetAllMajorsByBatchIdItem struct {
	Major          domains.Major `json:"major" validate:"required"`
	ClassroomCount int64         `json:"classroom_count" validate:"required"`
} // @name GetAllMajorsByBatchIdItem

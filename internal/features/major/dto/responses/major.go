package responses

import (
	"api/internal/features/major/domains"
	"api/internal/features/major/dto"
)

// type GetAllMajors struct {
// 	Items []domains.Major `json:"majors" validate:"required"`
// }

type GetAllMajorsByBatchId struct {
	Items []dto.GetAllMajorsByBatchIdItem `json:"items" validate:"required"`
} // @name GetAllMajorsByBatchIdRes

type GetMajor struct {
	Major domains.Major `json:"major" validate:"required"`
} // @name GetMajorRes

package dto

import "api/internal/features/batch/domains"

type GetAllBatchesItem struct {
	Batch      domains.Batch `json:"batch" validate:"required"`
	MajorCount int64         `json:"major_count" validate:"required"`
} // @name GetAllBatchesItem

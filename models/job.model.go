package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Job struct {
	ID        int64     `gorm:"size:20; primary_key" json:"id,omitempty"`
	JobName   string    `gorm:"size:50; not null" json:"job_name,omitempty"`
	JobStatus *bool     `gorm:"default:1; not null;" json:"job_status,omitempty"`
	JobSlug   string    `gorm:"size:50; not null; unique" json:"job_slug,omitempty"`
	CreatedAt time.Time `gorm:"" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"" json:"updated_at,omitempty"`
}

func ValidateStructJob[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

type CreateJobSchema struct {
	JobName string `json:"job_name" validate:"required"`
	JobSlug string `json:"job_slug,omitempty"`
}

type UpdateJobSchema struct {
	JobName string `json:"job_name,omitempty"`
	JobSlug string `json:"job_slug,omitempty"`
}

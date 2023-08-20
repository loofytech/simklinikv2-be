package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Education struct {
	ID              int64     `gorm:"size:20; primary_key" json:"id,omitempty"`
	EducationName   string    `gorm:"size:50; not null" json:"education_name,omitempty"`
	EducationActive *bool     `gorm:"default:1; not null;" json:"education_active,omitempty"`
	EducationSlug   string    `gorm:"size:50; not null; unique" json:"education_slug,omitempty"`
	CreatedAt       time.Time `gorm:"" json:"created_at,omitempty"`
	UpdatedAt       time.Time `gorm:"" json:"updated_at,omitempty"`
}

func ValidateStructEducation[T any](payload T) []*ErrorResponse {
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

type CreateEducationSchema struct {
	EducationName string `json:"education_name" validate:"required"`
	EducationSlug string `json:"education_slug,omitempty"`
}

type UpdateEducationSchema struct {
	EducationName string `json:"education_name,omitempty"`
	EducationSlug string `json:"education_slug,omitempty"`
}

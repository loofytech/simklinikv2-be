package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type MaritalStatus struct {
	ID            int64     `gorm:"size:20; primary_key" json:"id,omitempty"`
	MaritalName   string    `gorm:"size:50; not null" json:"marital_name,omitempty"`
	MaritalActive *bool     `gorm:"default:1; not null; unique" json:"marital_active,omitempty"`
	MaritalSlug   string    `gorm:"size:50; not null; unique" json:"marital_slug,omitempty"`
	CreatedAt     time.Time `gorm:"" json:"created_at,omitempty"`
	UpdatedAt     time.Time `gorm:"" json:"updated_at,omitempty"`
}

func ValidateStructMarital[T any](payload T) []*ErrorResponse {
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

type CreateMaritalStatusSchema struct {
	MaritalName   string `json:"marital_name" validate:"required"`
	MaritalActive *bool  `json:"marital_active" validate:"required"`
	MaritalSlug   string `json:"marital_slug,omitempty"`
}

type UpdateMaritalStatusSchema struct {
	MaritalName   string `json:"marital_name,omitempty"`
	MaritalActive *bool  `json:"marital_active,omitempty"`
	MaritalSlug   string `json:"marital_slug,omitempty"`
}

package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Religion struct {
	ID             int64     `gorm:"size:20; primary_key" json:"id,omitempty"`
	ReligionName   string    `gorm:"size:50; not null" json:"religion_name,omitempty"`
	ReligionActive *bool     `gorm:"default:1; not null; unique" json:"religion_active,omitempty"`
	ReligionSlug   string    `gorm:"size:50; not null; unique" json:"religion_slug,omitempty"`
	CreatedAt      time.Time `gorm:"" json:"created_at,omitempty"`
	UpdatedAt      time.Time `gorm:"" json:"updated_at,omitempty"`
}

func ValidateStructReligion[T any](payload T) []*ErrorResponse {
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

type CreateReligionSchema struct {
	ReligionName   string `json:"religion_name" validate:"required"`
	ReligionActive *bool  `json:"religion_active" validate:"required"`
	ReligionSlug   string `json:"religion_slug,omitempty"`
}

type UpdateReligionSchema struct {
	ReligionName   string `json:"religion_name,omitempty"`
	ReligionActive *bool  `json:"religion_active,omitempty"`
	ReligionSlug   string `json:"religion_slug,omitempty"`
}

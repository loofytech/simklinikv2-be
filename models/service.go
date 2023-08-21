package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	ID            int64     `gorm:"size:20; primary_key" json:"id,omitempty"`
	ServiceName   string    `gorm:"size:50; not null" json:"service_name,omitempty"`
	ServiceActive *bool     `gorm:"default:1; not null" json:"service_active,omitempty"`
	ServiceSlug   string    `gorm:"size:50; not null; unique" json:"service_slug,omitempty"`
	CreatedAt     time.Time `gorm:"" json:"created_at,omitempty"`
	UpdatedAt     time.Time `gorm:"" json:"updated_at,omitempty"`
}

func ValidateStructService[T any](payload T) []*ErrorResponse {
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

type CreateServiceSchema struct {
	ServiceName string `json:"service_name" validate:"required"`
	ServiceSlug string `json:"service_slug,omitempty"`
}

type UpdateServiceSchema struct {
	ServiceName string `json:"service_name,omitempty"`
	ServiceSlug string `json:"service_slug,omitempty"`
}

package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Ethnic struct {
	ID           int64     `gorm:"size:20; primary_key" json:"id,omitempty"`
	EthnicName   string    `gorm:"size:50; not null" json:"ethnic_name,omitempty"`
	EthnicActive *bool     `gorm:"default:1; not null; unique" json:"ethnic_active,omitempty"`
	EthnicSlug   string    `gorm:"size:50; not null; unique" json:"ethnic_slug,omitempty"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

func ValidateStructEthnic[T any](payload T) []*ErrorResponse {
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

type CreateEthnicSchema struct {
	EthnicName   string `json:"ethnic_name" validate:"required"`
	EthnicActive *bool  `json:"ethnic_active" validate:"required"`
	EthnicSlug   string `json:"ethnic_slug,omitempty"`
}

type UpdateEthnicSchema struct {
	EthnicName   string `json:"ethnic_name,omitempty"`
	EthnicActive *bool  `json:"ethnic_active,omitempty"`
	EthnicSlug   string `json:"ethnic_slug,omitempty"`
}

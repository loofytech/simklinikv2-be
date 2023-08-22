package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Unit struct {
	ID         int64     `gorm:"size:20; primary_key" json:"id,omitempty"`
	UnitName   string    `gorm:"size:50; not null" json:"unit_name,omitempty"`
	UnitStatus *bool     `gorm:"default:1; not null" json:"unit_status,omitempty"`
	UnitSlug   string    `gorm:"size:50; not null; unique" json:"unit_slug,omitempty"`
	ServiceId  int64     `gorm:"size:20; foreign_key" json:"service_id,omitempty"`
	Service    Service   `gorm:"references:id"`
	CreatedAt  time.Time `gorm:"" json:"created_at,omitempty"`
	UpdatedAt  time.Time `gorm:"" json:"updated_at,omitempty"`
}

func ValidateStructUnit[T any](payload T) []*ErrorResponse {
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

type CreateUnitSchema struct {
	UnitName   string `json:"unit_name" validate:"required"`
	UnitStatus *bool  `json:"unit_status,omitempty"`
	UnitSlug   string `json:"unit_slug,omitempty"`
	ServiceId  int64  `json:"service_id" validate:"required"`
}

type UpdateUnitSchema struct {
	UnitName   string `json:"unit_name,omitempty"`
	UnitStatus *bool  `json:"unit_status,omitempty"`
	UnitSlug   string `json:"unit_slug,omitempty"`
}

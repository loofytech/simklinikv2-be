package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type ClinicRate struct {
	ID               int64     `gorm:"size:20; primary_key" json:"id,omitempty"`
	ClinicRateName   string    `gorm:"size:50; not null" json:"rate_name,omitempty"`
	ClinicRatePrice  int64     `gorm:"size:50; not null" json:"rate_price,omitempty"`
	ClinicRateStatus *bool     `gorm:"default:1; not null;" json:"rate_status,omitempty"`
	UnitId           int64     `gorm:"size:20; foreign_key" json:"unit_id,omitempty"`
	Unit             Unit      `gorm:"references:id" json:"unit"`
	CreatedAt        time.Time `gorm:"" json:"created_at,omitempty"`
	UpdatedAt        time.Time `gorm:"" json:"updated_at,omitempty"`
}

func ValidateStructClinicRate[T any](payload T) []*ErrorResponse {
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

type CreateClinicRateSchema struct {
	ClinicRateName   string `json:"rate_name" validate:"required"`
	ClinicRatePrice  int64  `json:"rate_price" validate:"required"`
	ClinicRateStatus *bool  `json:"rate_status" validate:"required"`
	UnitId           int64  `json:"unit_id" validate:"required"`
}

type UpdateClinicRateSchema struct {
	ClinicRateName   string `json:"rate_name,omitempty"`
	ClinicRatePrice  int64  `json:"rate_price,omitempty"`
	ClinicRateStatus *bool  `json:"rate_status,omitempty"`
	UnitId           int64  `json:"unit_id,omitempty"`
}

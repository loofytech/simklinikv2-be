package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Measure struct {
	ID             int64        `gorm:"size:20; primary_key" json:"id,omitempty"`
	RegistrationId int64        `gorm:"size:20; foreign_key" json:"registration_id,omitempty"`
	ClinicRateId   int64        `gorm:"size:20; foreign_key" json:"clinic_rate_id,omitempty"`
	NakesFirst     string       `gorm:"size:50; not null" json:"nakes_first,omitempty"`
	NakesSecond    string       `gorm:"size:50; not null" json:"nakes_second,omitempty"`
	SubTotal       float64      `gorm:"size:50; not null" json:"sub_total,omitempty"`
	Registration   Registration `gorm:"references:id" json:"registration"`
	ClinicRate     ClinicRate   `gorm:"references:id" json:"clinic_rate"`
	CreatedAt      time.Time    `gorm:"" json:"created_at,omitempty"`
	UpdatedAt      time.Time    `gorm:"" json:"updated_at,omitempty"`
}

func ValidateStructMeasure[T any](payload T) []*ErrorResponse {
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

type CreateMeasureSchema struct {
	RegistrationId int64   `json:"registration_id" validate:"required"`
	ClinicRateId   int64   `json:"clinic_rate" validate:"required"`
	NakesFirst     string  `json:"nakes_first" validate:"required"`
	NakesSecond    string  `json:"nakes_second" validate:"required"`
	SubTotal       float64 `json:"sub_total" validate:"required"`
}

type UpdateMeasureSchema struct {
	RegistrationId int64   `json:"registration_id,omitempty"`
	ClinicRateId   int64   `json:"clinic_rate,omitempty"`
	NakesFirst     string  `json:"nakes_first,omitempty"`
	NakesSecond    string  `json:"nakes_second,omitempty"`
	SubTotal       float64 `json:"sub_total,omitempty"`
}

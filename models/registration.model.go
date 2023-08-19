package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Registration struct {
	ID                  int64     `gorm:"size:20; primary_key" json:"id,omitempty"`
	PatientId           int64     `gorm:"size:20; foreign_key" json:"patient_id,omitempty"`
	ResponsibleName     string    `gorm:"size:50; not null" json:"responsible_name,omitempty"`
	ResponsiblePhone    string    `gorm:"size:50; not null" json:"responsible_phone,omitempty"`
	ResponsibleAddress  string    `gorm:"size:50; not null" json:"responsible_address,omitempty"`
	ResponsibleRelation string    `gorm:"size:50; not null" json:"responsible_relation,omitempty"`
	ServiceId           int64     `gorm:"size:20; foreign_key" json:"service_id,omitempty"`
	Patient             Patient   `gorm:"references:id"`
	Service             Service   `gorm:"references:id"`
	CreatedAt           time.Time `gorm:"" json:"created_at,omitempty"`
	UpdatedAt           time.Time `gorm:"" json:"updated_at,omitempty"`
}

func ValidateStructRegistration[T any](payload T) []*ErrorResponse {
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

type CreateRegistrationSchema struct {
	ResponsibleName     string `json:"responsible_name" validate:"required"`
	ResponsiblePhone    string `json:"responsible_phone" validate:"required"`
	ResponsibleAddress  string `json:"responsible_address" validate:"required"`
	ResponsibleRelation string `json:"responsible_relation" validate:"required"`
}

type UpdateRegistrationSchema struct {
	ResponsibleName     string `json:"responsible_name,omitempty"`
	ResponsiblePhone    string `json:"responsible_phone,omitempty"`
	ResponsibleAddress  string `json:"responsible_address,omitempty"`
	ResponsibleRelation string `json:"responsible_relation,omitempty"`
}

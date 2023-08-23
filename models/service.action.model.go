package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type ServiceAction struct {
	ID             int64        `gorm:"size:20; primary_key" json:"id,omitempty"`
	UnitId         int64        `gorm:"size:20; foreign_key" json:"unit_id,omitempty"`
	RegistrationId int64        `gorm:"size:20; foreign_key" json:"registration_id,omitempty"`
	UserId         int64        `gorm:"size:20; foreign_key" json:"user_id,omitempty"`
	User           User         `gorm:"references:id" json:"user"`
	Unit           Unit         `gorm:"references:id" json:"unit"`
	Registration   Registration `gorm:"references:id"`
	CreatedAt      time.Time    `gorm:"" json:"created_at,omitempty"`
	UpdatedAt      time.Time    `gorm:"" json:"updated_at,omitempty"`
}

func ValidateStructServiceAction[T any](payload T) []*ErrorResponse {
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

type CreateServiceActionSchema struct {
	UnitId         int64 `gorm:"unit_id" validate:"required"`
	RegistrationId int64 `gorm:"registration_id" validate:"required"`
	UserId         int64 `gorm:"user_id" validate:"required"`
}

type UpdateServiceActionSchema struct {
	UnitId         int64 `gorm:"unit_id,omitempty"`
	RegistrationId int64 `gorm:"registration,omitempty"`
	UserId         int64 `gorm:"user_id,omitempty"`
}

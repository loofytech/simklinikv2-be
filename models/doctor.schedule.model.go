package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type DoctorSchedule struct {
	ID            int64     `gorm:"size:20; primary_key" json:"id,omitempty"`
	UserId        int64     `gorm:"size:20; foreign_key" json:"user_id,omitempty"`
	UnitId        int64     `gorm:"size:20; foreign_key" json:"unit_id,omitempty"`
	User          User      `gorm:"references:id" json:"user"`
	Unit          Unit      `gorm:"references:id" json:"unit"`
	Day           string    `gorm:"size:10; not null" json:"day,omitempty"`
	OpenPractice  string    `gorm:"size:10; not null" json:"open_practice,omitempty"`
	ClosePractice string    `gorm:"size:10; not null" json:"close_practice,omitempty"`
	CreatedAt     time.Time `gorm:"" json:"created_at,omitempty"`
	UpdatedAt     time.Time `gorm:"" json:"updated_at,omitempty"`
}

func ValidateStructDoctorSchedule[T any](payload T) []*ErrorResponse {
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

type CreateDoctorScheduleSchema struct {
	UserId        int64  `json:"user_id" validate:"required"`
	UnitId        int64  `json:"unit_id" validate:"required"`
	Day           string `json:"day" validate:"required"`
	OpenPractice  string `json:"open_practice" validate:"required"`
	ClosePractice string `json:"close_practice" validate:"required"`
}

type UpdateDoctorScheduleSchema struct {
	UserId        int64  `json:"user_id,omitempty"`
	UnitId        int64  `json:"unit_id,omitempty"`
	Day           string `json:"day,omitempty"`
	OpenPractice  string `json:"open_practice,omitempty"`
	ClosePractice string `json:"close_practice,omitempty"`
}

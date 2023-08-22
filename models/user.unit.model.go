package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type UserUnit struct {
	ID        int64     `gorm:"size:20; primary_key" json:"id,omitempty"`
	UserId    int64     `gorm:"size:20; foreign_key" json:"user_id,omitempty"`
	UnitId    int64     `gorm:"size:20; foreign_key" json:"unit_id,omitempty"`
	User      User      `gorm:"references:id"`
	Unit      Unit      `gorm:"references:id"`
	CreatedAt time.Time `gorm:"" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"" json:"updated_at,omitempty"`
}

func ValidateStructUserUnit[T any](payload T) []*ErrorResponse {
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

type CreateUserUnitSchema struct {
	UserId int64 `json:"user_id" validate:"required"`
	UnitId int64 `json:"unit_id" validate:"required"`
}

type UpdateUserUnitSchema struct {
	UserId int64 `json:"user_id,omitempty"`
	UnitId int64 `json:"unit_id,omitempty"`
}

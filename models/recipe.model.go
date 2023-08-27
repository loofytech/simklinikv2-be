package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Recipe struct {
	ID             int64        `gorm:"size:20; primary_key" json:"id,omitempty"`
	RegistrationId int64        `gorm:"size:20; foreign_key" json:"registration_id,omitempty"`
	MedicineId     int64        `gorm:"size:20; foreign_key" json:"medicine_id,omitempty"`
	Quantity       int64        `gorm:"size:50; not null" json:"quantity,omitempty"`
	SubTotal       float64      `gorm:"size:50; not null" json:"sub_total,omitempty"`
	Registration   Registration `gorm:"references:id" json:"registration"`
	Medicine       Medicine     `gorm:"references:id" json:"medicine"`
	CreatedAt      time.Time    `gorm:"" json:"created_at,omitempty"`
	UpdatedAt      time.Time    `gorm:"" json:"updated_at,omitempty"`
}

func ValidateStructRecipe[T any](payload T) []*ErrorResponse {
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

type CreateRecipeSchema struct {
	RegistrationId int64   `json:"registration_id" validate:"required"`
	MedicineId     int64   `json:"medicine_id" validate:"required"`
	Quantity       int64   `json:"quantity" validate:"required"`
	SubTotal       float64 `json:"sub_total" validate:"required"`
}

type UpdateRecipeSchema struct {
	RegistrationId int64   `json:"registration_id,omitempty"`
	MedicineId     int64   `json:"medicine_id,omitempty"`
	Quantity       int64   `json:"quantity,omitempty"`
	SubTotal       float64 `json:"sub_total,omitempty"`
}

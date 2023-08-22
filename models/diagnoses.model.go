package models

import "github.com/go-playground/validator/v10"

type Diagnoses struct {
	ID            int64  `gorm:"size:20; primary_key" json:"id,omitempty"`
	DiagnosesName string `gorm:"size:255; not null" json:"diagnoses_name,omitempty"`
	DiagnosesCode string `gorm:"size:50; not null" json:"diagnoses_code,omitempty"`
}

func ValidateStructDiagnosis[T any](payload T) []*ErrorResponse {
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

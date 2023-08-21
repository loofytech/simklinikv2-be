package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Payment struct {
	ID            int64     `gorm:"size:20; primary_key" json:"id,omitempty"`
	PaymentName   string    `gorm:"size:50; not null" json:"payment_name,omitempty"`
	PaymentStatus *bool     `gorm:"default:1; not null" json:"payment_status,omitempty"`
	PaymentSlug   string    `gorm:"size:50; not null; unique" json:"payment_slug,omitempty"`
	CreatedAt     time.Time `gorm:"" json:"created_at,omitempty"`
	UpdatedAt     time.Time `gorm:"" json:"updated_at,omitempty"`
}

func ValidateStructPayment[T any](payload T) []*ErrorResponse {
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

type CreatePaymentSchema struct {
	PaymentName   string `json:"payment_name" validate:"required"`
	PaymentStatus *bool  `json:"payment_status,omitempty"`
	PaymentSlug   string `json:"payment_slug,omitempty"`
}

type UpdatePaymentSchema struct {
	PaymentName   string `json:"payment_name,omitempty"`
	PaymentStatus *bool  `json:"payment_status,omitempty"`
	PaymentSlug   string `json:"payment_slug,omitempty"`
}

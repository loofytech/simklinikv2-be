package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type InsuranceProduct struct {
	ID                          int64          `gorm:"size:20; primary_key" json:"id,omitempty"`
	InsuranceProductName        string         `gorm:"size:50; not null" json:"insurance_product_name,omitempty"`
	InsuranceProductAdminFee    float64        `gorm:"size:255; null" json:"insurance_product_admin_fee,omitempty"`
	InsuranceProductMaxAdminFee float64        `gorm:"size:50; not null" json:"insurance_product_max_admin_fee,omitempty"`
	InsuranceProductStamp       float64        `gorm:"size:50; not null" json:"insurance_product_stamp,omitempty"`
	CreatedAt                   time.Time      `gorm:"" json:"created_at,omitempty"`
	UpdatedAt                   time.Time      `gorm:"" json:"updated_at,omitempty"`
	RelationAgencyId            int64          `gorm:"size:20; foreign_key" json:"relation_agency_id,omitempty"`
	RelationAgency              RelationAgency `gorm:"references:id"`
}

func ValidateStructInsuranceProduct[T any](payload T) []*ErrorResponse {
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

type CreateInsuranceProductSchema struct {
	InsuranceProductName        string  `json:"insurance_product_name" validate:"required"`
	InsuranceProductAdminFee    float64 `json:"insurance_product_admin_fee,omitempty" validate:"required" `
	InsuranceProductMaxAdminFee float64 `json:"insurance_product_max_admin_fee,omitempty" validate:"required"`
	InsuranceProductStamp       float64 `json:"insurance_product_stamp" validate:"required"`
	RelationAgencyId            int64   `json:"relation_agency_id" validate:"required"`
}

type UpdateInsuranceProductSchema struct {
	InsuranceProductName        string  `json:"insurance_product_name,omitempty"`
	InsuranceProductAdminFee    float64 `json:"insurance_product_admin_fee,omitempty"`
	InsuranceProductMaxAdminFee float64 `json:"insurance_product_max_admin_fee,omitempty"`
	InsuranceProductStamp       float64 `json:"insurance_product_stamp,omitempty"`
	RelationAgencyId            int64   `json:"relation_agency_id,omitempty"`
}

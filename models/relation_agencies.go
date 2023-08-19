package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type RelationAgency struct {
	ID                    int64     `gorm:"size:20; primary_key" json:"id,omitempty"`
	RelationAgencyName    string    `gorm:"size:50; not null" json:"relation_agency_name,omitempty"`
	RelationAgencyAddress string    `gorm:"size:255; null" json:"relation_agency_address,omitempty"`
	RelationAgencyPhone   string    `gorm:"size:50; not null" json:"relation_agency_phone,omitempty"`
	RelationAgencyEmail   string    `gorm:"size:50; not null" json:"relation_agency_email,omitempty"`
	RelationAgencyWebsite string    `gorm:"size:50; null" json:"relation_agency_website,omitempty"`
	CreatedAt             time.Time `gorm:"" json:"created_at,omitempty"`
	UpdatedAt             time.Time `gorm:"" json:"updated_at,omitempty"`
}

func ValidateStructRelationAgency[T any](payload T) []*ErrorResponse {
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

type CreateRelationAgencySchema struct {
	RelationAgencyName    string `json:"relation_agency_name" validate:"required"`
	RelationAgencyAddress string `json:"relation_agency_address,omitempty"`
	RelationAgencyPhone   string `json:"relation_agency_phone,omitempty"`
	RelationAgencyEmail   string `json:"relation_agency_email" validate:"required"`
	RelationAgencyWebsite string `json:"relation_agency_website,omitempty"`
}

type UpdateRelationAgencySchema struct {
	RelationAgencyName    string `json:"relation_agency_name,omitempty"`
	RelationAgencyAddress string `json:"relation_agency_address,omitempty"`
	RelationAgencyPhone   string `json:"relation_agency_phone,omitempty"`
	RelationAgencyEmail   string `json:"relation_agency_email,omitempty"`
	RelationAgencyWebsite string `json:"relation_agency_website,omitempty"`
}

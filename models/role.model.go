package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Role struct {
	ID         int64     `gorm:"size:20; primary_key" json:"id,omitempty"`
	RoleName   string    `gorm:"size:50; not null" json:"role_name,omitempty"`
	RoleStatus *bool     `gorm:"default:1; not null; unique" json:"role_status,omitempty"`
	RoleSlug   string    `gorm:"size:50; not null; unique" json:"role_slug,omitempty"`
	CreatedAt  time.Time `gorm:"default:null" json:"created_at,omitempty"`
	UpdatedAt  time.Time `gorm:"default:null" json:"updated_at,omitempty"`
}

func ValidateStructRole[T any](payload T) []*ErrorResponse {
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

type CreateRoleSchema struct {
	RoleName   string `json:"role_name" validate:"required"`
	RoleStatus *bool  `json:"role_status" validate:"required"`
	RoleSlug   string `json:"role_slug,omitempty"`
}

type UpdateRoleSchema struct {
	RoleName   string `json:"role_name,omitempty"`
	RoleStatus int    `json:"role_status,omitempty"`
	RoleSlug   string `json:"role_slug,omitempty"`
}

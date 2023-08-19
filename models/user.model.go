package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        int64     `gorm:"size:20; primary_key" json:"id,omitempty"`
	Name      string    `gorm:"size:50; not null" json:"name,omitempty"`
	Username  string    `gorm:"size:50; not null; unique" json:"username,omitempty"`
	Email     string    `gorm:"size:50; not null; unique" json:"email,omitempty"`
	Password  string    `gorm:"size:255; not null" json:"password,omitempty"`
	RoleId    int64     `gorm:"size:20; foreign_key" json:"role_id,omitempty"`
	CreatedAt time.Time `gorm:"" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"" json:"updated_at,omitempty"`
	Role      Role      `gorm:"references:id"`
}

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
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

type CreateUserSchema struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	RoleId   int64  `json:"role_id,omitempty"`
}

type UpdateUserSchema struct {
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	RoleId   int64  `json:"role_id,omitempty"`
}

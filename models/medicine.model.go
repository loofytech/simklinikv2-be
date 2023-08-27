package models

import (
	"database/sql/driver"
	"time"

	"github.com/go-playground/validator/v10"
)

type medicineType string

const (
	Pill   medicineType = "Tablet"
	Tablet medicineType = "Pill"
)

func (en *medicineType) Scan(value interface{}) error {
	*en = medicineType(value.([]byte))
	return nil
}

func (en medicineType) Value() (driver.Value, error) {
	return string(en), nil
}

type Medicine struct {
	ID            int64        `gorm:"size:20; primary_key" json:"id,omitempty"`
	MedicineName  string       `gorm:"size:50; not null" json:"medicine_name,omitempty"`
	MedicineType  medicineType `gorm:"column:medicine_type;type:enum('Pill','Tablet')" json:"medicine_type"`
	MedicineHPP   float64      `gorm:"size:50; not null" json:"medicine_hpp,omitempty"`
	MedicineHNA   float64      `gorm:"size:50; not null" json:"medicine_hna,omitempty"`
	MedicineStock int64        `gorm:"size:50; not null" json:"medicine_stock,omitempty"`
	CreatedAt     time.Time    `gorm:"" json:"created_at,omitempty"`
	UpdatedAt     time.Time    `gorm:"" json:"updated_at,omitempty"`
}

func ValidateStructMedicine[T any](payload T) []*ErrorResponse {
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

type CreateMedicineSchema struct {
	MedicineName  string       `json:"medicine_name" validate:"required"`
	MedicineType  medicineType `json:"medicine_type" validate:"required"`
	MedicineHPP   float64      `json:"medicine_hpp" validate:"required"`
	MedicineHNA   float64      `json:"medicine_hna" validate:"required"`
	MedicineStock int64        `json:"medicine_stock" validate:"required"`
}

type UpdateMedicineSchema struct {
	MedicineName  string       `json:"medicine_name,omitempty"`
	MedicineType  medicineType `json:"medicine_type,omitempty"`
	MedicineHPP   float64      `json:"medicine_hpp,omitempty"`
	MedicineHNA   float64      `json:"medicine_hna,omitempty"`
	MedicineStock int64        `json:"medicine_stock,omitempty"`
}

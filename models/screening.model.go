package models

import (
	"database/sql/driver"
	"time"

	"github.com/go-playground/validator/v10"
)

type enum string

const (
	Positif        enum = "0"
	Negatif        enum = "1"
	TidakDiketahui enum = "2"
)

func (en *enum) Scan(value interface{}) error {
	*en = enum(value.([]byte))
	return nil
}

func (en enum) Value() (driver.Value, error) {
	return string(en), nil
}

type Screening struct {
	ID                     int64         `gorm:"size:20; primary_key" json:"id,omitempty"`
	IdScreening            string        `gorm:"size:20; not null; unique" json:"id_screening"`
	BodyWeight             float64       `gorm:"size:50; not null" json:"body_weight,omitempty"`
	BodyHeight             float64       `gorm:"size:50; not null" json:"body_height,omitempty"`
	BodyTemperature        float64       `gorm:"size:50; not null" json:"body_temperature,omitempty"`
	BodyBreath             float64       `gorm:"size:50; not null" json:"body_breath,omitempty"`
	BodyPulse              float64       `gorm:"size:50; not null" json:"body_pulse,omitempty"`
	BodyBloodPressureMM    float64       `gorm:"size:50; not null" json:"body_blood_pressure_mm,omitempty"`
	BodyBloodPressureHG    float64       `gorm:"size:50; not null" json:"body_blood_pressure_hg,omitempty"`
	BodyIMT                float64       `gorm:"size:50; not null" json:"body_imt,omitempty"`
	BodyOxygenSaturation   float64       `gorm:"size:50; not null" json:"body_oxygen_saturation,omitempty"`
	BodyDiabetes           enum          `gorm:"column:body_diabetes;type:enum('0','1','2')" json:"body_diabetes"`
	BodyHaemopilia         enum          `gorm:"column:body_haemopilia;type:enum('0','1','2')" json:"body_haemopilia"`
	BodyHeartDisease       enum          `gorm:"column:body_heart_desease;type:enum('0','1','2')" json:"body_heart_desease"`
	AbdominalCircumference float64       `gorm:"size:50; not null" json:"abdominal_circumference,omitempty"`
	HistoryOtherDesease    string        `gorm:"size:1000; not null" json:"history_other_desease,omitempty"`
	HistoryTreatment       string        `gorm:"size:1000; not null" json:"history_treatment,omitempty"`
	AllergyMedicine        string        `gorm:"size:1000; not null" json:"allergy_medicine,omitempty"`
	AllergyFood            string        `gorm:"size:1000; not null" json:"allergy_food,omitempty"`
	IsSubmit               string        `gorm:"size:10; not null" json:"is_submit,omitempty"`
	ServiceActionId        int64         `gorm:"size:20; foreign_key" json:"service_action_id,omitempty"`
	ServiceAction          ServiceAction `gorm:"references:id" json:"service_action"`
	PatientId              int64         `gorm:"size:20; foreign_key" json:"patient_id,omitempty"`
	Patient                Patient       `gorm:"references:id" json:"patient"`
	CreatedAt              time.Time     `gorm:"" json:"created_at,omitempty"`
	UpdatedAt              time.Time     `gorm:"" json:"updated_at,omitempty"`
}

func ValidateStructScreening[T any](payload T) []*ErrorResponse {
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

type CreateScreeningSchema struct {
	IdScreening            string  `json:"id_screening" validate:"required"`
	BodyWeight             float64 `json:"body_weight" validate:"required"`
	BodyHeight             float64 `json:"body_height" validate:"required"`
	BodyTemperature        float64 `json:"body_temperature" validate:"required"`
	BodyBreath             float64 `json:"body_breath" validate:"required"`
	BodyPulse              float64 `json:"body_pulse" validate:"required"`
	BodyBloodPressureMM    float64 `json:"body_blood_pressure_mm" validate:"required"`
	BodyBloodPressureHG    float64 `json:"body_blood_pressure_hg" validate:"required"`
	BodyIMT                float64 `json:"body_imt" validate:"required"`
	BodyOxygenSaturation   float64 `json:"body_oxygen_saturation" validate:"required"`
	BodyDiabetes           enum    `json:"body_diabetes" validate:"required"`
	BodyHaemopilia         enum    `json:"body_haemopilia" validate:"required"`
	BodyHeartDisease       enum    `json:"body_heart_desease" validate:"required"`
	AbdominalCircumference float64 `json:"abdominal_circumference" validate:"required"`
	HistoryOtherDesease    string  `json:"history_other_desease" validate:"required"`
	HistoryTreatment       string  `json:"history_treatment" validate:"required"`
	AllergyMedicine        string  `json:"allergy_medicine" validate:"required"`
	AllergyFood            string  `json:"allergy_food" validate:"required"`
	IsSubmit               bool    `json:"is_submit" validate:"required"`
	ServiceActionId        int64   `json:"service_action_id" validate:"required"`
}

type UpdateScreeningSchema struct {
	IdScreening            string  `json:"id_screening,omitempty"`
	BodyWeight             float64 `json:"body_weight,omitempty"`
	BodyHeight             float64 `json:"body_height,omitempty"`
	BodyTemperature        float64 `json:"body_temperature,omitempty"`
	BodyBreath             float64 `json:"body_breath,omitempty"`
	BodyPulse              float64 `json:"body_pulse,omitempty"`
	BodyBloodPressureMM    float64 `json:"body_blood_pressure_mm,omitempty"`
	BodyBloodPressureHG    float64 `json:"body_blood_pressure_hg,omitempty"`
	BodyIMT                float64 `json:"body_imt,omitempty"`
	BodyOxygenSaturation   float64 `json:"body_oxygen_saturation,omitempty"`
	BodyDiabetes           enum    `json:"body_diabetes,omitempty"`
	BodyHaemopilia         enum    `json:"body_haemopilia,omitempty"`
	BodyHeartDisease       enum    `json:"body_heart_desease,omitempty"`
	AbdominalCircumference float64 `json:"abdominal_circumference,omitempty"`
	HistoryOtherDesease    string  `json:"history_other_desease,omitempty"`
	HistoryTreatment       string  `json:"history_treatment,omitempty"`
	AllergyMedicine        string  `json:"allergy_medicine,omitempty"`
	AllergyFood            string  `json:"allergy_food,omitempty"`
	IsSubmit               bool    `json:"is_submit,omitempty"`
	ServiceActionId        int64   `json:"service_action_id,omitempty"`
	PatientId              int64   `json:"patient_id,omitempty"`
}

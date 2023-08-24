package models

import (
	"database/sql/driver"
	"time"

	"github.com/go-playground/validator/v10"
)

type patientGender string
type patientBloodType string

const (
	Lakilaki  patientGender = "L"
	Perempuan patientGender = "P"
)

const (
	TypeA  patientBloodType = "A"
	TypeB  patientBloodType = "B"
	TypeO  patientBloodType = "O"
	TypeAB patientBloodType = "AB"
)

func (pg *patientGender) Scan(value interface{}) error {
	*pg = patientGender(value.([]byte))
	return nil
}

func (pg patientGender) Value() (driver.Value, error) {
	return string(pg), nil
}

func (bt *patientBloodType) Scan(value interface{}) error {
	*bt = patientBloodType(value.([]byte))
	return nil
}

func (bt patientBloodType) Value() (driver.Value, error) {
	return string(bt), nil
}

type Patient struct {
	ID               int64            `gorm:"size:20; primary_key" json:"id,omitempty"`
	MedicalRecord    string           `gorm:"size:20; not null; unique" json:"medical_record"`
	PatientName      string           `gorm:"size:50; not null" json:"patient_name,omitempty"`
	PatientAddress   string           `gorm:"size:255; null" json:"patient_address,omitempty"`
	PatientPhone     string           `gorm:"size:50; not null" json:"patient_phone,omitempty"`
	PatientNik       string           `gorm:"size:50; not null" json:"patient_nik,omitempty"`
	BirthPlace       string           `gorm:"size:50; null" json:"birth_place,omitempty"`
	BirthDate        string           `gorm:"size:50; not null" json:"birth_date"`
	Province         string           `gorm:"size:50; not null" json:"province,omitempty"`
	Regency          string           `gorm:"size:50; not null" json:"regency,omitempty"`
	District         string           `gorm:"size:50; not null" json:"district,omitempty"`
	SubDistrict      string           `gorm:"size:50; not null" json:"sub_district,omitempty"`
	PatientGender    patientGender    `gorm:"column:patient_gender;type:enum('L','P')" json:"patient_gender"`
	PatientBloodType patientBloodType `gorm:"column:patient_blood_type;type:enum('A','B','O','AB')" json:"patient_blood_type"`
	CreatedAt        time.Time        `gorm:"" json:"created_at,omitempty"`
	UpdatedAt        time.Time        `gorm:"" json:"updated_at,omitempty"`
	ReligionId       int64            `gorm:"size:20; foreign_key" json:"religion_id,omitempty"`
	EthnicId         int64            `gorm:"size:20; foreign_key" json:"ethnic_id,omitempty"`
	JobId            int64            `gorm:"size:20; foreign_key" json:"job_id,omitempty"`
	EducationId      int64            `gorm:"size:20; foreign_key" json:"education_id,omitempty"`
	MaritalStatusId  int64            `gorm:"size:20; foreign_key" json:"marital_status_id,omitempty"`
	Religion         Religion         `gorm:"references:id" json:"religion"`
	Ethnic           Ethnic           `gorm:"references:id" json:"ethnic"`
	Job              Job              `gorm:"references:id" json:"job"`
	MaritalStatus    MaritalStatus    `gorm:"references:id" json:"marital_status"`
	Education        Education        `gorm:"references:id" json:"education"`
}

func ValidateStructPatient[T any](payload T) []*ErrorResponse {
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

type CreatePatientSchema struct {
	MedicalRecord    string           `json:"medical_record"`
	PatientName      string           `json:"patient_name" validate:"required"`
	PatientAddress   string           `json:"patient_address" validate:"required"`
	PatientPhone     string           `json:"patient_phone" validate:"required"`
	PatientNik       string           `json:"patient_nik" validate:"required"`
	BirthPlace       string           `json:"birth_place" validate:"required"`
	BirthDate        string           `json:"birth_date" validate:"required"`
	Province         string           `json:"province" validate:"required"`
	Regency          string           `json:"regency" validate:"required"`
	District         string           `json:"district" validate:"required"`
	SubDistrict      string           `json:"sub_district" validate:"required"`
	PatientGender    patientGender    `json:"patient_gender" validate:"required"`
	PatientBloodType patientBloodType `json:"patient_blood_type" validate:"required"`
	ReligionId       int64            `json:"religion_id" validate:"required"`
	EthnicId         int64            `json:"ethnic_id" validate:"required"`
	JobId            int64            `json:"job_id" validate:"required"`
	EducationId      int64            `json:"education_id" validate:"required"`
	MaritalStatusId  int64            `json:"marital_status_id" validate:"required"`
}

type UpdatePatientSchema struct {
	MedicalRecord    string           `json:"medical_record,omitempty"`
	PatientName      string           `json:"patient_name" validate:"required"`
	PatientAddress   string           `json:"patient_address,omitempty"`
	PatientPhone     string           `json:"patient_phone,omitempty"`
	PatientNik       string           `json:"patient_nik,omitempty"`
	BirthPlace       string           `json:"birth_place,omitempty"`
	BirthDate        string           `json:"birth_date,omitempty"`
	Province         string           `json:"province,omitempty"`
	Regency          string           `json:"regency,omitempty"`
	District         string           `json:"district,omitempty"`
	SubDistrict      string           `json:"sub_district,omitempty"`
	PatientGender    patientGender    `json:"patient_gender,omitempty"`
	PatientBloodType patientBloodType `json:"patient_blood_type,omitempty"`
	ReligionId       int64            `json:"religion_id,omitempty"`
	EthnicId         int64            `json:"ethnic_id,omitempty"`
	JobId            int64            `json:"job_id,omitempty"`
	EducationId      int64            `json:"education_id,omitempty"`
	MaritalStatusId  int64            `json:"marital_status_id,omitempty"`
}

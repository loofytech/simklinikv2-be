package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Inspection struct {
	ID                  int64         `gorm:"size:20; primary_key" json:"id,omitempty"`
	Anamnesis           string        `gorm:"size:50; not null" json:"anamnesis,omitempty"`
	Objective           string        `gorm:"size:50; not null" json:"objective,omitempty"`
	KU                  string        `gorm:"size:50; not null" json:"k_u,omitempty"`
	Thoraks             string        `gorm:"size:50; not null" json:"thoraks,omitempty"`
	Therapy             string        `gorm:"size:50; not null" json:"therapy,omitempty"`
	Educations          string        `gorm:"size:50; not null" json:"educations,omitempty"`
	Instructions        string        `gorm:"size:50; not null" json:"instructions,omitempty"`
	Abd                 string        `gorm:"size:50; not null" json:"abd,omitempty"`
	Extremity           string        `gorm:"size:50; not null" json:"extremity,omitempty"`
	WorkingDiagnosis    string        `gorm:"size:50; not null" json:"working_diagnosis,omitempty"`
	PhysicalExamination string        `gorm:"size:50; not null" json:"physical_examination,omitempty"`
	Explanation         string        `gorm:"size:1000; not null" json:"explanation,omitempty"`
	AttachmentBefore    string        `gorm:"size:50; not null" json:"attachment_before,omitempty"`
	AttachmentAfter     string        `gorm:"size:50; not null" json:"attachment_after,omitempty"`
	ServiceActionId     int64         `gorm:"size:20; foreign_key" json:"service_action_id,omitempty"`
	DiagnoseId          string        `gorm:"size:50; is null" json:"diagnoses_id,omitempty"`
	ServiceAction       ServiceAction `gorm:"references:id" json:"service_action"`
	CreatedAt           time.Time     `gorm:"" json:"created_at,omitempty"`
	UpdatedAt           time.Time     `gorm:"" json:"updated_at,omitempty"`
}

func ValidateStructInspection[T any](payload T) []*ErrorResponse {
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

type CreateInspectionSchema struct {
	Anamnesis           string `json:"anamnesis" validate:"required"`
	Objective           string `json:"objective" validate:"required"`
	KU                  string `json:"k_u" validate:"required"`
	Thoraks             string `json:"thoraks" validate:"required"`
	Therapy             string `json:"therapy" validate:"required"`
	Education           string `json:"educations" validate:"required"`
	Instructions        string `json:"instructions" validate:"required"`
	Abd                 string `json:"abd" validate:"required"`
	Extremity           string `json:"extremity" validate:"required"`
	WorkingDiagnosis    string `json:"working_diagnosis" validate:"required"`
	DiagnoseId          string `json:"diagnoses_id" validate:"required"`
	PhysicalExamination string `json:"physical_examination" validate:"required"`
	Explanation         string `json:"explanation" validate:"required"`
	AttachmentBefore    string `json:"attachment_before" validate:"required"`
	AttachmentAfter     string `json:"attachment_after" validate:"required"`
	ServiceActionId     int64  `json:"service_action_id" validate:"required"`
}

type UpdateInspectionSchema struct {
	Anamnesis           string `json:"anamnesis,omitempty"`
	Objective           string `json:"objective,omitempty"`
	KU                  string `json:"k_u,omitempty"`
	Thoraks             string `json:"thoraks,omitempty"`
	Therapy             string `json:"therapy,omitempty"`
	Education           string `json:"educations,omitempty"`
	Instructions        string `json:"instructions,omitempty"`
	Abd                 string `json:"abd,omitempty"`
	Extremity           string `json:"extremity,omitempty"`
	WorkingDiagnosis    string `json:"working_diagnosis,omitempty"`
	DiagnoseId          string `json:"diagnoses_id,omitempty"`
	PhysicalExamination string `json:"physical_examination,omitempty"`
	Explanation         string `json:"explanation,omitempty"`
	AttachmentBefore    string `json:"attachment_before,omitempty"`
	AttachmentAfter     string `json:"attachment_after,omitempty"`
	ServiceActionId     int64  `json:"service_action_id" validate:"required"`
}

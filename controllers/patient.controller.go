package controllers

import (
	"fmt"
	"sim-klinikv2/config"
	"sim-klinikv2/models"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreatePatientHandler(c *fiber.Ctx) error {
	var patient *models.Patient
	var payload *models.CreatePatientSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	var counting int64
	config.DB.Find(&patient).Count(&counting)

	mr := strings.Join([]string{"A", fmt.Sprintf("%06d", counting+1)}, "")

	fmt.Print(payload.MedicalRecord == "")

	now := time.Now()
	newPatient := models.Patient{
		MedicalRecord:   mr,
		PatientName:     payload.PatientName,
		PatientAddress:  payload.PatientAddress,
		PatientPhone:    payload.PatientPhone,
		PatientNik:      payload.PatientNik,
		BirthPlace:      payload.BirthPlace,
		BirthDate:       payload.BirthDate,
		Province:        payload.Province,
		Regency:         payload.Regency,
		District:        payload.District,
		SubDistrict:     payload.SubDistrict,
		PatientGender:   payload.PatientGender,
		JobId:           payload.JobId,
		EthnicId:        payload.EthnicId,
		ReligionId:      payload.ReligionId,
		EducationId:     payload.EducationId,
		MaritalStatusId: payload.MaritalStatusId,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	result := config.DB.Create(&newPatient)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Patient already exist"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   fiber.Map{"Patient": newPatient},
	})
}

func FindPatient(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var patient []models.Patient
	results := config.DB.Limit(intLimit).Offset(offset).Preload("Job").Preload("Ethnic").Preload("Religion").Preload("Education").Preload("MaritalStatus").Find(&patient)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(patient), "patient": patient})
}

func UpdatePatient(c *fiber.Ctx) error {
	patientId := c.Params("patientId")

	var payload *models.UpdatePatientSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var patient models.Patient
	result := config.DB.First(&patient, "id = ?", patientId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Patient with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})

	if payload.PatientName != "" {
		updates["patient_name"] = payload.PatientName
	}

	if payload.PatientAddress != "" {
		updates["patient_address"] = payload.PatientAddress
	}

	if payload.PatientPhone != "" {
		updates["patient_phone"] = payload.PatientPhone
	}

	if payload.PatientNik != "" {
		updates["patient_nik"] = payload.PatientNik
	}

	if payload.BirthPlace != "" {
		updates["birth_place"] = payload.BirthPlace
	}

	if payload.BirthDate != "" {
		updates["birth_date"] = payload.BirthDate
	}

	if payload.Province != "" {
		updates["province"] = payload.Province
	}

	if payload.Regency != "" {
		updates["regency"] = payload.Regency
	}

	if payload.District != "" {
		updates["district"] = payload.District
	}

	if payload.PatientGender != "" {
		updates["patient_gender"] = payload.PatientGender
	}

	if payload.PatientBloodType != "" {
		updates["patient_blood_type"] = payload.PatientBloodType
	}

	if payload.SubDistrict != "" {
		updates["sub_district"] = payload.SubDistrict
	}

	updates["updated_at"] = time.Now()

	config.DB.Model(&patient).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"data": patient}})
}

func FindPatientById(c *fiber.Ctx) error {
	patientId := c.Params("patientId")

	var patient models.Patient
	result := config.DB.Preload("Job").Preload("Ethnic").Preload("Religion").Preload("Education").Preload("MaritalStatus").First(&patient, "id = ?", patientId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Patient with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"data": patient}})
}

func PatientDelete(c *fiber.Ctx) error {
	patientId := c.Params("patientId")

	result := config.DB.Delete(&models.Patient{}, "id = ?", patientId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Patient with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
}

func FindPatientByMR(c *fiber.Ctx) error {
	medicalRecord := c.Params("medicalRecord")

	var patient models.Patient
	result := config.DB.Preload("Job").Preload("Ethnic").Preload("Religion").Preload("Education").Preload("MaritalStatus").First(&patient, "medical_record = ?", medicalRecord)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(400).JSON(fiber.Map{"status": "fail", "message": "No Patient with that Medical Record exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": patient})
}

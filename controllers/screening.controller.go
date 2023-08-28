package controllers

import (
	"sim-klinikv2/config"
	"sim-klinikv2/models"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UpdateScreening(c *fiber.Ctx) error {
	var patient models.Patient
	var mmx *models.UpdatePatientSchema

	pt := struct {
		ID               *int64 `json:"patient_id"`
		PatientBloodType string `json:"patient_blood_type"`
	}{}

	screeningId := c.Params("screeningId")

	var payload *models.UpdateScreeningSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var screening models.Screening
	result := config.DB.First(&screening, "id = ?", screeningId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Screening with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})

	if payload.BodyWeight != 0 {
		updates["body_weight"] = payload.BodyWeight
	}

	if payload.BodyHeight != 0 {
		updates["body_height"] = payload.BodyHeight
	}

	if payload.BodyTemperature != 0 {
		updates["body_temperature"] = payload.BodyTemperature
	}

	if payload.BodyBreath != 0 {
		updates["body_breath"] = payload.BodyBreath
	}

	if payload.BodyPulse != 0 {
		updates["body_pulse"] = payload.BodyPulse
	}

	if payload.BodyBloodPressureMM != 0 {
		updates["body_blood_pressure_mm"] = payload.BodyBloodPressureMM
	}

	if payload.BodyBloodPressureHG != 0 {
		updates["body_blood_pressure_hg"] = payload.BodyBloodPressureHG
	}

	if payload.BodyIMT != 0 {
		updates["body_imt"] = payload.BodyOxygenSaturation
	}

	if payload.BodyDiabetes != "" {
		updates["body_diabetes"] = payload.BodyDiabetes
	}

	if payload.BodyHaemopilia != "" {
		updates["body_haemopilia"] = payload.BodyHaemopilia
	}

	if payload.BodyHeartDisease != "" {
		updates["body_heart_desease"] = payload.BodyHeartDisease
	}

	if payload.AbdominalCircumference != 0 {
		updates["abdominal_circumference"] = payload.AbdominalCircumference
	}

	if payload.HistoryOtherDesease != "0" {
		updates["history_other_desease"] = payload.HistoryOtherDesease
	}

	if payload.HistoryTreatment != "" {
		updates["history_treatment"] = payload.HistoryTreatment
	}

	if payload.AllergyMedicine != "" {
		updates["allergy_medicine"] = payload.AllergyMedicine
	}

	if payload.AllergyFood != "" {
		updates["allergy_food"] = payload.AllergyFood
	}

	if payload.ServiceActionId != 0 {
		updates["service_action_id"] = payload.ServiceActionId
	}

	updates["is_submit"] = "true"
	updates["updated_at"] = time.Now()

	config.DB.Model(&screening).Updates(updates)

	if err := c.BodyParser(&pt); err != nil {
		return err
	}

	if pt.ID != nil {
		chkDB := config.DB.Where(&models.Patient{ID: *pt.ID}).First(&patient)
		if chkDB.RowsAffected == 0 {
			return c.Status(400).JSON(fiber.Map{"status": "bad request", "message": "Data Patient tidak ditemukan"})
		}

		if err := c.BodyParser(&mmx); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
		}

		result := chkDB
		if err := result.Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Patient with that Id exists"})
			}
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
		}

		updates := make(map[string]interface{})
		updates["patient_blood_type"] = mmx.PatientBloodType

		config.DB.Model(&patient).Updates(updates)
		if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Patient already exist"})
		} else if result.Error != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
		}
	}

	// create Inspection
	newInspection := models.Inspection{
		Anamnesis:           "",
		Objective:           "",
		KU:                  "",
		Thoraks:             "",
		Therapy:             "",
		Educations:          "",
		Instructions:        "",
		Abd:                 "",
		Extremity:           "",
		WorkingDiagnosis:    "",
		PhysicalExamination: "",
		Explanation:         "",
		AttachmentBefore:    "",
		AttachmentAfter:     "",
		DiagnoseId:          "",
		ServiceActionId:     payload.ServiceActionId,
	}

	resInsection := config.DB.Create(&newInspection)

	if resInsection.Error != nil && strings.Contains(resInsection.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Screening already exist"})
	} else if resInsection.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": resInsection.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": screening})
}

func FindScreening(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var screening []models.Screening

	results := config.DB.Limit(intLimit).Offset(offset).Preload("Patient").Find(&screening)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(screening), "data": screening})
}

func FindScreeningById(c *fiber.Ctx) error {
	screeningId := c.Params("screeningId")

	var screening models.Screening
	result := config.DB.Preload("Patient").First(&screening, "id = ?", screeningId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Screening with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": screening})
}

func ScreeningDelete(c *fiber.Ctx) error {
	screeningId := c.Params("screeningId")

	result := config.DB.Delete(&models.Screening{}, "id = ?", screeningId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Screening with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
}

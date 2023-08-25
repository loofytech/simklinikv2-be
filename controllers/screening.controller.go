package controllers

import (
	"sim-klinikv2/config"
	"sim-klinikv2/models"
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

	if payload.BodyBloodPressure != 0 {
		updates["body_blood_pressure"] = payload.BodyBloodPressure
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

	// if payload.ServiceActionId != nil {
	// 	updates["sub_district"] = payload.ServiceActionId
	// }

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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": screening})
}

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

func CreateMeasure(c *fiber.Ctx) error {
	var payload *models.CreateMeasureSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	now := time.Now()
	newMeasure := models.Measure{
		RegistrationId: payload.RegistrationId,
		ClinicRateId:   payload.ClinicRateId,
		NakesFirst:     payload.NakesFirst,
		SubTotal:       payload.SubTotal,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	result := config.DB.Create(&newMeasure)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Measure already exist"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   newMeasure,
	})
}

func FindMeasure(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var measure []models.Measure
	results := config.DB.Limit(intLimit).Offset(offset).Find(&measure)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(measure), "data": measure})
}

func UpdateMeasure(c *fiber.Ctx) error {
	measureId := c.Params("measureId")

	var payload *models.UpdateMeasureSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var measure models.Measure
	result := config.DB.First(&measure, "id = ?", measureId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Measure with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	if payload.RegistrationId != 0 {
		updates["registration_id"] = payload.RegistrationId
	}

	if payload.ClinicRateId != 0 {
		updates["clinic_rate_id"] = payload.ClinicRateId
	}

	if payload.NakesFirst != "" {
		updates["nakes_first"] = payload.NakesFirst
	}

	if payload.NakesSecond != "" {
		updates["nakes_second"] = payload.NakesSecond
	}

	if payload.SubTotal != 0 {
		updates["sub_total"] = payload.SubTotal
	}

	updates["updated_at"] = time.Now()

	config.DB.Model(&measure).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": measure})
}

func FindMeasureById(c *fiber.Ctx) error {
	measureId := c.Params("measureId")

	var measure models.Measure
	result := config.DB.First(&measure, "id = ?", measureId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Measure with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": measure})
}

func MeasureDelete(c *fiber.Ctx) error {
	measureId := c.Params("measureId")

	result := config.DB.Delete(&models.Measure{}, "id = ?", measureId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Measure with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
}

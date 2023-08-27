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

func CreateClinicRate(c *fiber.Ctx) error {
	var payload *models.CreateClinicRateSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	now := time.Now()
	newClinicRate := models.ClinicRate{
		ClinicRateName:   payload.ClinicRateName,
		ClinicRatePrice:  payload.ClinicRatePrice,
		ClinicRateStatus: payload.ClinicRateStatus,
		UnitId:           payload.UnitId,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	result := config.DB.Create(&newClinicRate)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "ClinicRate already exist"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   newClinicRate,
	})
}

func FindClinicRate(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var clinicRate []models.ClinicRate
	results := config.DB.Limit(intLimit).Offset(offset).Find(&clinicRate)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(clinicRate), "data": clinicRate})
}

func UpdateClinicRate(c *fiber.Ctx) error {
	clinicRateId := c.Params("clinicRateId")

	var payload *models.UpdateClinicRateSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var clinicRate models.ClinicRate
	result := config.DB.First(&clinicRate, "id = ?", clinicRateId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No ClinicRate with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	if payload.ClinicRateName != "" {
		updates["rate_name"] = payload.ClinicRateName
	}

	if payload.ClinicRatePrice != 0 {
		updates["rate_price"] = payload.ClinicRatePrice
	}

	if payload.ClinicRateStatus != nil {
		updates["rate_status"] = payload.ClinicRateStatus
	}

	if payload.UnitId != 0 {
		updates["unit_id"] = payload.UnitId
	}

	updates["updated_at"] = time.Now()

	config.DB.Model(&clinicRate).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": clinicRate})
}

func FindClinicRateById(c *fiber.Ctx) error {
	clinicRateId := c.Params("clinicRateId")

	var clinicRate models.ClinicRate
	result := config.DB.First(&clinicRate, "id = ?", clinicRateId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No ClinicRate with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"ClinicRate": clinicRate}})
}

func ClinicRateDelete(c *fiber.Ctx) error {
	clinicRateId := c.Params("clinicRateId")

	result := config.DB.Delete(&models.ClinicRate{}, "id = ?", clinicRateId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No ClinicRate with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
}

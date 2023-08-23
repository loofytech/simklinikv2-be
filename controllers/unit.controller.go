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

func CreateUnitHandler(c *fiber.Ctx) error {
	var payload *models.CreateUnitSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	now := time.Now()
	newUnit := models.Unit{
		UnitName:   payload.UnitName,
		UnitStatus: payload.UnitStatus,
		UnitSlug:   payload.UnitSlug,
		ServiceId:  payload.ServiceId,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	result := config.DB.Create(&newUnit)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Unit already exist"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   newUnit,
	})
}

func FindUnit(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var unit []models.Unit
	results := config.DB.Limit(intLimit).Offset(offset).Preload("Service").Find(&unit)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(unit), "data": unit})
}

func UpdateUnit(c *fiber.Ctx) error {
	unitId := c.Params("unitId")

	var payload *models.UpdateUnitSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var unit models.Unit
	result := config.DB.First(&unit, "id = ?", unitId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Unit with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	if payload.UnitName != "" {
		updates["unit_name"] = payload.UnitName
	}

	if payload.UnitStatus != nil {
		updates["unit_status"] = payload.UnitStatus
	}

	updates["updated_at"] = time.Now()

	config.DB.Model(&unit).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"unit": unit}})
}

func FindUnitById(c *fiber.Ctx) error {
	unitId := c.Params("unitId")

	var unit models.Unit
	result := config.DB.Preload("Service").First(&unit, "id = ?", unitId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Unit with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"data": unit}})
}

func UnitDelete(c *fiber.Ctx) error {
	unitId := c.Params("unitId")

	result := config.DB.Delete(&models.Unit{}, "id = ?", unitId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Unit with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
}

func FindUnitByServiceId(c *fiber.Ctx) error {
	serviceId := c.Params("serviceId")

	var service models.Unit
	result := config.DB.First(&service, "service_id = ?", serviceId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Unit with that Srevice Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"data": service}})
}

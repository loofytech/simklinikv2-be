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

func CreateMaritalStatusHandler(c *fiber.Ctx) error {
	var payload *models.CreateMaritalStatusSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	now := time.Now()
	newMaritalStatus := models.MaritalStatus{
		MaritalName: payload.MaritalName,
		MaritalSlug: payload.MaritalSlug,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	result := config.DB.Create(&newMaritalStatus)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "MaritalStatus already exist"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   newMaritalStatus,
	})
}

func FindMaritalStatus(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var maritalStatus []models.MaritalStatus
	results := config.DB.Limit(intLimit).Offset(offset).Find(&maritalStatus)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(maritalStatus), "data": maritalStatus})
}

func UpdateMaritalStatus(c *fiber.Ctx) error {
	maritalStatusId := c.Params("maritalStatusId")

	var payload *models.UpdateMaritalStatusSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var maritalStatus models.MaritalStatus
	result := config.DB.First(&maritalStatus, "id = ?", maritalStatusId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No MaritalStatus with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	if payload.MaritalName != "" {
		updates["marital_name"] = payload.MaritalName
	}
	// if payload.MaritalActive != nil {
	// 	updates["marital_active"] = payload.MaritalActive
	// }

	updates["updated_at"] = time.Now()

	config.DB.Model(&maritalStatus).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": maritalStatus}})
}

func FindMaritalStatusById(c *fiber.Ctx) error {
	maritalStatusId := c.Params("maritalStatusId")

	var maritalStatus models.MaritalStatus
	result := config.DB.First(&maritalStatus, "id = ?", maritalStatusId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No MaritalStatus with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": maritalStatus}})
}

func MaritalStatusDelete(c *fiber.Ctx) error {
	maritalStatusId := c.Params("maritalStatusId")

	result := config.DB.Delete(&models.MaritalStatus{}, "id = ?", maritalStatusId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No MaritalStatus with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
}

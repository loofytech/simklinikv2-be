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

func CreateServiceActionHandler(c *fiber.Ctx) error {
	var payload *models.CreateServiceActionSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	now := time.Now()
	newService := models.ServiceAction{
		UnitId:         payload.UnitId,
		UserId:         payload.UserId,
		RegistrationId: payload.RegistrationId,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	result := config.DB.Create(&newService)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Service already exist"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   fiber.Map{"Service": newService},
	})
}

func FindServiceAction(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var serviceAction []models.ServiceAction
	results := config.DB.Limit(intLimit).Offset(offset).Preload("User").Preload("Unit").Preload("Registration").Find(&serviceAction)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(serviceAction), "user": serviceAction})
}

func UpdateServiceAction(c *fiber.Ctx) error {
	serviceActionId := c.Params("serviceActionId")

	var payload *models.UpdateServiceActionSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var service models.ServiceAction
	result := config.DB.First(&service, "id = ?", serviceActionId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Service with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	// if payload.ServiceName != "" {
	// 	updates["service_name"] = payload.ServiceName
	// }

	// if payload.ServiceActive != nil {
	// 	updates["service_active"] = payload.ServiceActive
	// }

	updates["updated_at"] = time.Now()

	config.DB.Model(&service).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": service}})
}

func FindServiceActionById(c *fiber.Ctx) error {
	serviceActionId := c.Params("serviceActionId")

	var service models.ServiceAction
	result := config.DB.Preload("User").Preload("Unit").Preload("Registration").First(&service, "id = ?", serviceActionId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Service with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": service}})
}

func ServiceActionDelete(c *fiber.Ctx) error {
	serviceActionId := c.Params("serviceActionId")

	result := config.DB.Delete(&models.Service{}, "id = ?", serviceActionId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Service with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
}

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

func CreateEthnicHandler(c *fiber.Ctx) error {
	var payload *models.CreateEthnicSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	now := time.Now()
	newEthnic := models.Ethnic{
		EthnicName:   payload.EthnicName,
		EthnicActive: payload.EthnicActive,
		EthnicSlug:   payload.EthnicSlug,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	result := config.DB.Create(&newEthnic)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Ethnic already exist"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   fiber.Map{"Ethnic": newEthnic},
	})
}

func FindEthnic(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var ethnic []models.Ethnic
	results := config.DB.Limit(intLimit).Offset(offset).Find(&ethnic)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(ethnic), "user": ethnic})
}

func UpdateEthnic(c *fiber.Ctx) error {
	ethnicId := c.Params("ethnicId")

	var payload *models.UpdateEthnicSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var ethnic models.Ethnic
	result := config.DB.First(&ethnic, "id = ?", ethnicId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Ethnic with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	if payload.EthnicName != "" {
		updates["ethnic_name"] = payload.EthnicName
	}
	if payload.EthnicActive != nil {
		updates["ethnic_active"] = payload.EthnicActive
	}

	updates["updated_at"] = time.Now()

	config.DB.Model(&ethnic).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": ethnic}})
}

func FindEthnicById(c *fiber.Ctx) error {
	ethnicId := c.Params("ethnicId")

	var ethnic models.Ethnic
	result := config.DB.First(&ethnic, "id = ?", ethnicId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Ethnic with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": ethnic}})
}

func EthnicDelete(c *fiber.Ctx) error {
	ethnicId := c.Params("ethnicId")

	result := config.DB.Delete(&models.Ethnic{}, "id = ?", ethnicId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Ethnic with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
}

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

func CreateUserUnitHandler(c *fiber.Ctx) error {
	var payload *models.CreateUserUnitSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	now := time.Now()
	newUserUnit := models.UserUnit{
		UserId:    payload.UserId,
		UnitId:    payload.UnitId,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := config.DB.Create(&newUserUnit)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "UserUnit already exist"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   fiber.Map{"UserUnit": newUserUnit},
	})
}

func FindUserUnit(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var userUnit []models.UserUnit
	results := config.DB.Limit(intLimit).Offset(offset).Find(&userUnit)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(userUnit), "data": userUnit})
}

func UpdateUserUnit(c *fiber.Ctx) error {
	userUnitId := c.Params("userUnitId")

	var payload *models.UpdateUserUnitSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var userUnit models.UserUnit
	result := config.DB.First(&userUnit, "id = ?", userUnitId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No UserUnit with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	// if payload.UserUnitName != "" {
	// 	updates["UserUnit_name"] = payload.UserUnitName
	// }

	// if payload.UserUnitStatus != nil {
	// 	updates["UserUnit_status"] = payload.UserUnitStatus
	// }

	updates["updated_at"] = time.Now()

	config.DB.Model(&userUnit).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"data": userUnit}})
}

func FindUserUnitById(c *fiber.Ctx) error {
	userUnitId := c.Params("userUnitId")

	var userUnit models.UserUnit
	result := config.DB.First(&userUnit, "id = ?", userUnitId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No UserUnit with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"data": userUnit}})
}

func UserUnitDelete(c *fiber.Ctx) error {
	userUnitId := c.Params("userUnitId")

	result := config.DB.Delete(&models.UserUnit{}, "id = ?", userUnitId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No UserUnit with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
}

package controllers

import (
	"sim-klinikv2/config"
	"sim-klinikv2/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// type APIDIagnoses struct {
// 	DiagnosesName string
// 	DiagnosesCode string
// }

func FindDiagnoses(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var diagnoses []models.Diagnoses
	results := config.DB.Limit(intLimit).Offset(offset).Find(&diagnoses)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(diagnoses), "data": diagnoses})
}

func FindDiagnosesByName(c *fiber.Ctx) error {
	diagnosesCode := c.Params("diagnosesCode")

	var diagnoses models.Diagnoses
	result := config.DB.Where("diagnoses_code LIKE ?", diagnosesCode).Find(&diagnoses)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Diagnoses with that Name exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"diagnoses": diagnoses}})
}

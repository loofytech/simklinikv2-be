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

func CreateMedicine(c *fiber.Ctx) error {
	var payload *models.CreateMedicineSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	now := time.Now()
	newMedicine := models.Medicine{
		MedicineName: payload.MedicineName,
		MedicineType: payload.MedicineType,
		MedicineHPP:  payload.MedicineHPP,
		MedicineHNA:  payload.MedicineHNA,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	result := config.DB.Create(&newMedicine)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Medicine already exist"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   newMedicine,
	})
}

func FindMedicine(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var medicine []models.Medicine
	results := config.DB.Limit(intLimit).Offset(offset).Find(&medicine)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(medicine), "data": medicine})
}

func UpdateMedicine(c *fiber.Ctx) error {
	medicineId := c.Params("medicineId")

	var payload *models.UpdateMedicineSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var medicine models.Medicine
	result := config.DB.First(&medicine, "id = ?", medicineId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Medicine with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	if payload.MedicineName != "" {
		updates["medicine_name"] = payload.MedicineName
	}

	if payload.MedicineType != "" {
		updates["medicine_type"] = payload.MedicineType
	}

	if payload.MedicineHPP != 0 {
		updates["medicine_hpp"] = payload.MedicineHPP
	}

	if payload.MedicineHNA != 0 {
		updates["medicine_hna"] = payload.MedicineHNA
	}

	if payload.MedicineStock != 0 {
		updates["medicine_stock"] = payload.MedicineStock
	}

	updates["updated_at"] = time.Now()

	config.DB.Model(&medicine).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": medicine})
}

func FindMedicineById(c *fiber.Ctx) error {
	medicineId := c.Params("medicineId")

	var medicine models.Medicine
	result := config.DB.First(&medicine, "id = ?", medicineId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Medicine with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"Medicine": medicine}})
}

func MedicineDelete(c *fiber.Ctx) error {
	medicineId := c.Params("medicineId")

	result := config.DB.Delete(&models.Medicine{}, "id = ?", medicineId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Medicine with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
}

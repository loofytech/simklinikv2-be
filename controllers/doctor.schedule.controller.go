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

func CreateDoctorScheduleHandler(c *fiber.Ctx) error {
	var payload *models.CreateDoctorScheduleSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	now := time.Now()
	newDoctorSchedule := models.DoctorSchedule{
		UserId:        payload.UserId,
		UnitId:        payload.UnitId,
		Day:           payload.Day,
		OpenPractice:  payload.OpenPractice,
		ClosePractice: payload.ClosePractice,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	result := config.DB.Create(&newDoctorSchedule)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "DoctorSchedule already exist"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   newDoctorSchedule,
	})
}

func FindDoctorSchedule(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var doctorSchedule []models.DoctorSchedule
	results := config.DB.Limit(intLimit).Offset(offset).Joins("User").Joins("Unit").Find(&doctorSchedule)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(doctorSchedule), "data": doctorSchedule})
}

func UpdateDoctorSchedule(c *fiber.Ctx) error {
	doctorScheduleId := c.Params("doctorScheduleId")

	var payload *models.UpdateDoctorScheduleSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var doctorSchedule models.DoctorSchedule
	result := config.DB.First(&doctorSchedule, "id = ?", doctorScheduleId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No DoctorSchedule with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	if payload.Day != "" {
		updates["day"] = payload.Day
	}

	if payload.OpenPractice != "" {
		updates["open_practice"] = payload.OpenPractice
	}

	if payload.ClosePractice != "" {
		updates["close_practice"] = payload.ClosePractice
	}

	updates["updated_at"] = time.Now()

	config.DB.Model(&doctorSchedule).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": doctorSchedule})
}

func FindDoctorScheduleById(c *fiber.Ctx) error {
	doctorScheduleId := c.Params("doctorScheduleId")

	var doctorSchedule models.DoctorSchedule
	result := config.DB.Preload("User").Preload("Unit").First(&doctorSchedule, "id = ?", doctorScheduleId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No DoctorSchedule with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": doctorSchedule})
}

func DoctorScheduleDelete(c *fiber.Ctx) error {
	doctorScheduleId := c.Params("doctorScheduleId")

	result := config.DB.Delete(&models.DoctorSchedule{}, "id = ?", doctorScheduleId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No DoctorSchedule with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
}

func FindScheduleByUnit(c *fiber.Ctx) error {
	unitId := c.Params("unitId")

	var schedule []models.DoctorSchedule
	result := config.DB.Joins("User").Find(&schedule, "unit_id = ?", unitId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Unit with that Srevice Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": schedule})
}

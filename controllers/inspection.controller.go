package controllers

import (
	"sim-klinikv2/config"
	"sim-klinikv2/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func FindInspection(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var inspection []models.Inspection
	results := config.DB.Limit(intLimit).Offset(offset).Find(&inspection)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(inspection), "data": inspection})
}

func UpdateInspection(c *fiber.Ctx) error {
	inspectionId := c.Params("inspectionId")

	var payload *models.UpdateInspectionSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var inspection models.Inspection
	result := config.DB.First(&inspection, "id = ?", inspectionId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Inspection with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	if payload.Anamnesis != "" {
		updates["anamnesis"] = payload.Anamnesis
	}

	if payload.Objective != "" {
		updates["objective"] = payload.Objective
	}

	if payload.KU != "" {
		updates["k_u"] = payload.KU
	}

	if payload.Thoraks != "" {
		updates["thoraks"] = payload.Thoraks
	}

	if payload.Therapy != "" {
		updates["therapy"] = payload.Therapy
	}

	if payload.Education != "" {
		updates["educations"] = payload.Education
	}

	if payload.Instructions != "" {
		updates["instructions"] = payload.Instructions
	}

	if payload.Abd != "" {
		updates["abd"] = payload.Abd
	}

	if payload.Extremity != "" {
		updates["extremity"] = payload.Extremity
	}

	if payload.WorkingDiagnosis != "" {
		updates["working_diagnosis"] = payload.WorkingDiagnosis
	}

	if payload.DiagnoseId != "" {
		updates["diagnosis"] = payload.DiagnoseId
	}

	if payload.PhysicalExamination != "" {
		updates["physical_examniation"] = payload.PhysicalExamination
	}

	if payload.Explanation != "" {
		updates["explanation"] = payload.Explanation
	}

	if payload.AttachmentBefore != "" {
		updates["attachment_before"] = payload.AttachmentBefore
	}

	if payload.AttachmentAfter != "" {
		updates["attachment_after"] = payload.AttachmentAfter
	}

	if payload.ServiceActionId != 0 {
		updates["service_action_id"] = payload.ServiceActionId
	}

	updates["updated_at"] = time.Now()

	config.DB.Model(&inspection).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"Inspection": inspection}})
}

func InspectionDelete(c *fiber.Ctx) error {
	InspectionId := c.Params("InspectionId")

	result := config.DB.Delete(&models.Inspection{}, "id = ?", InspectionId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Inspection with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
}

func FindInspectionById(c *fiber.Ctx) error {
	inspectionId := c.Params("inspectionId")

	var inspection models.Inspection
	result := config.DB.First(&inspection, "id = ?", inspectionId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Inspection with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": inspection})
}

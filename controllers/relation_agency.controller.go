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

func CreateRelationAgencyHandler(c *fiber.Ctx) error {
	var payload *models.CreateRelationAgencySchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	now := time.Now()
	newRelationAgency := models.RelationAgency{
		RelationAgencyName:    payload.RelationAgencyName,
		RelationAgencyAddress: payload.RelationAgencyAddress,
		RelationAgencyPhone:   payload.RelationAgencyPhone,
		RelationAgencyEmail:   payload.RelationAgencyEmail,
		RelationAgencyWebsite: payload.RelationAgencyWebsite,
		CreatedAt:             now,
		UpdatedAt:             now,
	}

	result := config.DB.Create(&newRelationAgency)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "RelationAgency already exist"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   newRelationAgency,
	})
}

func FindRelationAgency(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var relationAgency []models.RelationAgency
	results := config.DB.Limit(intLimit).Offset(offset).Find(&relationAgency)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(relationAgency), "data": relationAgency})
}

func UpdateRelationAgency(c *fiber.Ctx) error {
	relationAgencyId := c.Params("relationAgencyId")

	var payload *models.UpdateRelationAgencySchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var relationAgency models.RelationAgency
	result := config.DB.First(&relationAgency, "id = ?", relationAgencyId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Relation Agency with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	if payload.RelationAgencyName != "" {
		updates["relation_agency_name"] = payload.RelationAgencyName
	}

	if payload.RelationAgencyAddress != "" {
		updates["relation_agency_address"] = payload.RelationAgencyAddress
	}

	if payload.RelationAgencyPhone != "" {
		updates["relation_agency_phone"] = payload.RelationAgencyPhone
	}

	if payload.RelationAgencyEmail != "" {
		updates["relation_agency_email"] = payload.RelationAgencyEmail
	}

	if payload.RelationAgencyWebsite != "" {
		updates["relation_agency_website"] = payload.RelationAgencyWebsite
	}

	updates["updated_at"] = time.Now()

	config.DB.Model(&relationAgency).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": relationAgency}})
}

func FindRelationAgencyById(c *fiber.Ctx) error {
	relationAgencyId := c.Params("relationAgencyId")

	var relationAgency models.RelationAgency
	result := config.DB.First(&relationAgency, "id = ?", relationAgencyId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No RelationAgency with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": relationAgency}})
}

func RelationAgencyDelete(c *fiber.Ctx) error {
	relationAgencyId := c.Params("RelationAgencyId")

	result := config.DB.Delete(&models.RelationAgency{}, "id = ?", relationAgencyId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No RelationAgency with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
}

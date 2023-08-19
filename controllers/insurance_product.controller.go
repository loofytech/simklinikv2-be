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

func CreateInsuranceProductHandler(c *fiber.Ctx) error {
	var payload *models.CreateInsuranceProductSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	now := time.Now()
	newInsuranceProduct := models.InsuranceProduct{
		InsuranceProductName:        payload.InsuranceProductName,
		InsuranceProductAdminFee:    payload.InsuranceProductAdminFee,
		InsuranceProductMaxAdminFee: payload.InsuranceProductMaxAdminFee,
		InsuranceProductStamp:       payload.InsuranceProductStamp,
		RelationAgencyId:            payload.RelationAgencyId,
		CreatedAt:                   now,
		UpdatedAt:                   now,
	}

	result := config.DB.Create(&newInsuranceProduct)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "InsuranceProduct already exist"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   fiber.Map{"InsuranceProduct": newInsuranceProduct},
	})
}

func FindInsuranceProduct(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var insuranceProduct []models.InsuranceProduct
	results := config.DB.Limit(intLimit).Offset(offset).Find(&insuranceProduct)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(insuranceProduct), "user": insuranceProduct})
}

func UpdateInsuranceProduct(c *fiber.Ctx) error {
	insuranceProductId := c.Params("insuranceProductId")

	var payload *models.UpdateInsuranceProductSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var insuranceProduct models.InsuranceProduct
	result := config.DB.First(&insuranceProduct, "id = ?", insuranceProductId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No InsuranceProduct with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	if payload.InsuranceProductName != "" {
		updates["insurance_product_name"] = payload.InsuranceProductName
	}

	if payload.InsuranceProductAdminFee != 0 {
		updates["insurance_product_admin_fee"] = payload.InsuranceProductAdminFee
	}

	if payload.InsuranceProductMaxAdminFee != 0 {
		updates["insurance_product_max_admin_fee"] = payload.InsuranceProductAdminFee
	}

	if payload.InsuranceProductStamp != 0 {
		updates["insurance_product_stamp"] = payload.InsuranceProductStamp
	}

	// if payload.RelationAgencyId != nil {
	// 	updates["relation_agency_id"] = payload.RelationAgencyId
	// }

	updates["updated_at"] = time.Now()

	config.DB.Model(&insuranceProduct).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": insuranceProduct}})
}

func FindInsuranceProductById(c *fiber.Ctx) error {
	insuranceProductId := c.Params("insuranceProductId")

	var insuranceProduct models.InsuranceProduct
	result := config.DB.First(&insuranceProduct, "id = ?", insuranceProductId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No InsuranceProduct with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": insuranceProduct}})
}

func InsuranceProductDelete(c *fiber.Ctx) error {
	insuranceProductId := c.Params("insuranceProductId")

	result := config.DB.Delete(&models.InsuranceProduct{}, "id = ?", insuranceProductId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No InsuranceProduct with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
}

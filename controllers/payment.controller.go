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

func CreatePaymentHandler(c *fiber.Ctx) error {
	var payload *models.CreatePaymentSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	now := time.Now()
	newPayment := models.Payment{
		PaymentName:   payload.PaymentName,
		PaymentStatus: payload.PaymentStatus,
		PaymentSlug:   payload.PaymentSlug,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	result := config.DB.Create(&newPayment)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Payment already exist"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   fiber.Map{"Payment": newPayment},
	})
}

func FindPayment(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var payment []models.Payment
	results := config.DB.Limit(intLimit).Offset(offset).Find(&payment)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(payment), "data": payment})
}

func UpdatePayment(c *fiber.Ctx) error {
	paymentId := c.Params("paymentId")

	var payload *models.UpdatePaymentSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var payment models.Payment
	result := config.DB.First(&payment, "id = ?", paymentId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Payment with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	if payload.PaymentName != "" {
		updates["payment_name"] = payload.PaymentName
	}

	if payload.PaymentStatus != nil {
		updates["payment_active"] = payload.PaymentStatus
	}

	updates["updated_at"] = time.Now()

	config.DB.Model(&payment).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"data": payment}})
}

func FindPaymentById(c *fiber.Ctx) error {
	paymentId := c.Params("paymentId")

	var payment models.Payment
	result := config.DB.First(&payment, "id = ?", paymentId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Payment with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"data": payment}})
}

func PaymentDelete(c *fiber.Ctx) error {
	paymentId := c.Params("paymentId")

	result := config.DB.Delete(&models.Payment{}, "id = ?", paymentId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Payment with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
}

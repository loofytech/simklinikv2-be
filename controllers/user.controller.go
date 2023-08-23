package controllers

import (
	"strconv"
	"strings"
	"time"

	"sim-klinikv2/config"
	"sim-klinikv2/models"
	"sim-klinikv2/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateUserHandler(c *fiber.Ctx) error {
	var payload *models.CreateUserSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	password, _ := utils.HashPassword(payload.Password)

	now := time.Now()
	newUser := models.User{
		Name:      payload.Name,
		Username:  payload.Username,
		Email:     payload.Email,
		Password:  password,
		RoleId:    payload.RoleId,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := config.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "User already exist"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   fiber.Map{"user": newUser},
	})
}

func FindUser(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var user []models.User
	results := config.DB.Limit(intLimit).Offset(offset).Preload("Role").Find(&user)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(user), "user": user})
}

func UpdateUser(c *fiber.Ctx) error {
	userId := c.Params("userId")

	var payload *models.UpdateUserSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var user models.User
	result := config.DB.First(&user, "id = ?", userId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No user with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	if payload.Name != "" {
		updates["name"] = payload.Name
	}
	if payload.Username != "" {
		updates["username"] = payload.Username
	}
	if payload.Email != "" {
		updates["email"] = payload.Email
	}

	password, _ := utils.HashPassword(payload.Password)
	if payload.Password != "" {
		updates["password"] = password
	}
	// if payload.RoleId := "" {
	// 	updates["role_id"] = payload.RoleId
	// }

	updates["updated_at"] = time.Now()

	config.DB.Model(&user).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}

func FindUserById(c *fiber.Ctx) error {
	userId := c.Params("userId")

	var user models.User
	result := config.DB.Preload("Role").First(&user, "id = ?", userId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No user with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}

func UserDelete(c *fiber.Ctx) error {
	userId := c.Params("userId")

	result := config.DB.Delete(&models.User{}, "id = ?", userId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No user with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
}
